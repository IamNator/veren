package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/IamNator/veren/internal/models"
)

//RegisterUserHandler ...
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var input models.RegisterNewUserRequest
	json.NewDecoder(r.Body).Decode(&input)
}
