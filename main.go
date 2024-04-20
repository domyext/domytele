// main.go

package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"hanacore/utils/console"
	"os"
	"os/signal"

	"hanacore/config"
	"hanacore/module"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
		bot.WithCallbackQueryDataHandler("button", bot.MatchTypePrefix, callbackHandler),
	}

	b, err := bot.New(config.BotToken, opts...) // Update with your bot token
	if err != nil {
		panic(err)
	}

	console.ShowBotInfo(b)
	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	module.DispatchMessage(ctx, b, update)
}

func callbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	module.Dispatchcallback(ctx, b, update)
}
