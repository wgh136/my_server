package internal

import (
	"my_server/pkg/httpex"
	"net/http"
	"strings"
)

func StorageHandler(request *httpex.HttpRequest) httpex.Response {
	uri := request.RouteParam()["subPath"]
	if uri[0] == '/' {
		uri = uri[1:]
	}
	if strings.Index(uri, "https:/") != -1 && strings.Index(uri, "https://") == -1 {
		uri = strings.Replace(uri, "https:/", "https://", 1)
	}
	res, err := http.Get(uri)
	if err != nil || res.StatusCode != 200 {
		return httpex.BadRequestResponse()
	}
	return httpex.ByteResponse(res.Body, res.Header.Get("content-type"))
}
