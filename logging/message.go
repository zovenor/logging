package logging

type MessageType uint16

const (
	PrintMessageType MessageType = iota
	InfoMessageType
	WarningMessageType
	FatalMessageType
	SuccessMessageType
)

func NewMessage(value string, msgType MessageType) *Message {
	return &Message{value: value, msgType: msgType}
}

type Message struct {
	value   string
	msgType MessageType
}

func (m *Message) Value() string { return m.value }

func (m *Message) SetValue(v string) { m.value = v }

func (m *Message) Type() MessageType { return m.msgType }

func (m *Message) SetType(t MessageType) { m.msgType = t }
