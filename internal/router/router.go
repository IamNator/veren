package router

import (
	"github.com/IamNator/veren/internal/controllers"
	"github.com/gorilla/mux"
)

//New routes all http requests within the app
func New() *mux.Router {
	mx := mux.NewRouter()

	mx.HandleFunc("/", controllers.HomePageHandler)

	return mx
}
