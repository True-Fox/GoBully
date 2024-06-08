package main

import (
	"fmt"
	"myapp/routes"
	"myapp/server"
	"myapp/utils"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: go run main.go --id <server-id> --list <comma-separated list of ports>")
		os.Exit(1)
	}

	var idPort string
	var listPorts []string

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--id":
			if i+1 < len(os.Args) {
				idPort = os.Args[i+1]
				i++
			} else {
				fmt.Println("Missing value for --id")
				os.Exit(1)
			}
		case "--list":
			if i+1 < len(os.Args) {
				listPorts = strings.Split(os.Args[i+1], ",")
				i++
			} else {
				fmt.Println("Missing value for --list")
				os.Exit(1)
			}
		}
	}

	if idPort == "" || len(listPorts) == 0 {
		fmt.Println("Usage: go run main.go --id <server-id> --list <comma-separated list of ports>")
		os.Exit(1)
	}

	var leader string
	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.Root)
	mux.HandleFunc("/hello", routes.Hello)
	mux.HandleFunc("/alive", routes.AliveHandler)
	mux.HandleFunc("/request-leader", func(w http.ResponseWriter, r *http.Request) {
		routes.RequestLeaderHandler(w, r, leader)
	})
	mux.HandleFunc("/new-leader", func(w http.ResponseWriter, r *http.Request) {
		routes.NewLeaderHandler(w, r, &leader)
	})

	go server.StartServer(mux, idPort)
	utils.BullyAlgorithm(idPort, listPorts)
	select {}
}
