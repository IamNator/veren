package main

import (
	"net/http"

	"github.com/IamNator/veren/http/config"
	"github.com/IamNator/veren/http/logger"
)

func main() {

	port := config.Config.PORT
	logger.Logger.Println("server running on port :" + port)

	http.ListenAndServe(":"+port, nil)
}
