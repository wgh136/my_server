package internal

import (
	"my_server/pkg/tg_bot"
	"strconv"
	"strings"
)

var (
	userNameToChatID = map[string]int{}
)

func OnUpdate(update tg_bot.Update) {
	userNameToChatID[update.Message.From.Username] = update.Message.Chat.ID
	if *update.Message.Text == "/start" {
		tg_bot.SendMessage("Welcome!\nYou can send me message.", update.Message.Chat.ID)
		return
	}
	if update.Message.From.Username == adminName {
		meChatId = strconv.Itoa(update.Message.Chat.ID)
		if (*update.Message.Text)[0] == '@' {
			username := strings.Split(*update.Message.Text, "\n")[0]
			username = username[1:]
			chatId := userNameToChatID[username]
			tg_bot.SendMessage(strings.Replace(*update.Message.Text, "@"+username+"\n", "", 1), chatId)
		}
		tg_bot.SendMessage("ok", meChatId)
		return
	}
	if meChatId != "" {
		tg_bot.ForwardMessage(update.Message, meChatId)
		tg_bot.SendMessage("@"+update.Message.From.Username, meChatId)
	}
}
