package utils

import (
	"fmt"
	"os"
	"strings"
)

func Parse(args []string) (string, []string) {
	if len(args) < 5 {
		fmt.Println("Usage: ./go-bully --id <server-id> --list <comma-separated list of ports>")
		os.Exit(1)
	}

	var idPort string
	var listPorts []string

	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "--id":
			if i+1 < len(args) {
				idPort = args[i+1]
				i++
			} else {
				fmt.Println("Missing value for --id")
				os.Exit(1)
			}
		case "--list":
			if i+1 < len(args) {
				listPorts = strings.Split(args[i+1], ",")
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

	return idPort, listPorts
}