package types

import "testing"

func TestErrorsChannel(t *testing.T) {
	empty := loggerConfigs.ErrorsChannelIsEmpty()
	if !empty {
		t.Fatal("error channel is not empty")
	}
	_ = loggerConfigs.GetErrorUpdates()
	empty = loggerConfigs.ErrorsChannelIsEmpty()
	if empty {
		t.Fatal("error channel is empty")
	}
	loggerConfigs.CloseErrorChan()
	empty = loggerConfigs.ErrorsChannelIsEmpty()
	if !empty {
		t.Fatal("error channel is not empty")
	}
}
