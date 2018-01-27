package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Concat concatenates a bunch of strings, separated by spaces.
// It returns an empty string and an error if no strings were passed in.
func concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("no strings supplied")
	}
	return strings.Join(parts, " "), nil
}

func main() {
	result, _ := concat(os.Args[1:]...)
	fmt.Printf("Concatenated string: '%s'\n", result)
}
