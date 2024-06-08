package server

import (
	"fmt"
	"net/http"
	"time"
)

func StartHeartBeat(idport string, listPorts []string) {
	for _, port := range listPorts {
		go func(port string) {
			for {
				resp, err := http.Get("http://localhost:" + port)
				if err != nil {
					fmt.Printf("Error sending the heartbeat to port %s: %s\n", port, err)
				} else {
					fmt.Printf("Heartbeat sent to port %s, status code %d\n", port, resp.StatusCode)
					resp.Body.Close()
				}
				time.Sleep(10 * time.Second)
			}
		}(port)
	}
}