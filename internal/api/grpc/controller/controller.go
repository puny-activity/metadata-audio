package controller

import (
	metadataparseraudio "github.com/puny-activity/metadata-parser-audio/api/proto/gen"
	"log/slog"
)

type Controller struct {
	metadataparseraudio.MetadataParserServer
	log *slog.Logger
}

func New(log *slog.Logger) *Controller {
	return &Controller{
		log: log,
	}
}
