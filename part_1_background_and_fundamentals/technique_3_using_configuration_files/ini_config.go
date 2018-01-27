package main

import (
	"log"

	"gopkg.in/gcfg.v1"
	_ "gopkg.in/gcfg.v1"
)

func main() {
	config := struct {
		Section struct {
			Enabled bool
			Path    string
		}
	}{}
	if err := gcfg.ReadFileInto(&config, "part_1_background_and_fundamentals/technique_3_using_configuration_files/config.ini"); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s %v", config.Section.Path, config.Section.Enabled)
}
