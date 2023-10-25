package internal

import (
	"my_server/pkg/httpex"
	"os"
	"strings"
)

func DownloadHandler(request *httpex.HttpRequest) httpex.Response {
	filePath := resourcesPath + string(os.PathSeparator) + "apps"
	splits := strings.Split(request.RouteParam()["subPath"], "/")
	for _, element := range splits {
		if element != "" {
			filePath += string(os.PathSeparator)
			filePath += element
		}
	}
	return httpex.FileResponse(filePath)
}
