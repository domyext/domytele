package module

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"hanacore/utils/console"
	"math/rand"
	"strings"
	"time"
)

type ToDModule struct{}

func (m *ToDModule) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {
	moduleName := "Truth or Dare"
	senderID := bot.EscapeMarkdown(fmt.Sprintf("%d", update.Message.From.ID)) // Convert int64 to string

	message := update.Message.Text
	switch {
	case strings.HasPrefix(message, "/truth"):
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   getTruthMsg(),
		})
		console.ShowLog(moduleName, senderID)
	case strings.HasPrefix(message, "/dare"):
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   getDareMsg(),
		})
		console.ShowLog(moduleName, senderID)
	case strings.HasPrefix(message, "/tod"):
		challenge, challengeType := getRandomChallenge()

		text := fmt.Sprintf(`<b>You got %s !</b>

%s`,
			challengeType, challenge)

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			ParseMode: "HTML",
			Text:      text,
		})
		console.ShowLog(moduleName, senderID)
	}
}

func init() {
	RegisterModule(&ToDModule{})
}

func getTruthMsg() string {
	messages := []string{
		"Name 3 things you hate about your friend!",
		"What weird desire do you wish to have?",
		"Who is your crush?",
		"Who do you want to meet right now?",
	}
	rand.Seed(time.Now().UnixNano())
	return messages[rand.Intn(len(messages))]
}

func getDareMsg() string {
	messages := []string{
		"Do push-ups for 30 seconds by the side of the road.",
		"Invite a random person to vlog with you for 2 minutes.",
		"Say \"I love you\" to your crush.",
		"Call your crush or partner, then say \"I'm hungry\".",
		"Take a selfie right now and post it to your Instagram Story.",
	}
	rand.Seed(time.Now().UnixNano())
	return messages[rand.Intn(len(messages))]
}

func getRandomChallenge() (string, string) {
	truthMsg := getTruthMsg()
	dareMsg := getDareMsg()
	challenges := []string{truthMsg, dareMsg}
	challengeType := []string{"Truth", "Dare"}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(challenges))
	return challenges[index], challengeType[index]
}
