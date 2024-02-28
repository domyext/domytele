package module

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Module defines the interface for a bot module.
type Module interface {
	Handle(ctx context.Context, b *bot.Bot, update *models.Update)
}

func RegisterModule(m Module) {
	modules = append(modules, m)
}

var modules []Module

func DispatchMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	for _, m := range modules {
		m.Handle(ctx, b, update)
	}
}
