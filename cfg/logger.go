package cfg

type Logger struct {
	Level string
}

func (c Logger) GetLevel() string {
	return c.Level
}
