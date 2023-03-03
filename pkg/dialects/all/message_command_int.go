//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Send a command with up to seven parameters to the MAV, where params 5 and 6 are integers and the other values are floats. This is preferred over COMMAND_LONG when sending positional data in params 5 and 6, as it allows for greater precision when sending latitudes/longitudes. Param 5 and 6 encode positional data as scaled integers, where the scaling depends on the actual command value. NaN or INT32_MAX may be used in float/integer params (respectively) to indicate optional/default values (e.g. to use the component's current latitude, yaw rather than a specific value). The command microservice is documented at https://mavlink.io/en/services/command.html
type MessageCommandInt = common.MessageCommandInt
