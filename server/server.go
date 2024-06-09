package server

import (
	"myapp/routes"
	"errors"
	"fmt"
	"net/http"
	"os"
)

//Start a http server
func StartServer(mux *http.ServeMux, port string) {
	fmt.Printf("Server started at port %s\n", port)
	err := http.ListenAndServe(":"+port, mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server on port %s closed\n", port)
	} else if err != nil {
		fmt.Printf("Error starting the server on port %s: %s\n", port, err)
		os.Exit(1)
	}
}

//Set Routes of a http server
func SetRoutes(leader *string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/alive", routes.AliveHandler)
	mux.HandleFunc("/request-leader", func(w http.ResponseWriter, r *http.Request) {
		routes.RequestLeaderHandler(w, r, *leader)
	})
	mux.HandleFunc("/new-leader", func(w http.ResponseWriter, r *http.Request) {
		routes.NewLeaderHandler(w, r, leader)
	})
	return mux
}