package main

import (
	"log"
	"nandcli/internal/config"

)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	config.Current = cfg

}
