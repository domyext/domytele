package module

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"hanacore/utils/console"
	"strings"
	"time"
)

type TimeModule struct{}

func (m *TimeModule) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {
	moduleName := "Time"
	moduleCommand := "/time"
	senderID := bot.EscapeMarkdown(fmt.Sprintf("%d", update.Message.From.ID)) // Convert int64 to string

	message := update.Message.Text
	currentTime := time.Now().Format("Mon Jan 2 15:04:05 UTC 2006")
	respondMessage := fmt.Sprint("It's currently " + currentTime)
	if strings.HasPrefix(message, moduleCommand) {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   respondMessage,
		})
		console.ShowLog(moduleName, senderID)
	}
}

func init() {
	RegisterModule(&TimeModule{})
}
