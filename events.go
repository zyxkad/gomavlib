package gomavlib

// Event is the interface implemented by all events received through node.Events()
type Event interface {
	isEvent()
}

// EventFrame is the event fired when a frame is received
type EventFrame struct {
	// the frame
	Frame Frame

	// the node which sent the frame
	Node NodeIdentifier

	// the channel from which the frame was received
	Channel *EndpointChannel
}

func (*EventFrame) isEvent() {}

// Message returns the message inside the frame.
func (res *EventFrame) Message() Message {
	return res.Frame.GetMessage()
}

// SystemId returns the frame system id.
func (res *EventFrame) SystemId() byte {
	return res.Frame.GetSystemId()
}

// ComponentId returns the frame component id.
func (res *EventFrame) ComponentId() byte {
	return res.Frame.GetComponentId()
}

// EventParseError is the event fired when a parse error occurs
type EventParseError struct {
	// the error
	Error error

	// the channel used to send the frame
	Channel *EndpointChannel
}

func (*EventParseError) isEvent() {}

// EventChannelOpen is the event fired when a channel is opened
type EventChannelOpen struct {
	Channel *EndpointChannel
}

func (*EventChannelOpen) isEvent() {}

// EventChannelClose is the event fired when a channel is closed
type EventChannelClose struct {
	Channel *EndpointChannel
}

func (*EventChannelClose) isEvent() {}

// EventNodeAppear is the event fired when a new node is detected
type EventNodeAppear struct {
	Node NodeIdentifier
}

func (*EventNodeAppear) isEvent() {}

// EventNodeDisappear is the event fired when a node disappears (i.e. times out)
type EventNodeDisappear struct {
	Node NodeIdentifier
}

func (*EventNodeDisappear) isEvent() {}