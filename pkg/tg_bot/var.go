package tg_bot

var (
	token = ""
)

func Init(botToken string, webhook string, secretToken string) {
	token = botToken
	startWebhook(webhook, secretToken)
}
