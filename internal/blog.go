package internal

import (
	"io"
	"my_server/pkg/httpex"
	"os"
)

func BlogListHandler(request *httpex.HttpRequest) httpex.Response {
	file, err := os.Open(resourcesPath + string(os.PathSeparator) + "blogs" + string(os.PathSeparator) + "info.json")
	defer file.Close()
	if err != nil {
		return httpex.InternalErrorResponse("error")
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return httpex.InternalErrorResponse("error")
	}
	return httpex.Response{StatusCode: 200, Data: data, Headers: map[string]string{
		"content-type":                "application/json",
		"Access-Control-Allow-Origin": "*",
	}}
}

func BlogDetailHandler(request *httpex.HttpRequest) httpex.Response {
	id := request.RouteParam()["id"]
	file, err := os.Open(resourcesPath + string(os.PathSeparator) + "blogs" + string(os.PathSeparator) + id + ".md")
	defer file.Close()
	if err != nil {
		return httpex.InternalErrorResponse("error")
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return httpex.InternalErrorResponse("error")
	}
	return httpex.Response{StatusCode: 200, Data: data, Headers: map[string]string{
		"content-type":                "text/plain",
		"Access-Control-Allow-Origin": "*",
	}}
}
