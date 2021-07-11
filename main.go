package main

import (
	"net/http"

	"github.com/IamNator/veren/internal/config"
	"github.com/IamNator/veren/internal/logger"
)

func main() {

	port := config.Config.PORT
	logger.Logger.Println("server running on port :" + port)

	http.ListenAndServe(":"+port, nil)
}
