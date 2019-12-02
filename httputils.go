package main

import (
	"net/http"
	"strings"
)

// TODO: receive method (GET, POST, PUT, PATCH, etc)
// and pass to request handler
func makeRequest(method, url, body string) (error, *http.Request) {
	// Need to strip \n from passed in string to
	// prevent illegal character error

	request, err := http.NewRequest(method, strings.TrimSuffix(url, "\n"), nil)
	if err != nil {
		return err, nil
	}

	return nil, request

}

func getRequestResponse(request *http.Request) (error, *http.Response) {
	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return err, nil
	}

	return nil, resp
}

func GetResponse(method, url, body string) (error, *http.Response) {
	err, request := makeRequest(method, url, body)

	if err != nil {
		return err, nil
	}

	return getRequestResponse(request)
}
