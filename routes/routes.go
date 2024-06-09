package routes

import (
	"fmt"
	"io"
	"net/http"
)

func AliveHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func RequestLeaderHandler(w http.ResponseWriter, r *http.Request, leader string) {
	w.Write([]byte(leader))
}

func NewLeaderHandler(w http.ResponseWriter, r *http.Request, leader *string) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	*leader = string(body)
	fmt.Printf("New leader announced: %s\n", *leader)
}
