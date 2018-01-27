package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "hello_cli"
	app.Usage = "Print hello world"
	app.Version = "0.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "World",
			Usage: "Who to say hello to.",
		},
	}
	app.Action = func(context *cli.Context) error {
		name := context.GlobalString("name")
		fmt.Printf("Hello, %s!", name)
		return nil
	}
	app.Run(os.Args)
}
