package main

import (
	"myapp/server"
	"myapp/utils"
	"os"
)

func main() {
	var leader string

	idPort, listPorts := utils.Parse(os.Args)

	mux := server.SetRoutes(&leader)
	go server.StartServer(mux, idPort)
	
	utils.BullyAlgorithm(idPort, listPorts, &leader)
	select {}
}
