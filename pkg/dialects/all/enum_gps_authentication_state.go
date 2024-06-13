//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/development"
)

// Signal authentication state in a GPS receiver.
type GPS_AUTHENTICATION_STATE = development.GPS_AUTHENTICATION_STATE

const (
	// The GPS receiver does not provide GPS signal authentication info.
	GPS_AUTHENTICATION_STATE_UNKNOWN GPS_AUTHENTICATION_STATE = development.GPS_AUTHENTICATION_STATE_UNKNOWN
	// The GPS receiver is initializing signal authentication.
	GPS_AUTHENTICATION_STATE_INITIALIZING GPS_AUTHENTICATION_STATE = development.GPS_AUTHENTICATION_STATE_INITIALIZING
	// The GPS receiver encountered an error while initializing signal authentication.
	GPS_AUTHENTICATION_STATE_ERROR GPS_AUTHENTICATION_STATE = development.GPS_AUTHENTICATION_STATE_ERROR
	// The GPS receiver has correctly authenticated all signals.
	GPS_AUTHENTICATION_STATE_OK GPS_AUTHENTICATION_STATE = development.GPS_AUTHENTICATION_STATE_OK
	// GPS signal authentication is disabled on the receiver.
	GPS_AUTHENTICATION_STATE_DISABLED GPS_AUTHENTICATION_STATE = development.GPS_AUTHENTICATION_STATE_DISABLED
)
