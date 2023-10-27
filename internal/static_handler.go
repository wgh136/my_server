package internal

import (
	"bytes"
	"io"
	"my_server/pkg/httpex"
	"os"
	"strings"
)

func StaticHandler(r *httpex.HttpRequest) httpex.Response {
	fileName := r.RouteParam()["fileName"]
	file, err := os.Open(resourcesPath + string(os.PathSeparator) + "static" + string(os.PathSeparator) + fileName)
	defer file.Close()
	if err != nil {
		return httpex.NotFoundResponse()
	}
	contentType := "text/plain"
	lr := strings.Split(fileName, ".")
	if lr[1] == "png" || lr[1] == "jpg" {
		contentType = "img/jpeg"
	}
	if lr[1] == "css" {
		contentType = "text/css"
	}
	data, err := io.ReadAll(file)

	return httpex.ByteResponse(bytes.NewReader(data), contentType)
}
