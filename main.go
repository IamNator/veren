package main

import (
	"net/http"

	"github.com/IamNator/veren/internal/config"
	"github.com/IamNator/veren/internal/logger"
	"github.com/IamNator/veren/internal/router"
)

func main() {

	port := config.Config.PORT
	
	router := router.New()
	logger.Logger.Println("server running on port :" + port)

	http.ListenAndServe(":"+port, router)
}
