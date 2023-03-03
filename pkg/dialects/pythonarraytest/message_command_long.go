//autogenerated:yes
//nolint:revive,misspell,govet,lll
package pythonarraytest

import (
	"github.com/aler9/gomavlib/pkg/dialects/common"
)

// Send a command with up to seven parameters to the MAV. COMMAND_INT is generally preferred when sending MAV_CMD commands where param 5 and param 6 contain latitude/longitude data, as sending these in floats can result in a significant loss of precision. COMMAND_LONG is required for commands that mandate float values in params 5 and 6. The command microservice is documented at https://mavlink.io/en/services/command.html
type MessageCommandLong = common.MessageCommandLong
