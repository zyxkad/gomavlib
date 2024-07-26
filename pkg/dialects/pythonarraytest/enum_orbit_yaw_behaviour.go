//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package pythonarraytest

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/common"
)

// Yaw behaviour during orbit flight.
type ORBIT_YAW_BEHAVIOUR = common.ORBIT_YAW_BEHAVIOUR

const (
	// Vehicle front points to the center (default).
	ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TO_CIRCLE_CENTER ORBIT_YAW_BEHAVIOUR = common.ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TO_CIRCLE_CENTER
	// Vehicle front holds heading when message received.
	ORBIT_YAW_BEHAVIOUR_HOLD_INITIAL_HEADING ORBIT_YAW_BEHAVIOUR = common.ORBIT_YAW_BEHAVIOUR_HOLD_INITIAL_HEADING
	// Yaw uncontrolled.
	ORBIT_YAW_BEHAVIOUR_UNCONTROLLED ORBIT_YAW_BEHAVIOUR = common.ORBIT_YAW_BEHAVIOUR_UNCONTROLLED
	// Vehicle front follows flight path (tangential to circle).
	ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TANGENT_TO_CIRCLE ORBIT_YAW_BEHAVIOUR = common.ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TANGENT_TO_CIRCLE
	// Yaw controlled by RC input.
	ORBIT_YAW_BEHAVIOUR_RC_CONTROLLED ORBIT_YAW_BEHAVIOUR = common.ORBIT_YAW_BEHAVIOUR_RC_CONTROLLED
	// Vehicle uses current yaw behaviour (unchanged). The vehicle-default yaw behaviour is used if this value is specified when orbit is first commanded.
	ORBIT_YAW_BEHAVIOUR_UNCHANGED ORBIT_YAW_BEHAVIOUR = common.ORBIT_YAW_BEHAVIOUR_UNCHANGED
)
