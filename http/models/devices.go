package models

//AddDeviceRequest is the request made to added new devices by a user
type AddDeviceRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}



