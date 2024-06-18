package types

import (
	"fmt"
	"path"
	"time"
)

const (
	timeFormatter     = "15:04:05"
	timeFileFormatter = "2006-01-02-15h"
	baseLogsDir       = "logs"
)

const (
	successPrefix = "success"
	warningPrefix = "warning"
	fatalPrefix   = "fatal"
	infoPrefix    = "info"
)

type LoggerConfigs struct {
	listenChan        chan *Message
	timeFormatter     string
	timeFileFormatter string
	logsDirPath       string
	binPath           string

	removeLogsDelay time.Duration

	lastCheckingTime time.Time
	checkingDelay    time.Duration
}

func BaseLoggerConfigs() (*LoggerConfigs, error) {
	binPath, err := getBinPath()
	if err != nil {
		return nil, err
	}
	return &LoggerConfigs{
		timeFormatter:     timeFormatter,
		timeFileFormatter: timeFileFormatter,
		logsDirPath:       path.Join(binPath, baseLogsDir),
		binPath:           binPath,
		removeLogsDelay:   14 * 24 * time.Hour, // 2 weeks
		lastCheckingTime:  time.Time{},
		checkingDelay:     1 * time.Hour,
	}, nil
}

func (lc *LoggerConfigs) sendToChan(msg *Message) error {
	if lc.listenChan == nil {
		return fmt.Errorf("listen channel is nil")
	}
	lc.listenChan <- msg
	return nil
}

func (lc *LoggerConfigs) SetRemoveLogDelay(timeout time.Duration) {
	lc.removeLogsDelay = timeout
}

func (lc *LoggerConfigs) SetCheckingDelay(delay time.Duration) {
	lc.checkingDelay = delay
}
