package httpRequest

import (
	"bytes"
	"net/http"
)

func HttpRequest(method, url, contentType, jsonString string, headerContents map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader([]byte(jsonString)))
	switch contentType {
	case "json":
		req.Header.Add("Content-Type", "application/json; charset=utf-8")
	case "form":
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	default:
	}

	if headerContents != nil {
		for k, v := range headerContents {
			req.Header.Add(k, v)
		}
	}
	res, err := http.DefaultClient.Do(req)
	return res, err
}
