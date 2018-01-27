package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	ErrRejected = errors.New("the request was rejected")
	ErrTimeout  = errors.New("the request timed out")
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func SendRequest(req string) (string, error) {
	switch random.Int() % 3 {
	case 0:
		return req, nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}
}

func main() {
	response, err := SendRequest("Hello")
	for err == ErrTimeout {
		fmt.Println("Timeout. Retrying.")
		response, err = SendRequest("Hello")
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}
