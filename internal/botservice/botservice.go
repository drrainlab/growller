package botservice

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	bot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func init() {

}

type Config struct {
	Port  string `env:"BOT_SERVICE_PORT,default=9999"`
	Token string `env:"BOT_API_TOKEN,default=6558502198:AAHny0Opfd9XbEpJ5t5t5gBuhb4v_B4aNxI"`
}

type IBotService interface {
	Run(ctx context.Context) error
	Instance() *bot.Bot
	SendMessage(msg string) error
	Stop(ctx context.Context)
}

func NewBotService(cfg *Config) (IBotService, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config is nil")
	}

	ctx := context.Background()

	service := &botService{
		cfg: cfg,
	}

	b, err := bot.New(cfg.Token,
		bot.WithDefaultHandler(service.NewHandler()),
		bot.WithHTTPClient(time.Second*1, http.DefaultClient))
	if err != nil {
		log.Fatalf("error creating bot service: %v", err)
	}
	fmt.Println("connected to TG")

	b.RegisterHandler(bot.HandlerTypeMessageText, "/", bot.MatchTypePrefix, service.NewCommandHandler())

	service.bot = b

	// call methods.SetWebhook if needed

	// b.SetWebhook(ctx, )

	// go func() {
	// 	http.ListenAndServe(":2000", b.WebhookHandler())
	// }()

	go b.Start(ctx)
	if err != nil {
		return nil, err
	}

	// service.bot.Debug = true
	return service, nil
}

type botService struct {
	cfg   *Config
	bot   *bot.Bot
	updCh chan *models.Update
	cmdCh chan *models.Update

	done <-chan struct{}
}

func (s *botService) Instance() *bot.Bot {
	return s.bot
}

func (s *botService) NewHandler() func(ctx context.Context, b *bot.Bot, update *models.Update) {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		fmt.Println(update.Message.Chat.ID)
		// this handler will be called for all updates
	}
}

func (s *botService) NewCommandHandler() func(ctx context.Context, bot *bot.Bot, update *models.Update) {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		fmt.Println(update)
		// this handler will be called for all updates
	}
}

func (s *botService) SendMessage(msg string) error {
	p := &bot.SendMessageParams{
		ChatID:    150247499,
		ParseMode: models.ParseModeHTML,
	}
	p.Text = msg

	_, err := s.bot.SendMessage(context.Background(), p)
	return err
}

func (s *botService) Run(ctx context.Context) error {

	log.Println("running telegram bot service")

	go func() {
		defer func() {
			log.Println("exit botservice")
		}()

		for {
			select {
			case <-ctx.Done():
				s.Stop(ctx)
				return
			case <-s.done:
				s.Stop(ctx)
				return
			}
		}
	}()
	return nil
}

func (s *botService) Stop(ctx context.Context) {
	s.bot.Close(ctx)
}
