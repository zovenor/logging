package logging

func UpdateChannel(value string, logConfigs *LogConfigs, opts ...func(*LogConfigs)) (*LogConfigs, error) {
	for _, opt := range opts {
		opt(logConfigs)
	}
	if logConfigs.sendToChan {
		err := loggerConfigs.SendMessageToChannel(&Message{
			value:   value,
			msgType: logConfigs.msgType,
		})
		if err != nil {
			return nil, err
		}
	}
	return logConfigs, nil
}

func getOptsFromInterface(values ...interface{}) (filteredValues []interface{}, opts []func(*LogConfigs)) {
	for _, iface := range values {
		if opt, ok := iface.(func(*LogConfigs)); ok {
			opts = append(opts, opt)
		} else {
			filteredValues = append(filteredValues, opt)
		}
	}
	return filteredValues, opts
}
