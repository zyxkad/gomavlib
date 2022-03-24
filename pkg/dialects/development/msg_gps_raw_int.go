//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

// The global position, as returned by the Global Positioning System (GPS). This is
// NOT the global position estimate of the system, but rather a RAW sensor value. See message GLOBAL_POSITION_INT for the global position estimate.
type MessageGpsRawInt struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// GPS fix type.
	FixType GPS_FIX_TYPE `mavenum:"uint8"`
	// Latitude (WGS84, EGM96 ellipsoid)
	Lat int32
	// Longitude (WGS84, EGM96 ellipsoid)
	Lon int32
	// Altitude (MSL). Positive for up. Note that virtually all GPS modules provide the MSL altitude in addition to the WGS84 altitude.
	Alt int32
	// GPS HDOP horizontal dilution of position (unitless * 100). If unknown, set to: UINT16_MAX
	Eph uint16
	// GPS VDOP vertical dilution of position (unitless * 100). If unknown, set to: UINT16_MAX
	Epv uint16
	// GPS ground speed. If unknown, set to: UINT16_MAX
	Vel uint16
	// Course over ground (NOT heading, but direction of movement) in degrees * 100, 0.0..359.99 degrees. If unknown, set to: UINT16_MAX
	Cog uint16
	// Number of satellites visible. If unknown, set to UINT8_MAX
	SatellitesVisible uint8
	// Altitude (above WGS84, EGM96 ellipsoid). Positive for up.
	AltEllipsoid int32 `mavext:"true"`
	// Position uncertainty.
	HAcc uint32 `mavext:"true"`
	// Altitude uncertainty.
	VAcc uint32 `mavext:"true"`
	// Speed uncertainty.
	VelAcc uint32 `mavext:"true"`
	// Heading / track uncertainty
	HdgAcc uint32 `mavext:"true"`
	// Yaw in earth frame from north. Use 0 if this GPS does not provide yaw. Use UINT16_MAX if this GPS is configured to provide yaw and is currently unable to provide it. Use 36000 for north.
	Yaw uint16 `mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageGpsRawInt) GetID() uint32 {
	return 24
}