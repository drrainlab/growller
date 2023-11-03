package httpservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"growwwler/internal/botservice"
	"growwwler/internal/models"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	cors "github.com/rs/cors"
)

type HTTPService struct {
	srv http.Server
}

func (h *HTTPService) Run(ctx context.Context, cfg *Config) {

	var serveErr error

	serveErr = h.srv.ListenAndServe()

	if serveErr != nil && !errors.Is(serveErr, http.ErrServerClosed) {
		log.Fatal(serveErr)
	}
}

func (h *HTTPService) Shutdown(ctx context.Context) error {
	err := h.srv.Shutdown(ctx)

	return err
}

func writeToFile(data string) error {
	var f *os.File
	f, err := os.Open("sensor_data.json")
	if errors.Is(err, fs.ErrNotExist) {
		f, err = os.Create("sensor_data.json")
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	_, err = f.WriteString(data)
	return err
}

func NewHTTPService(cfg *Config, tgBot botservice.IBotService) (*HTTPService, error) {
	// amw, err := authmiddleware.NewAuthenticationMiddleware(cfg.KeyFilePath, "/docs")
	// if err != nil {
	// 	return nil, err
	// }

	r := mux.NewRouter()
	r.Methods("POST").Path("/data").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var boxData models.BoxData
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("err reading body: ", err)
			w.WriteHeader(500)
			return
		}
		fmt.Println(string(raw))
		json.Unmarshal(raw, &boxData)
		tgmsg := fmt.Sprintf(`
			time: %s
			temperature: %.2f
			humidity: %.2f
			co2: %d
			ghum: %d
			fan on: %t
			humidifier on: %t
			pump on: %t
		`, boxData.Time,
			boxData.Temperature,
			boxData.Humidity,
			boxData.CO2,
			boxData.Ghum,
			boxData.FanState,
			boxData.HumidifierState,
			boxData.PumpState,
		)

		if err = tgBot.SendMessage(tgmsg); err != nil {
			log.Println(err)
		}

		// if err = writeToFile(string(raw)); err != nil {
		// 	log.Println(err)
		// }

		w.WriteHeader(200)
	})
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	middle := r
	// r.Use(handler.CORS)

	corsHandler := cors.AllowAll().Handler(middle)

	// r.Use(amw.AuthorizationByHeaders)

	return &HTTPService{
		srv: http.Server{
			Handler:      corsHandler,
			Addr:         fmt.Sprintf(":%d", cfg.Port),
			WriteTimeout: cfg.WriteTimeout,
			ReadTimeout:  cfg.ReadTimeout,
		},
	}, nil
}
