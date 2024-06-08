package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

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