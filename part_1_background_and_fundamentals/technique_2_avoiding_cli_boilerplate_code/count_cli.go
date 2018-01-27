package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "count_cli"
	app.Usage = "Count up and down."
	app.Commands = []cli.Command{{
		Action: func(context *cli.Context) error {
			start := context.Int("stop")
			if start <= 0 {
				fmt.Println("Stop cannot be negative.")
			}
			for i := 1; i <= start; i++ {
				fmt.Println(i)
			}
			return nil
		},
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "stop, s",
				Usage: "Value to count up to",
				Value: 10,
			},
		},
		Name:      "up",
		ShortName: "u",
		Usage:     "Count Up",
	}, {
		Action: func(context *cli.Context) error {
			start := context.Int("start")
			if start < 0 {
				fmt.Println("Start cannot be negative.")
			}
			for i := start; i >= 0; i-- {
				fmt.Println(i)
			}
			return nil
		},
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "start, s",
				Usage: "Start counting down from",
				Value: 10,
			},
		},
		Name:      "down",
		ShortName: "d",
		Usage:     "Count Down",
	}}
	app.Run(os.Args)
}
