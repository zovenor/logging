package logging

import "fmt"

type LoggerConfigs struct {
	ListenChan chan *Message
}

func (lc *LoggerConfigs) SendMessageToChannel(msg *Message) error {
	if lc.ListenChan == nil {
		return fmt.Errorf("listen channel is nil")
	}
	lc.ListenChan <- msg
	return nil
}

type LogConfigs struct {
	sendToChan bool
	msgType    MessageType
}

func WithChanneling() func(*LogConfigs) {
	return func(logC *LogConfigs) {
		logC.sendToChan = true
	}
}

func WithType(msgType MessageType) func(*LogConfigs) {
	return func(logC *LogConfigs) {
		logC.msgType = msgType
	}
}
