# httpex

## About
This package aim to develop a simple server program quickly.

For complex server, please consider to use Gin or other framework.

## Getting Start
```go
func handler(req *httpex.HttpRequest) httpex.Response{
    return httpex.TextResponse("Hello " + req.RouteParam()["name"])
}

func main(){
	httpex.Get("/<name>", handler)
	httpex.ListenAndServer("80")
}
```