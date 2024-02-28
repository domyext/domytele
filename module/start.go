package module

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"math/rand"
	"strings"
	"time"
)

type StartModule struct{}

func (m *StartModule) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := update.Message.Text
	if strings.HasPrefix(message, "/start") {
		randomMessage := getRandomMessage()
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   randomMessage,
		})
		fmt.Println("[LOG] Start module executed successfully")
	}
}

func init() {
	RegisterModule(&StartModule{})
}

func getRandomMessage() string {
	messages := []string{
		"Hello there! I am Hana",
		"G'day mate! What's up?",
		"Aye, howdy?",
		"Greetings!",
		"Yoyoyoyoyoyooo whaddup!",
	}
	rand.Seed(time.Now().UnixNano())
	return messages[rand.Intn(len(messages))]
}
