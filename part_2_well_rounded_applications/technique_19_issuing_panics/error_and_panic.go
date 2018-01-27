package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("can't divide by zero")

func preCheckDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return divide(a, b), nil
}

func divide(a, b int) int {
	return a / b
}

func main() {
	fmt.Println("Divide 1 by 0")
	if _, err := preCheckDivide(1, 0); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Println("Divide 2 by 0")
	divide(2, 0)
}
