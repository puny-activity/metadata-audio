package cfg

type Logger struct {
	Level string
}

func (l Logger) GetLevel() string {
	return l.Level
}
