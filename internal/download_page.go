package internal

import (
	"io"
	"my_server/pkg/httpex"
	"os"
)

func DownloadPageHtml(request *httpex.HttpRequest) httpex.Response {
	file, err := os.Open(resourcesPath + string(os.PathSeparator) + "web" + string(os.PathSeparator) + "download" + string(os.PathSeparator) + "download.html")
	if err != nil {
		return httpex.InternalErrorResponse("Error")
	}
	html, err := io.ReadAll(file)
	if err != nil {
		return httpex.InternalErrorResponse("Error")
	}
	return httpex.HtmlResponse(string(html))
}
