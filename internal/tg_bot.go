package internal

import (
	"fmt"
	"io"
	"my_server/pkg/tg_bot"
	"os"
	"strconv"
	"strings"
)

func OnUpdate(update tg_bot.Update) {
	if update.Message.Chat.Type == "group" || update.Message.Chat.Type == "supergroup" {
		log(*update.Message.Text, "Group: "+update.Message.From.Username, update.Message.Chat.ID)
		return
	}
	log(*update.Message.Text, update.Message.From.Username, update.Message.Chat.ID)
	AddData(update.Message.From.Username, update.Message.Chat.ID)
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
	file, err := os.Open(resourcesPath + string(os.PathSeparator) + "chatIDs.txt")
	if err != nil {
		return map[string]int{}
	}
	defer file.Close()
	var res map[string]int
	data, err := io.ReadAll(file)
	if err != nil {
		return map[string]int{}
	}
	str := string(data)
	items := strings.Split(str, "\n")
	for _, item := range items {
		lr := strings.Split(item, ":")
		i, err := strconv.Atoi(lr[1])
		if err != nil {
			continue
		}
		res[lr[0]] = i
	}
	return res
}

func AddData(userName string, chatID int) {
	file, err := os.OpenFile(resourcesPath+string(os.PathSeparator)+"chatIDs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	dataToAppend := fmt.Sprintf("%s:%d\n", userName, chatID)
	_, err = file.WriteString(dataToAppend)
	if err != nil {
		fmt.Println("无法写入文件:", err)
	}
}

func log(content string, userName string, chatID int) {
	file, err := os.OpenFile(resourcesPath+string(os.PathSeparator)+"botlog.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	dataToAppend := fmt.Sprintf("%s:%s:%d\n", content, userName, chatID)
	_, err = file.WriteString(dataToAppend)
	if err != nil {
		fmt.Println("无法写入文件:", err)
	}
}
