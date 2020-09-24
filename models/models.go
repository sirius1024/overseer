package models

// EndPoint that shoule be probed
type EndPoint struct {
	Region           string
	AvailabilityZone string
	NetworkZone      string
	CloudPlatform    string
	URI              string
}
