package main

import (
	"github.com/ehsandavari/go-logger"
)

func main() {
	iLogger := logger.NewLogger(true, "debug", 1, "example", "", "uuid", "1.0.0", "development", "1e56443f5a73adf5f4e26bc0f592b10a4caa282f")
	iLogger.Infof("Failed to fetch URL: %s", "url")
	err := iLogger.Sync()
	if err != nil {
		iLogger.Error(err)
	}
}
