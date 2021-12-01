//autogenerated:yes
//nolint:golint,misspell,govet,lll
package uavionix

// Current motion information from a designated system
type MessageFollowTarget struct {
	// Timestamp (time since system boot).
	Timestamp uint64
	// bit positions for tracker reporting capabilities (POS = 0, VEL = 1, ACCEL = 2, ATT + RATES = 3)
	EstCapabilities uint8
	// Latitude (WGS84)
	Lat int32
	// Longitude (WGS84)
	Lon int32
	// Altitude (MSL)
	Alt float32
	// target velocity (0,0,0) for unknown
	Vel [3]float32
	// linear target acceleration (0,0,0) for unknown
	Acc [3]float32
	// (0 0 0 0 for unknown)
	AttitudeQ [4]float32
	// (0 0 0 for unknown)
	Rates [3]float32
	// eph epv
	PositionCov [3]float32
	// button states or switches of a tracker device
	CustomState uint64
}

// GetID implements the msg.Message interface.
func (*MessageFollowTarget) GetID() uint32 {
	return 144
}