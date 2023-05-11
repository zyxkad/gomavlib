//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ardupilotmega

import (
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/common"
)

// Set gimbal manager pitch and yaw angles (high rate message). This message is to be sent to the gimbal manager (e.g. from a ground station) and will be ignored by gimbal devices. Angles and rates can be set to NaN according to use case. Use MAV_CMD_DO_GIMBAL_MANAGER_PITCHYAW for low-rate adjustments that require confirmation.
type MessageGimbalManagerSetPitchyaw = common.MessageGimbalManagerSetPitchyaw
