//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

// Get the current mode.
// This should be emitted on any mode change, and broadcast at low rate (nominally 0.5 Hz).
// It may be requested using MAV_CMD_REQUEST_MESSAGE.
type MessageCurrentMode struct {
	// Standard mode.
	StandardMode MAV_STANDARD_MODE `mavenum:"uint8"`
	// A bitfield for use for autopilot-specific flags
	CustomMode uint32
	// The custom_mode of the mode that was last commanded by the user (for example, with MAV_CMD_DO_SET_STANDARD_MODE, MAV_CMD_DO_SET_MODE or via RC). This should usually be the same as custom_mode. It will be different if the vehicle is unable to enter the intended mode, or has left that mode due to a failsafe condition. 0 indicates the intended custom mode is unknown/not supplied
	IntendedCustomMode uint32
}

// GetID implements the message.Message interface.
func (*MessageCurrentMode) GetID() uint32 {
	return 436
}