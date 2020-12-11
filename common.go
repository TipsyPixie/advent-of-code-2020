package advent_of_code_2020

import (
	"bufio"
	"io/ioutil"
	"os"
	"testing"
)

type Input interface {
	ReadLine() (string, bool, error)
	ReadAll() (string, error)
	Close() error
}

type fileInput struct {
	file    *os.File
	scanner *bufio.Scanner
}

// to make sure fileInput implements Input
var _ Input = (*fileInput)(nil)

func FromFile(path string) (*fileInput, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &fileInput{
		file:    file,
		scanner: bufio.NewScanner(file),
	}, nil
}

func (input *fileInput) ReadLine() (string, bool, error) {
	if input.scanner.Scan() {
		return input.scanner.Text(), true, nil
	}
	return "", false, input.scanner.Err()
}

func (input *fileInput) ReadAll() (string, error) {
	fileContents, err := ioutil.ReadAll(input.file)
	if err != nil {
		return "", err
	}
	return string(fileContents), nil
}

func (input *fileInput) Close() error {
	err := input.file.Close()
	if err != nil {
		return err
	}
	return nil
}

func CommonTest(t *testing.T, f func(string) (int, error)) {
	answer, err := f("./input.txt")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(answer)
}
