package httpex_client

import (
	"bytes"
	json2 "encoding/json"
	"io"
	"net/http"
	"strconv"
)

type HttpError struct {
	message string
}

func (e HttpError) Error() string {
	return e.message
}

func PostJson(url string, data interface{}, headers map[string]string) (interface{}, error) {
	json, err := json2.Marshal(data)
	if err != nil {
		return "", err
	}
	client := http.Client{
		Timeout: 2 * 1000 * 1000 * 1000,
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(json))
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	res, err := client.Do(req)

	if err != nil {
		return "", err
	}
	var response interface{}
	resData, err := io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return "", HttpError{"Invalid Status Code: " + strconv.Itoa(res.StatusCode) + "\nData:" + string(resData)}
	}
	if err != nil {
		return "", err
	}
	if resData == nil {
		return "", nil
	}
	err = json2.Unmarshal(resData, response)
	if err != nil {
		return "", err
	}
	return resData, nil
}
