package cfg

import "github.com/puny-activity/metadata-parser-audio/internal/api/grpc"

type App struct {
	GRPC GRPC
}

func (c App) GetGRPCConfig() grpc.Config {
	return c.GRPC
}
