package config

import (
	"context"
	"fmt"
	"growwwler/internal/botservice"
	"growwwler/internal/httpservice"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	BotConfig   botservice.Config
	HttpService httpservice.Config
}

// ConfigureWithDefaults sets up default values of Config
// change this method if you want to change default behavior
func (c *Config) ConfigureWithDefaults() (cfg *Config) {
	// ctx := context.Background()

	// local := new(Config)

	// if c.HTTPService.SSLCert == "" {
	// 	certFile, err := os.ReadFile("deploy/secret/hotelcert.pem")
	// 	if err != nil {
	// 		logger.Logger(ctx).Fatal("Error read pem-file", zap.Error(err))
	// 	}
	// 	c.HTTPService.SSLCert = string(certFile)
	// }

	// if c.HTTPService.SSLKey == "" {
	// 	certFile, err := os.ReadFile("deploy/secret/hotelkey.pem")
	// 	if err != nil {
	// 		logger.Logger(ctx).Fatal("Error read pem-file", zap.Error(err))
	// 	}
	// 	c.HTTPService.SSLKey = string(certFile)
	// }

	return
}

func New(ctx context.Context) (*Config, error) {
	conf := new(Config)
	if err := envconfig.Process(ctx, conf); err != nil {
		return nil, fmt.Errorf("config processing error: %s", err)
	}
	return conf, nil
}
