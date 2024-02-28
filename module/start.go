package module

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strings"
)

type StartModule struct{}

func (m *StartModule) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := update.Message.Text
	if strings.HasPrefix(message, "/start") {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Hello there! I am Hana",
		})
		fmt.Print("[LOG] Start module executed successfully")
	}
}

func init() {
	RegisterModule(&StartModule{})
}
