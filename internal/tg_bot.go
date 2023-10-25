package internal

import (
	"my_server/pkg/tg_bot"
	"strconv"
)

func OnUpdate(update tg_bot.Update) {
	if update.Message.From.Username == adminName {
		meChatId = strconv.Itoa(update.Message.Chat.ID)
		tg_bot.SendMessage("ok", meChatId)
		return
	}
	if meChatId != "" {
		tg_bot.ForwardMessage(update.Message, meChatId)
		tg_bot.SendMessage("@"+update.Message.From.Username, meChatId)
	}
}
