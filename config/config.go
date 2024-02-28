package config

import "os"

var BotToken = os.Getenv("TG_BOT_TOKEN")
