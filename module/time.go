package module

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strings"
	"time"
)

type TimeModule struct{}

func (m *TimeModule) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := update.Message.Text
	currentTime := time.Now().Format("Mon Jan 2 15:04:05 UTC 2006")
	respondMessage := fmt.Sprint("It's currently " + currentTime)
	if strings.HasPrefix(message, "/time") {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   respondMessage,
		})
		fmt.Println("[LOG] Time module executed successfully")
	}
}

func init() {
	RegisterModule(&TimeModule{})
}
