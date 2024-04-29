package module

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	logger "hanacore/utils/Logger"
	"hanacore/utils/console"
	"strings"
)

type ChatInfoModule struct{}

func (m *ChatInfoModule) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {
	moduleName := "Chat Info"
	moduleCommand := "/chat"
	senderID := bot.EscapeMarkdown(fmt.Sprintf("%d", update.Message.From.ID)) // Convert int64 to string

	message := update.Message.Text
	chatType := bot.EscapeMarkdown(update.Message.Chat.Type)
	chatID := bot.EscapeMarkdown(fmt.Sprintf("%d", update.Message.Chat.ID)) // Convert int64 to string

	var msgText string
	if chatType == "group" {
		chat := "*Chat Info*\n\n*Group ID:* `" + chatID + "`\n*Sender ID:* `" + senderID + "`"
		msgText = chat
	} else {
		chat := "*Chat Info*\n\n*Chat Type:* `" + chatType + "`\n*Sender ID:* `" + senderID + "`"
		msgText = chat
	}
	if strings.HasPrefix(message, moduleCommand) {

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    int64(update.Message.Chat.ID),
			ParseMode: "MarkdownV2",
			Text:      msgText,
		})
		console.ShowLog(moduleName, senderID)
		logger.SendLog(ctx, b, update, senderID, moduleName)
	}
}

func init() {
	RegisterModule(&ChatInfoModule{})
}
