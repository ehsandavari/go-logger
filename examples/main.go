package main

import (
	"context"
	"github.com/ehsandavari/go-logger"
)

func main() {
	iLogger := logger.NewLogger(false, "debug", 1, "example", "", "uuid", "1.0.0", "development", "1e56443f5a73adf5f4e26bc0f592b10a4caa282f")
	ctx := context.Background()
	ctx = context.WithValue(ctx, "RequestID", "valtest")
	ctx = context.WithValue(ctx, "TraceID", "asdfasdf24321")

	iLogger.WithBool("key", true).Debug(ctx, "asdad")
	err := iLogger.Sync()
	if err != nil {
		return
	}

}
