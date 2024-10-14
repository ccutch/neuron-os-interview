package main

import (
	"encoding/json"
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
		var (
			req neuronos.CommandRequest
			res neuronos.CommandResponse
			err error
		)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		switch req.Type {
		case "ping":
			res.Data, err = cmdr.Ping(req.Payload)
		case "sysinfo":
			res.Data, err = cmdr.GetSystemInfo()
		}

		if err != nil {
			res.Data = nil
			res.Error = err.Error()
		} else {
			res.Success = true
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
