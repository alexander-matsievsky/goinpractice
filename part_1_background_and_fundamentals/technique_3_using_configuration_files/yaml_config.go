package main

import (
	"log"

	"github.com/kylelemons/go-gypsy/yaml"
	_ "github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	config, err := yaml.ReadFile("part_1_background_and_fundamentals/technique_3_using_configuration_files/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(config.Get("path"))
	log.Print(config.GetBool("enabled"))
}
