package console

import (
	"fmt"
	"time"
)

func ShowLog(module, senderID string) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	logOutput := fmt.Sprintf("[%s] [%s] [LOG] %s Module executed successfully.", currentTime, senderID, module)
	fmt.Println(logOutput)
}
