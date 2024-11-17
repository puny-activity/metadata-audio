package cfg

type Config struct {
	Logger Logger
}

func Parse() (*Config, error) {
	return &Config{
		Logger: Logger{
			Level: "debug",
		},
	}, nil
}
