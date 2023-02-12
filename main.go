package main

import (
	"os"
	"telegram/telegram"
)

func main() {

	botApi := os.Getenv("BOT_API")
	telegram.SendTelegram(botApi, "Alerta de servidor down")

}
