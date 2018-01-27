package main

import (
	"fmt"
	"log"
)

type ParseError struct {
	Column  int
	Line    int
	Message string
}

func (e ParseError) Error() string {
	return fmt.Sprintf("[%d:%d] %s", e.Line, e.Column, e.Message)
}

func main() {
	log.Fatal(ParseError{
		Column:  13,
		Line:    666,
		Message: "ParseError: unexpected token '.'",
	})
}
