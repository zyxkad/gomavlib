package main

import (
	"log"

	"github.com/aler9/gomavlib"
	"github.com/aler9/gomavlib/pkg/dialect"
	"github.com/aler9/gomavlib/pkg/message"
)

// this is a custom message.
// It must be prefixed with "Message" and implement the message.Message interface.
type MessageCustom struct {
	Param1 uint8
	Param2 uint8
	Param3 uint32
}

func (*MessageCustom) GetID() uint32 {
	return 304
}

func main() {
	// create a custom dialect from messages
	dialect := &dialect.Dialect{3, []message.Message{
		&MessageCustom{},
	}}

	// create a node which
	// - communicates with a serial port
	// - understands our custom dialect
	// - writes messages with given system id
	node, err := gomavlib.NewNode(gomavlib.NodeConf{
		Endpoints: []gomavlib.EndpointConf{
			gomavlib.EndpointSerial{
				Device: "/dev/ttyUSB0",
				Baud:   57600,
			},
		},
		Dialect:     dialect,
		OutVersion:  gomavlib.V2, // change to V1 if you're unable to communicate with the target
		OutSystemID: 10,
	})
	if err != nil {
		panic(err)
	}
	defer node.Close()

	// print every message we receive
	for evt := range node.Events() {
		if frm, ok := evt.(*gomavlib.EventFrame); ok {
			log.Printf("received: id=%d, %+v\n", frm.Message().GetID(), frm.Message())
		}
	}
}
