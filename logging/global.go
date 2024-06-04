package logging

var loggerConfigs = LoggerConfigs{}

func MakeGlobalChannel(channelSize uint64) chan *Message {
	loggerConfigs.ListenChan = make(chan *Message, channelSize)
	return loggerConfigs.ListenChan
}
