package schema

//Log ...
type Log struct {
	ID       uint `json:"id"`
	DeviceID uint `json:"device_id"`
	State    int  `json:"state"` // saves the state of the device

	ModifiedAt string `json:"modified_at"`
}

//TableName ....
func (Log) TableName() string {
	return "log"
}
