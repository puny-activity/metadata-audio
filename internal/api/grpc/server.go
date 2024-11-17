package grpc

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	metadataparseraudio "github.com/puny-activity/metadata-parser-audio/api/proto/gen"
	"github.com/puny-activity/metadata-parser-audio/internal/api/grpc/controller"
	"github.com/puny-activity/metadata-parser-audio/pkg/werr"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type Config interface {
	GetPort() int
}

type Server struct {
	config     Config
	controller *controller.Controller
	server     *grpc.Server
	log        *slog.Logger
}

func New(config Config, controller *controller.Controller, log *slog.Logger) *Server {
	return &Server{
		config:     config,
		controller: controller,
		log:        log,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.GetPort()))
	if err != nil {
		return werr.Wrap("failed to listen", err)
	}

	s.server = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(),
		),
	)

	metadataparseraudio.RegisterMetadataParserServer(s.server, s.controller)

	go func() {
		err = s.server.Serve(listener)
		if err != nil {
			panic(werr.Wrap("failed to serve", err))
		}
	}()

	return nil
}

func (s *Server) Stop() {
	s.server.GracefulStop()
}
