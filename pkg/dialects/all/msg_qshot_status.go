//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Information about the shot operation.
type MessageQshotStatus struct {
	// Current shot mode.
	Mode MAV_QSHOT_MODE `mavenum:"uint16"`
	// Current state in the shot. States are specific to the selected shot mode.
	ShotState uint16
}

// GetID implements the msg.Message interface.
func (*MessageQshotStatus) GetID() uint32 {
	return 60020
}