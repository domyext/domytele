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
		"What is the most embarrassing thing that has ever happened to you?",
		"Have you ever lied to get out of trouble? What was the lie?",
		"What's the most childish thing you still do?",
		"Have you ever cheated on a test or exam? What happened?",
		"What is your biggest fear?",
		"Have you ever had a crush on someone in this group? Who?",
		"What is the most adventurous thing you've ever done?",
		"Have you ever stolen anything? What was it?",
		"If you could swap lives with someone for a day, who would it be and why?",
		"What's the most embarrassing thing you've done in front of a crush?",
		"What's the most embarrassing thing you've ever said or done in front of your crush?",
		"Have you ever pretended to be sick to get out of something? What was it?",
		"What's the weirdest dream you've ever had?",
		"Have you ever stalked someone on social media? Who was it?",
		"What's the most trouble you've ever gotten into at school?",
		"Have you ever peed in a pool?",
		"If you could erase one past experience, what would it be?",
		"What's the most embarrassing thing your parents have caught you doing?",
		"Have you ever been caught picking your nose?",
		"What's the most embarrassing thing you've worn in public?",
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
		"Go outside and sing \"Happy Birthday\" to a stranger.",
		"Send a message to your crush confessing your feelings (or a fake confession if you're already together).",
		"Do your best impression of a celebrity and post it on social media.",
		"Eat a spoonful of a condiment (e.g., mustard, ketchup, hot sauce) of the group's choice.",
		"Call a random number and try to have a conversation with whoever picks up.",
		"Wear socks on your hands for the next three rounds.",
		"Do 20 push-ups right now.",
		"Send a text to your parents saying you're dropping out of school (follow up later to let them know it's a joke!).",
		"Let someone draw a funny mustache on your face with permanent marker (or washable if preferred).",
		"Do a one-minute stand-up comedy routine improvised on the spot.",
		"Text your best friend and tell them you've won a million dollars and you're giving them half.",
		"Put an ice cube down your shirt and leave it there until it melts.",
		"Go outside and do the chicken dance on the sidewalk for one minute.",
		"Post a funny selfie on your social media accounts right now.",
		"Let someone give you a wedgie.",
		"Speak in an accent for the next three rounds.",
		"Swap clothes with the person next to you.",
		"Do a handstand against the wall for one minute.",
		"Send a voice message to your crush singing a love song.",
		"Wear your clothes backward for the next hour.",
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
