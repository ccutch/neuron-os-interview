package main

import (
	"log"
	"net/http"

	"neuronos"
)

func main() {
	commander := neuronos.NewCommander()
	server := &http.Server{
		Addr:    ":8080",
		Handler: handleRequests(commander),
	}
	log.Fatal(server.ListenAndServe())
}

func handleRequests(cmdr neuronos.Commander) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/execute", handleCommand(cmdr))
	return mux
}

func handleCommand(cmdr neuronos.Commander) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse request and execute command
	}
}
