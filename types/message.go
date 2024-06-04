package types

import (
	"fmt"
	"time"
)

type MessageType uint16

const (
	InfoMessageType MessageType = iota
	WarningMessageType
	FatalMessageType
	SuccessMessageType
)

func NewMessage(value any, msgType MessageType, opts ...func(*Message)) *Message {
	msg := &Message{value: fmt.Sprint(value), msgType: msgType, time: time.Now()}
	setOpts(msg, opts...)
	return msg
}

type Message struct {
	value   string
	msgType MessageType
	time    time.Time

	sendToChan bool
}

func (m *Message) Value() string     { return m.value }
func (m *Message) SetValue(v string) { m.value = v }

func (m *Message) Type() MessageType     { return m.msgType }
func (m *Message) SetType(t MessageType) { m.msgType = t }

func (m *Message) Time() time.Time     { return m.time }
func (m *Message) SetTime(t time.Time) { m.time = t }

// Opts

func WithChanneling(msg *Message) {
	msg.sendToChan = true
}
