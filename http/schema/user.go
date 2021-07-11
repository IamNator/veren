package schema

import "time"

//User ...
type User struct {
	ID             uint      `json:"id"`
	FirstName      string    `json:"first_name"`
	MiddleName     string    `json:"middle_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email" validate:"required"` //this has to be unique
	Location       string    `json:"location"`
	DoB            time.Time `json:"dob"`
	ModifiedAt     time.Time `json:"modified_at"`
	HashedPassword string    `json:"hashed_password"`
}

//TableName ....
func (User) TableName() string {
	return "user"
}
