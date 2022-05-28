package commands

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func POST(url string, contentType string, body []byte) ([]byte, error) {
	reader := bytes.NewReader(body)
	resp, err := http.Post(url, contentType, reader)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		errBody, _ := getBody(resp)
		return nil, errors.New(string(errBody))
	}

	return getBody(resp)
}

func GET(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		errBody, _ := getBody(resp)
		return nil, errors.New(string(errBody))
	}

	return getBody(resp)
}

func getBody(r *http.Response) ([]byte, error) {
	return ioutil.ReadAll(r.Body)
}
