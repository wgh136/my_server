package tg_bot

import (
	"fmt"
	"my_server/pkg/httpex_client"
)

func request(method string, data interface{}) interface{} {
	res, err := httpex_client.PostJson("https://api.telegram.org/bot"+token+"/"+method, data, map[string]string{
		"Content-Type": "application/json",
	})
	fmt.Println(err.Error())
	return res
}
