package logger

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"hanacore/config"
	"time"
)

func ShowLog(module, senderID string) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	logOutput := fmt.Sprintf("[%s] [%s] [LOG] %s Module executed successfully.", currentTime, senderID, module)
	fmt.Println(logOutput)
}

func SendLog(ctx context.Context, b *bot.Bot, update *models.Update, senderID, module string) {
	if config.BotLogChat != "" {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		msgText := fmt.Sprintf("[%s] [%s] [LOG] %s Module executed successfully.", currentTime, senderID, module)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: config.BotLogChat,
			Text:   msgText,
		})
	}
}
