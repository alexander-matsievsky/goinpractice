package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func read(r io.Reader) <-chan []byte {
	out := make(chan []byte)
	go func() {
		for {
			data := make([]byte, 1024)
			n, _ := r.Read(data)
			if n > 0 {
				out <- data
			}
		}
	}()
	return out
}

func main() {
	stdin := read(os.Stdin)
	timeout := time.After(10 * time.Second)
	for {
		select {
		case data := <-stdin:
			os.Stdout.Write(data)
		case <-timeout:
			fmt.Println("Timed out")
			os.Exit(0)
		}
	}
}
