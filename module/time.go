package module

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"time"
)

type TimeModule struct{}

func (m *TimeModule) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := update.Message.Text
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	respondMessage := fmt.Sprint("It's currently " + currentTime)
	if message == "/time" {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   respondMessage,
		})
		fmt.Print("[LOG] Time module executed successfully")
	}
}

func init() {
	RegisterModule(&TimeModule{})
}
