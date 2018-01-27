package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type words struct {
	sync.Mutex
	found map[string]int
}

func newWords() *words {
	return &words{found: map[string]int{}}
}

func (w *words) add(word string, n int) {
	w.Lock()
	defer w.Unlock()
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}

func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}

func main() {
	var waitGroup sync.WaitGroup
	words := newWords()
	for _, file := range os.Args[1:] {
		waitGroup.Add(1)
		go func(file string) {
			if err := tallyWords(file, words); err != nil {
				fmt.Println(err.Error())
			}
			waitGroup.Done()
		}(file)
	}
	waitGroup.Wait()
	fmt.Println("Words that appear more than once:")
	for word, count := range words.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
}
