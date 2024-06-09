package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

// implements the bully algorithm for leader election
func BullyAlgorithm(idPort string, listPorts []string, leader *string) {
	var leaderMutex sync.Mutex
	timeout := 5 * time.Second

	for {
		leaderMutex.Lock()
		currentLeader := *leader
		leaderMutex.Unlock()

		// Check if there is no leader or the current leader is not alive
		if currentLeader == "" || !isAlive(currentLeader) {
			newLeader := detectLeaderFailure(idPort, listPorts)

			leaderMutex.Lock()
			*leader = newLeader
			leaderMutex.Unlock()

			if newLeader == idPort {
				announceLeader(idPort, listPorts)
			}
		}

		time.Sleep(timeout)
	}
}

// detects the failure of the current leader and elects a new leader
func detectLeaderFailure(idPort string, listPorts []string) string {
	highestPort := idPort
	for _, port := range listPorts {
		if port > idPort {
			if isAlive(port) {
				highestPort = port
			}
		}
	}

	if idPort == highestPort {
		return idPort
	}

	for _, port := range listPorts {
		if port > idPort && isAlive(port) {
			return requestLeader(port)
		}
	}

	return idPort
}

// announces the current process as the new leader
func announceLeader(idPort string, listPorts []string) {
	for _, port := range listPorts {
		if port != idPort {
			sendLeader(port, idPort)
		}
	}
	fmt.Printf("Process %s is the leader\n", idPort)
}

//  checks if the process at the given port is alive
func isAlive(port string) bool {
	resp, err := http.Get(fmt.Sprintf("http://localhost:%s/alive", port))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

// requests the leader information from the process at the given port
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

// sends the leader information to the process at the given port
func sendLeader(port string, leaderPort string) {
	resp, err := http.Post(fmt.Sprintf("http://localhost:%s/new-leader", port), "text/plain", strings.NewReader(leaderPort))
	if err != nil {
		return
	}
	defer resp.Body.Close()
}
