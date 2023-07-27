//autogenerated:yes
//nolint:revive,misspell,govet,lll
package paparazzi

// This message is emitted as response to SCRIPT_REQUEST_LIST by the MAV to get the number of mission scripts.
type MessageScriptCount struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Number of script items in the sequence
	Count uint16
}

// GetID implements the message.Message interface.
func (*MessageScriptCount) GetID() uint32 {
	return 183
}