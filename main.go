package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/dig"
	"hanacore/config"
	"hanacore/module"
	"hanacore/utils/console"
	"os"
	"os/signal"
)

func NewBot(token string) (*bot.Bot, error) {
	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
		bot.WithCallbackQueryDataHandler("button", bot.MatchTypePrefix, callbackHandler),
	}
	return bot.New(token, opts...)
}

func main() {
	container := dig.New()

	container.Provide(func() string { return config.BotToken })
	container.Provide(NewBot)

	module.RegisterModules(container)

	err := container.Invoke(runBot)
	if err != nil {
		panic(err)
	}
}

func runBot(b *bot.Bot, modules []module.Module) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	for _, mod := range modules {
		module.RegisterModule(mod)
	}

	console.ShowBotInfo(b)
	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	module.DispatchMessage(ctx, b, update)
}

func callbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	module.DispatchCallback(ctx, b, update)
}
