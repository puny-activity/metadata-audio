package cfg

import "github.com/puny-activity/metadata-parser-audio/internal/app"

type Config struct {
	App    App
	Logger Logger
}

func (c *Config) GetAppConfig() app.Config {
	return c.App
}

func Parse() (*Config, error) {
	return &Config{
		App: App{
			GRPC: GRPC{
				Port: 9090,
			},
		},
		Logger: Logger{
			Level: "debug",
		},
	}, nil
}
