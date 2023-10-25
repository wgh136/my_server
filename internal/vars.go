package internal

import (
	json2 "encoding/json"
	"io"
	"os"
)

var (
	resourcesPath = ""
	Port          = "80"
	SecretToken   = ""
	meChatId      = ""
	BotToken      = ""
	WebbookUrl    = ""
	adminName     = ""
)

func Init() {
	file, _ := os.Open(os.Args[1])
	bytes, _ := io.ReadAll(file)
	var json map[string]string
	_ = json2.Unmarshal(bytes, &json)
	BotToken = json["botToken"]
	resourcesPath = json["resourcePath"]
	Port = json["port"]
	SecretToken = json["secretToken"]
	WebbookUrl = json["webhook"]
	adminName = json["admin"]
}
