package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func BullyAlgorithm(idPort string, listports []string) {
	leader := ""
	timeout := 5 * time.Second
	for {
		if leader == "" {
			leader = detectLeaderFailure(idPort, listports)
			if leader == idPort{
				announceLeader(idPort, listports)
			}
		}

		time.Sleep(timeout)
	}

}

func detectLeaderFailure(idPort string, listports[] string) string {
	highestPort := idPort
	for _,port := range listports{
		if port > idPort {
			alive := isAlive(port)
			if alive {
				highestPort = port
			}
		}
	}

	if idPort == highestPort {
		return idPort
	}

	for _, port := range listports {
		if port>idPort && isAlive(port){
			return requestLeader(port)
		}
	}

	return ""
}

func announceLeader(idPort string, listports[] string){
	for _, port := range listports {
		if port != idPort {
			sendLeader(port, idPort)
		}
	}
	fmt.Printf("Process %s is the leader\n", idPort)
}

func isAlive(port string) bool {
	resp, err := http.Get(fmt.Sprintf("http://localhost:%s/alive", port))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func requestLeader(port string) string {
	resp, err := http.Get(fmt.Sprintf("http://localhost:%s/request-leader", port))
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

func sendLeader(port string, leaderPort string){
	resp, err := http.Post(fmt.Sprintf("http://localhost:%s/new-leader", port), "text/plain", strings.NewReader(leaderPort))
	if err != nil {
		return
	}
	defer resp.Body.Close()
}