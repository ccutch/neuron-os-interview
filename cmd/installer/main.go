package main

import (
	"log"
	"neuronos"
)

func main() {
	cmd := neuronos.NewCommander()
	if _, err := cmd.GetSystemInfo(); err != nil {
		log.Fatal("Error", err)
	}
	log.Println("Start on boot enabled")
}
