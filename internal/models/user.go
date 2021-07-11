package models

import "time"

//UserInformationResponse is the response to send the user information
type UserInformationResponse struct {
	ID         uint      `json:"id"`
	FirstName  string    `json:"first_name"`
	MiddleName string    `json:"middle_name"`
	Email      string    `json:"email" validate:"required"`
	LastName   string    `json:"last_name"`
	Location   string    `json:"location"`
	DoB        time.Time `json:"dob"`
	ModifiedAt time.Time `json:"modified_at"`
}

//RegisterNewUserRequest ...
type RegisterNewUserRequest struct {
	FirstName  string    `json:"first_name" validate:"required"`
	MiddleName string    `json:"middle_name" validate:"required"`
	Email      string    `json:"email" validate:"required"`
	LastName   string    `json:"last_name" validate:"required"`
	Location   string    `json:"location" validate:"required"`
	DoB        time.Time `json:"dob" validate:"required"`
	PassWord   string    `json:"password" validate:"required"`
}

//UserLoginRequest ...
type UserLoginRequest struct {
	Email string `json:"email" validate:"required"`
}
