package logging

var loggerConfigs = LoggerConfigs{}

func MakeGlobalChannel(channelSize uint64) chan *Message {
	loggerConfigs.ListenChan = make(chan *Message)
	return loggerConfigs.ListenChan
}
