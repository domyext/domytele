package console

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"strings"
	"time"
)

func ShowBotInfo(b *bot.Bot) error {
	botInfo, err := b.GetMe(context.Background())
	if err != nil {
		return err
	}

	output := fmt.Sprintf("Running as: %s %s\nBot handle: @%s", botInfo.FirstName, botInfo.LastName, botInfo.Username)
	printBoxed(output)
	return nil
}

func printBoxed(content string) {
	lines := strings.Split(content, "\n")
	maxLength := maxLength(lines)
	border := strings.Builder{}
	border.WriteString(fmt.Sprintf("┌%s┐\n", strings.Repeat("─", maxLength+2)))
	for _, line := range lines {
		border.WriteString(fmt.Sprintf("│ %-*s │\n", maxLength, line))
	}
	border.WriteString(fmt.Sprintf("└%s┘", strings.Repeat("─", maxLength+2)))

	fmt.Println(border.String())
}

func maxLength(lines []string) int {
	maxLength := 0
	for _, line := range lines {
		length := len(line)
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}

func ShowLog(module, senderID string) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	logOutput := fmt.Sprintf("[%s] [%s] [LOG] %s Module executed successfully.", currentTime, senderID, module)
	fmt.Println(logOutput)
}
