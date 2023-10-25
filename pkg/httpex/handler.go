package httpex

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

var (
	handlers   []httpHandler
	logPrinter = func(content string) {
		fmt.Println(content)
	}
)

type httpHandler struct {
	route   string
	method  string
	handler HttpHandler
}

type HttpHandler = func(request *HttpRequest) Response

func Get(route string, handler HttpHandler) {
	handlers = append(handlers, httpHandler{route, "GET", handler})
}

func Post(route string, handler HttpHandler) {
	handlers = append(handlers, httpHandler{route, "POST", handler})
}

func Put(route string, handler HttpHandler) {
	handlers = append(handlers, httpHandler{route, "PUT", handler})
}

func Delete(route string, handler HttpHandler) {
	handlers = append(handlers, httpHandler{route, "DELETE", handler})
}

// Mount : match all request whose request path starts with [route]
func Mount(route string, handler HttpHandler) {
	if route[len(route)-1] == '/' {
		route = route[0 : len(route)-1]
	}
	handlers = append(handlers, httpHandler{route, "MOUNT", handler})
}

// Forward request to another server
func Forward(route string, to string) {
	Mount(route, func(request *HttpRequest) Response {
		return forward(request, to)
	})
}

// check whether handler's route match request uri
func match(route string, uri string, params map[string]string, moute bool) bool {
	if strings.Contains(uri, "?") {
		uri = strings.Split(uri, "?")[0]
	}

	if route == uri {
		return true
	}

	routeSegments := strings.Split(route, "/")
	requestSegments := strings.Split(uri, "/")

	if !moute && len(routeSegments) != len(requestSegments) {
		return false
	}

	if moute && len(routeSegments) > len(requestSegments) {
		return false
	}

	i, j := 0, 0

	for ; i < len(routeSegments) && j < len(requestSegments); i, j = i+1, j+1 {
		ro, re := routeSegments[i], requestSegments[j]
		if ro == re {
			continue
		}
		if len(ro) > 2 && ro[0] == '<' && ro[len(ro)-1] == '>' {
			params[ro[1:len(ro)-1]] = re
			continue
		}
		return false
	}

	if j < len(requestSegments) {
		subPath := ""
		for ; j < len(requestSegments); j++ {
			subPath += "/"
			subPath += requestSegments[j]
		}
		params["subPath"] = subPath
	}

	return true
}

// handler request
func handleFunc(w http.ResponseWriter, r *http.Request) {
	logPrinter(r.Method + ": " + r.RequestURI)

	for _, handler := range handlers {
		routeParams := map[string]string{}
		if (handler.method == r.Method && match(handler.route, r.RequestURI, routeParams, false)) ||
			(handler.method == "MOUNT" && match(handler.route, r.RequestURI, routeParams, true)) {
			request, err := getRequestFromHttp(r)
			request.routeParam = routeParams
			if err != nil {
				w.WriteHeader(500)
				_, err := fmt.Fprint(w, "An error occurred!")
				if err != nil {
					logPrinter(err.Error())
					return
				}
				return
			}
			res := handler.handler(request)

			for key, value := range res.headers {
				w.Header().Set(key, value)
			}

			if err != nil {
				w.WriteHeader(500)
				logPrinter(err.Error())
			} else {
				w.WriteHeader(res.statusCode)
			}

			if res.data != nil {
				_, err = w.Write(res.data)
			} else if res.reader != nil {
				_, err = io.Copy(w, res.reader)
			}

			return
		}
	}
	w.WriteHeader(404)
	_, err := fmt.Fprintf(w, "Not found")
	if err != nil {
		logPrinter(err.Error())
		return
	}
}

// ListenAndServer : start http server
func ListenAndServer(port string) {
	http.HandleFunc("/", handleFunc)
	logPrinter("Start server at: http://localhost:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		logPrinter("Failed to start server: " + err.Error())
		return
	}
}

func SetLogPrinter(printer func(content string)) {
	logPrinter = printer
}
