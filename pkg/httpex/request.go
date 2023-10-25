package httpex

import (
	"io"
	"net/http"
	"strings"
)

type HttpRequest struct {
	method       string
	uri          string
	requestParam map[string]string
	routeParam   map[string]string
	data         []byte
	headers      http.Header
}

func (h HttpRequest) Method() string {
	return h.method
}

func (h HttpRequest) Uri() string {
	return h.uri
}

func (h HttpRequest) RequestParam() map[string]string {
	return h.requestParam
}

func (h HttpRequest) RouteParam() map[string]string {
	return h.routeParam
}

func (h HttpRequest) Data() []byte {
	return h.data
}

func (h HttpRequest) Headers() http.Header {
	return h.headers
}

func getRequestFromHttp(r *http.Request) (*HttpRequest, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	params := map[string]string{}
	if strings.Contains(r.RequestURI, "?") {
		paramStr := strings.Split(r.RequestURI, "?")[1]
		paramItems := strings.Split(paramStr, "&")
		for _, item := range paramItems {
			lr := strings.Split(item, "=")
			params[lr[0]] = lr[1]
		}
	}
	return &HttpRequest{r.Method, r.RequestURI, params, nil, body, r.Header}, nil
}
