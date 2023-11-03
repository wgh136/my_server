package internal

import (
	json2 "encoding/json"
	"io"
	"my_server/pkg/tg_bot"
	"os"
	"strconv"
	"strings"
)

func OnUpdate(update tg_bot.Update) {
	addData(update.Message.From.Username, update.Message.Chat.ID)
	if *update.Message.Text == "/start" {
		tg_bot.SendMessage("Welcome!\nYou can send me message.", update.Message.Chat.ID)
		return
	}
	if update.Message.From.Username == adminName {
		meChatId = strconv.Itoa(update.Message.Chat.ID)
		if (*update.Message.Text)[0] == '@' {
			username := strings.Split(*update.Message.Text, "\n")[0]
			username = username[1:]
			chatId := readData()[username]
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

func readData() map[string]int {
	file, err := os.Open(resourcesPath + string(os.PathSeparator) + "charIDs.json")
	if err != nil {
		return map[string]int{}
	}
	defer file.Close()
	var res map[string]int
	data, err := io.ReadAll(file)
	if err != nil {
		return map[string]int{}
	}
	json2.Unmarshal(data, &res)
	return res
}

func addData(userName string, chatID int) {
	data := readData()
	data[userName] = chatID
	file, err := os.Open(resourcesPath + string(os.PathSeparator) + "charIDs.json")
	defer file.Close()
	if err != nil {
		return
	}
	bytes, err := json2.Marshal(data)
	if err != nil {
		return
	}
	file.Write(bytes)
}
