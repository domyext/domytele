package console

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"strings"
	"time"
)

func ShowBotInfo(b *bot.Bot) {
	botInfo, _ := b.GetMe(context.Background())

	output := fmt.Sprintf("Running as: %s %s\nBot handle: @%s", botInfo.FirstName, botInfo.LastName, botInfo.Username)

	printBoxed(output)
}

func printBoxed(content string) {
	lines := strings.Split(content, "\n")
	maxLength := maxLength(lines)
	border := fmt.Sprintf("┌%s┐\n", strings.Repeat("─", maxLength+2))
	for _, line := range lines {
		border += fmt.Sprintf("│ %-*s │\n", maxLength, line)
	}
	border += fmt.Sprintf("└%s┘", strings.Repeat("─", maxLength+2))

	fmt.Println(border)
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
