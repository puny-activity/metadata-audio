package lggr

import (
	"context"
	"github.com/puny-activity/metadata-parser-audio/pkg/base"
	"log/slog"
)

type ContextHandler struct {
	slog.Handler
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	actionID := ctx.Value(base.CtxActionID)
	if actionID != nil {
		r.AddAttrs(slog.String(base.CtxActionID, actionID.(string)))
	}
	return h.Handler.Handle(ctx, r)
}
