package telegram

import (
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Message struct {
	Text    string `json:"text"`
	GroupId int64  `json:"group_id"`
}

func SendTelegram(botApi string, mensagem string) {

	bot, err := tgbotapi.NewBotAPI(botApi)
	if err != nil {
		panic(err)
	}
	message := Message{}
	message.Text = mensagem
	groupId := os.Getenv("TELEGRAM_GROUP_ID")
	if groupId == "" {
		panic("TELEGRAM_GROUP_ID is not set")
	}
	message.GroupId, err = strconv.ParseInt(groupId, 10, 64)
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	alertText := tgbotapi.NewMessage(message.GroupId, message.Text)
	bot.Send(alertText)

}
