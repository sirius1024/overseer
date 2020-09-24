package models

import "time"

// Ping is the request of overseer self
type Ping struct {
	Cloud            string
	Region           string
	AvailabilityZone string
	NetworkZone      string
}

// Pong is the response of ping
type Pong struct {
	RequestInfo  Ping
	Error        string
	RequestDate  time.Time
	ResponseDate time.Time
}
