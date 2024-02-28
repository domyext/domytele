package module

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type ChatInfoModule struct{}

func (m *ChatInfoModule) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := update.Message.Text
	senderID := bot.EscapeMarkdown(fmt.Sprintf("%d", update.Message.From.ID)) // Convert int64 to string
	chatType := bot.EscapeMarkdown(update.Message.Chat.Type)
	chatID := bot.EscapeMarkdown(fmt.Sprintf("%d", update.Message.Chat.ID)) // Convert int64 to string

	var msgText string
	if chatType == "group" {
		chat := `
* Chat Info *

* Group ID: * %s
* Sender ID: * %s
`
		msgText = fmt.Sprintf(chat, chatID, senderID)
	} else {
		chat := `
* Chat Info *

* Chat Type: * %s
* Sender ID: * %s
`
		msgText = fmt.Sprintf(chat, chatType, senderID)
	}

	if message == "/chat" {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    int64(update.Message.Chat.ID),
			ParseMode: "MarkdownV2",
			Text:      msgText,
		})
		fmt.Println("[LOG] Chat info module executed successfully")
	}
}

func init() {
	RegisterModule(&ChatInfoModule{})
}
