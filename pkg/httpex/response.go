package httpex

import (
	"io"
	"os"
	"strconv"
)

type Response struct {
	statusCode int
	// Response data, if it is nil, reader will be used
	data    []byte
	headers map[string]string
	// Used for large response
	reader io.Reader
}

func HtmlResponse(html string) Response {
	return Response{200, []byte(html), map[string]string{
		"Content-Type": "text/html",
	}, nil}
}

func CssResponse(css string) Response {
	return Response{200, []byte(css), map[string]string{
		"Content-Type": "text/css",
	}, nil}
}

func TextResponse(text string) Response {
	return Response{200, []byte(text), map[string]string{
		"Content-Type": "text/plain",
	}, nil}
}

func InternalErrorResponse(text string) Response {
	return Response{500, []byte(text), map[string]string{
		"Content-Type": "text/plain",
	}, nil}
}

func NotFoundResponse() Response {
	return Response{404, []byte("Not Found"), map[string]string{
		"Content-Type": "text/plain",
	}, nil}
}

func BadRequestResponse() Response {
	return Response{400, []byte("Bad Request"), map[string]string{
		"Content-Type": "text/plain",
	}, nil}
}

func FileResponse(filePath string) Response {
	file, err := os.Open(filePath)
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return NotFoundResponse()
	}
	return Response{200, nil, map[string]string{
		"Content-Type":   "application/octet-stream",
		"Content-Length": strconv.FormatInt(fileInfo.Size(), 10),
	}, file}
}

func ByteResponse(reader io.Reader, contentType string) Response {
	return Response{200, nil, map[string]string{
		"Content-Type": contentType,
	}, reader}
}
