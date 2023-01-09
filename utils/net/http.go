package net

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpParam struct {
	Url     string            `json:"url"`
	Method  string            `json:"method"`
	Header  map[string]string `json:"header"`
	Body    string            `json:"body"`
	Timeout int               `json:"timeout"`
}

func HttpRequest(url string, method string, body string, header map[string]string) (output []byte, err error) {

	var httpParam HttpParam
	httpParam.Url = url
	httpParam.Method = method
	httpParam.Header = header
	httpParam.Body = body
	httpParam.Timeout = 30

	res, err := httpParam.httDo()
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return response, fmt.Errorf("error while sending data to external server with status %d", res.StatusCode)
	}

	return response, nil
}

func (httpParam *HttpParam) httDo() (*http.Response, error) {
	headers := makeHeader(httpParam.Header)
	timeout := time.Duration(httpParam.Timeout) * time.Second

	switch httpParam.Method {
	case "get":
		return get(httpParam.Url, headers, timeout)
	default:
		return nil, errors.New("invalid method")
	}
}

func makeHeader(headers map[string]string) http.Header {
	result := http.Header{}
	for key, value := range headers {
		result.Add(key, value)
	}
	return result
}

func get(url string, headers http.Header, timeout time.Duration) (*http.Response, error) {
	var client = &http.Client{Timeout: timeout}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("faield http request get: %v", err)
	}

	request.Header = headers
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("faield client get: %v", err)
	}

	return response, nil
}
