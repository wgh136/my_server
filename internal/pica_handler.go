package internal

import (
	"io"
	"my_server/pkg/httpex"
	"os"
)

func VersionHandler(_ *httpex.HttpRequest) httpex.Response {
	content := ""
	file, err := os.Open(resourcesPath + string(os.PathSeparator) + "version.txt")
	defer file.Close()
	if err != nil {
		content = "Internal Error"
	} else {
		contentBytes, err1 := io.ReadAll(file)
		if err1 != nil {
			content = "Internal Error"
		} else {
			content = string(contentBytes)
		}
	}
	return httpex.TextResponse(content)
}

func UpdatesHandler(_ *httpex.HttpRequest) httpex.Response {
	content := ""
	file, err := os.Open(resourcesPath + string(os.PathSeparator) + "updates.txt")
	defer file.Close()
	if err != nil {
		content = "Internal Error"
	} else {
		contentBytes, err1 := io.ReadAll(file)
		if err1 != nil {
			content = "Internal Error"
		} else {
			content = string(contentBytes)
		}
	}
	return httpex.TextResponse(content)
}
