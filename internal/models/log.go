package models

import "github.com/IamNator/veren/internal/schema"

//LogEndpointResponse is the response when user access the log of a device
type LogEndpointResponse struct {
	schema.Log
}
