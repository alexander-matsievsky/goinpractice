package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func compress(file string) error {
	in, err := os.Open(file)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(fmt.Sprintf("%s.gz", file))
	if err != nil {
		return err
	}
	defer out.Close()
	gout := gzip.NewWriter(out)
	_, err = io.Copy(gout, in)
	gout.Close()
	return err
}

func main() {
	var group sync.WaitGroup
	for _, file := range os.Args[1:] {
		group.Add(1)
		go func(file string) {
			compress(file)
			group.Done()
		}(file)
	}
	group.Wait()
}
