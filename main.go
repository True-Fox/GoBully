package main

import (
	"fmt"
	"myapp/server"
	"myapp/utils"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: ./go-bully --id <server-id> --list <comma-separated list of ports>")
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
	mux := server.SetRoutes(&leader)
	go server.StartServer(mux, idPort)
	utils.BullyAlgorithm(idPort, listPorts, &leader)
	select {}
}
