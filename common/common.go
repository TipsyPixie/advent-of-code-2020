package common

import (
	"bufio"
	"net/http"
)

type Input interface {
	ReadLine()
	ReadAll()
	Close()
}

type urlInput struct {
	response *http.Response
	bufferedReader *bufio.Reader
}

func (input *urlInput) Close() error {
	err := input.response.Body.Close()
	if err != nil {
		return err
	}
	return nil
}

func FromUrl(url string) (*urlInput, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return &urlInput{
		response: response,
		bufferedReader: bufio.NewReader(response.Body),
	}, nil
}
