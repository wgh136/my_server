package httpex

import (
	"bytes"
	"io"
	"net/http"
	"slices"
)

var (
	shouldRemoveFromHeader = []string{"connection", "host", "sec-fetch-site", "sec-fetch-user", "sec-fetch-dest",
		"sec-fetch-mode", "sec-ch-ua-platform", "sec-ch-ua", "upgrade-insecure-requests",
		"sec-ch-ua-mobile", "cache-control", "date", "transfer-encoding",
		"strict-transport-security", "content-encoding", "cf-cache-status", "server",
		"cf-ray", "Transfer-Encoding"}
)

func forward(request *HttpRequest, to string) Response {
	if request.routeParam["subPath"][0] != '/' {
		request.routeParam["subPath"] = "/" + request.routeParam["subPath"]
	}
	if to[len(to)-1] == '/' {
		to = to[0 : len(to)-1]
	}
	to += request.routeParam["subPath"]
	if len(request.requestParam) != 0 {
		to += "?"
		for key, value := range request.requestParam {
			to += key
			to += "="
			to += value
			to += "&"
		}
		to = to[0 : len(to)-1]
	}
	logPrinter("Fetch data From: " + to)
	client := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       3 * 1000 * 1000 * 1000,
	}

	req, err := http.NewRequest(request.Method(), to, bytes.NewBuffer(request.data))

	for key, values := range request.Headers() {
		if !(slices.Contains(shouldRemoveFromHeader, key)) {
			var value = ""
			for index, element := range values {
				value += element
				if index != len(values)-1 {
					value += ";"
				}
			}
			req.Header.Set(key, value)
		}
	}

	if err != nil {
		return BadRequestResponse()
	}

	response, err := client.Do(req)
	if err != nil {
		return BadRequestResponse()
	}

	resBody, _ := io.ReadAll(response.Body)

	headerMap := make(map[string]string)
	for key, values := range response.Header {
		if slices.Contains(shouldRemoveFromHeader, key) {
			continue
		}

		var value = ""
		for index, element := range values {
			value += element
			if index != len(values)-1 {
				value += ";"
			}
		}
		headerMap[key] = value
	}

	return Response{
		response.StatusCode,
		resBody,
		headerMap,
		nil,
	}
}
