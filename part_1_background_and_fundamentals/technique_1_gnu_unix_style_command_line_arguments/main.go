package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	Name    string `default:"World" description:"A name to say hello to." long:"name" short:"n"`
	Spanish bool   `description:"Use Spanish Language" long:"spanish" short:"s"`
}

func main() {
	flags.Parse(&opts)

	if opts.Spanish == true {
		fmt.Printf("Hola %s!\n", opts.Name)
	} else {
		fmt.Printf("Hello %s!\n", opts.Name)
	}
}
