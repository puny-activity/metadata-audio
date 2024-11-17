package controller

import (
	"context"
	metadataparseraudio "github.com/puny-activity/metadata-parser-audio/api/proto/gen"
	"log/slog"
)

func (c *Controller) Extract(ctx context.Context, request *metadataparseraudio.ExtractRequest) (*metadataparseraudio.ExtractResponse, error) {
	c.log.LogAttrs(ctx, slog.LevelInfo, "Extract",
		slog.Any("request", request))
	return &metadataparseraudio.ExtractResponse{}, nil
}
