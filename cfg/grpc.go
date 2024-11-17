package cfg

type GRPC struct {
	Port int
}

func (c GRPC) GetPort() int {
	return c.Port
}
