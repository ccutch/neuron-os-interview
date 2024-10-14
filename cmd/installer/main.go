package main

import (
	"log"
	"neuronos"
)

func main() {
	if err := neuronos.EnableStartOnBoot(); err != nil {
		log.Fatal("Error", err)
	}
	log.Println("Start on boot enabled")
}
