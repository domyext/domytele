package module

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Module interface {
	Handle(ctx context.Context, b *bot.Bot, update *models.Update)
}

type CallbackModule interface {
	CallbackHandle(ctx context.Context, b *bot.Bot, update *models.Update)
}

var modules []Module

func RegisterModule(m Module) {
	modules = append(modules, m)
}

func DispatchMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	for _, m := range modules {
		m.Handle(ctx, b, update)
	}
}

func DispatchCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	for _, m := range modules {
		if callbackModule, ok := m.(CallbackModule); ok {
			callbackModule.CallbackHandle(ctx, b, update)
		}
	}
}
