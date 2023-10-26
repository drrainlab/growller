package httpservice

import "time"

type Config struct {
	ServiceName   string        `env:"SERVICE_NAME,default=growwwler"`
	AppEnv        string        `env:"APP_ENV,default=dev"`
	Debug         bool          `env:"DEBUG,default=false"`
	Port          int           `env:"PORT,default=420"`
	SSLEnabled    bool          `env:"SSL_ENABLED,default=false"`
	SSLKey        string        `env:"SSL_KEY,default="`
	SSLCert       string        `env:"SSL_CERT,default="`
	HTTPTimeout   time.Duration `env:"HTTP_TIMEOUT,default=70s"`
	ReadTimeout   time.Duration `env:"SERVER_READ_TIMEOUT,default=15s"`
	WriteTimeout  time.Duration `env:"SERVER_WRITE_TIMEOUT,default=15s"`
	KeyFilePath   string        `env:"SERVICE_KEYS_FILE_PATH"`
	AlwaysGateway bool          `env:"ALWAYS_GATEWAY,default=false"`
	CORSServices  []string      `env:"CORS_SERVICES,default="`
}
