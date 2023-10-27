package main

import (
	"fmt"
	"my_server/internal"
	"my_server/pkg/httpex"
	"my_server/pkg/tg_bot"
)

func main() {
	internal.InitVars()
	tg_bot.Init(internal.BotToken, internal.WebbookUrl, internal.SecretToken)
	tg_bot.SetOnUpdate(internal.OnUpdate)
	httpex.Get("/", internal.HomeHandler)
	httpex.Get("/version", internal.VersionHandler)
	httpex.Get("/updates", internal.UpdatesHandler)
	httpex.Forward("/picaapi/", "https://picaapi.picacomic.com")
	httpex.Get("/download", internal.DownloadPageHtml)
	httpex.Get("/static/<fileName>", internal.StaticHandler)
	httpex.Mount("/download", internal.DownloadHandler)
	httpex.Mount("/storage", internal.StorageHandler)
	httpex.Mount("/storage/download", internal.StorageHandler)
	httpex.Post("/log", func(request *httpex.HttpRequest) httpex.Response {
		return httpex.TextResponse("ok")
	})
	httpex.Forward("/jmComic/", "https://www.jmapinode3.cc")
	httpex.Post("/logs", func(request *httpex.HttpRequest) httpex.Response {
		return httpex.TextResponse("ok")
	})
	httpex.Post("/tgbot", func(request *httpex.HttpRequest) httpex.Response {
		if request.Headers().Get("X-Telegram-Bot-Api-Secret-Token") != internal.SecretToken {
			fmt.Println("Error Header: ", request.Headers().Get("X-Telegram-Bot-Api-Secret-Token"))
			return httpex.BadRequestResponse()
		}
		tg_bot.HandleWebhookData(request.Data())
		return httpex.TextResponse("ok")
	})
	httpex.ListenAndServer(internal.Port)
}
