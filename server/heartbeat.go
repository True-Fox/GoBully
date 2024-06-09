package server

import (
	"fmt"
	"net/http"
	"time"
)

//Heartbeat Ring-based implementation
func StartHeartBeat(idport string, listPorts []string) {
	for i := range listPorts {
		currPort := listPorts[i]
		nextPort := listPorts[(i+1)%len(listPorts)]
		go func(currPort, nextPort string) {
			for {
				resp, err := http.Get("http://localhost:" + nextPort)
				if err != nil {
					fmt.Printf("Error sending the heartbeat to port %s: %s\n", nextPort, err)
				} else {
					fmt.Printf("Heartbeat sent to port %s, status code %d\n", nextPort, resp.StatusCode)
					resp.Body.Close()
				}
				time.Sleep(10 * time.Second)
			}
		}(currPort, nextPort)
	}
}
