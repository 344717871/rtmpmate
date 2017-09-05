package Message

import ()

type Message struct {
	Header
	Payload []byte
}

type Header struct {
	Type      byte
	Length    int // 3 bytes
	Timestamp uint32
	StreamID  uint32 // 3 bytes
}

func New() (*Message, error) {
	var m Message
	return &m, nil
}
