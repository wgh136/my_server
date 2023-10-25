package internal

import (
	"io"
	"my_server/pkg/httpex"
	"os"
)

func HomeHandler(_ *httpex.HttpRequest) httpex.Response {
	content := ""
	file, err := os.Open(resourcesPath + string(os.PathSeparator) + "web" + string(os.PathSeparator) + "home.html")
	defer file.Close()
	if err != nil {
		return httpex.InternalErrorResponse("Internal Error")
	} else {
		contentBytes, err1 := io.ReadAll(file)
		if err1 != nil {
			return httpex.InternalErrorResponse("Internal Error")
		} else {
			content = string(contentBytes)
		}
	}
	return httpex.HtmlResponse(content)
}

func HomeCSSHandler(_ *httpex.HttpRequest) httpex.Response {
	content := ""
	file, err := os.Open(resourcesPath + string(os.PathSeparator) + "web" + string(os.PathSeparator) + "home.css")
	defer file.Close()
	if err != nil {
		return httpex.InternalErrorResponse("Internal Error")
	} else {
		contentBytes, err1 := io.ReadAll(file)
		if err1 != nil {
			return httpex.InternalErrorResponse("Internal Error")
		} else {
			content = string(contentBytes)
		}
	}
	return httpex.CssResponse(content)
}
