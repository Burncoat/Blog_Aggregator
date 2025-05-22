package main

import (
	"fmt"
	"log"

	"github.com/Burncoat/Blog_Aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config %v", err)
	}
	fmt.Printf("Reading config file: %+v\n", cfg)

	err = cfg.SetUser("hunter")
	if err != nil {
		log.Fatalf("couldn't set username %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config %v", err)
	}
	fmt.Printf("Reading config file again: %+v\n", cfg)
}