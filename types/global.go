package types

var loggerConfigs, _ = BaseLoggerConfigs()

func MakeGlobalChannel(channelSize uint64) chan *Message {
	loggerConfigs.listenChan = make(chan *Message, channelSize)
	return loggerConfigs.listenChan
}
