package main

import (
	"errors"
	"fmt"
)

func yikes() {
	panic(errors.New("something bad happened"))
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Trapped panic: %s (%T)\n", err, err)
		}
	}()
	yikes()
}
