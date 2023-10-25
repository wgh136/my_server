package tg_bot

import (
	json2 "encoding/json"
	"fmt"
)

var (
	onUpdate = func(update Update) {
		fmt.Println("New update: ", update.UpdateID)
	}
)

// Start webhook.
//
// Not: this package do not have http server feature.
func startWebhook(url string, secretToken string) {
	request("setWebhook", map[string]string{
		"url":          url,
		"secret_token": secretToken,
	})
}

// HandleWebhookData should be called when new data received from webhook url
func HandleWebhookData(data []byte) {
	var update Update
	err := json2.Unmarshal(data, &update)
	if err != nil {
		return
	}
	onUpdate(update)
}

// SetOnUpdate set how to deal with update
func SetOnUpdate(function func(update Update)) {
	onUpdate = function
}

// ForwardMessage
//
// chatId must be string or int
func ForwardMessage(message *Message, chatId any) {
	json := map[string]interface{}{
		"chat_id":      chatId,
		"from_chat_id": message.Chat.ID,
		"message_id":   message.MessageID,
	}
	request("forwardMessage", json)
}

// SendMessage
//
// chatId must be string or int
func SendMessage(text string, chatId any) {
	json := map[string]interface{}{
		"chat_id": chatId,
		"text":    text,
	}
	request("sendMessage", json)
}
