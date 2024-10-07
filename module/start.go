package module

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"math/rand"
	"time"
)

type StartModule struct{}

func ProvideStartModule() *StartModule {
	return &StartModule{}
}

func (m *StartModule) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message.Text == "/start" {
		messages := []string{
			"Hello there! I am Hana",
			"G'day mate! What's up?",
			"Aye, howdy?",
			"Greetings!",
			"Yoyoyoyoyoyooo whaddup!",
		}
		rand.Seed(time.Now().UnixNano())
		msgText := messages[rand.Intn(len(messages))]

		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   msgText,
		})
	}
}

func (m *StartModule) CallbackHandle(ctx context.Context, b *bot.Bot, update *models.Update) {

}
