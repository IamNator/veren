package schema

//Device record the devices owned by a user
type Device struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ModifiedAt  string `json:"modified_at"`
}

//TableName returns table name
func (Device) TableName() string {
	return "device"
}
