package main

import (
	"fmt"
	"time"
)

func send2(ch chan<- string) chan<- bool {
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				println("Done")
				close(ch)
				return
			default:
				ch <- "hello"
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	return done
}

func main() {
	msg := make(chan string)
	until := time.After(5 * time.Second)
	done := send2(msg)
	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			done <- true
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}
