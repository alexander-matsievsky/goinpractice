package main

import (
	"encoding/json"
	"log"
	"os"
)

type configuration struct {
	Enabled bool   `json:"enabled"`
	Path    string `json:"path"`
}

func main() {
	config := configuration{}
	file, err := os.Open("part_1_background_and_fundamentals/technique_3_using_configuration_files/config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s %v", config.Path, config.Enabled)
}
