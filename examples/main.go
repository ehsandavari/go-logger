package main

import (
	"context"
	"github.com/ehsandavari/go-logger"
)

func main() {
	iLogger := logger.NewLogger(true, "debug", 1, "example", "", "uuid", "1.0.0", "development", "1e56443f5a73adf5f4e26bc0f592b10a4caa282f")
	ctx := context.Background()
	ctx = context.WithValue(ctx, "RequestID", "valtest")
	ctx = context.WithValue(ctx, "TraceID", "asdfasdf24321")

	iLogger.WithBool("key", true).Debug(ctx, "asdad")
	iLogger.WithBool("1", false).Fatal(ctx, "1")
	iLogger.WithBool("2", false).Debug(ctx, "2")
	iLogger.WithBool("3", false).Debug(ctx, "3")
	iLogger.WithBool("4", false).Debug(ctx, "4")
	iLogger.WithBool("5", false).Debug(ctx, "5")
	iLogger.WithBool("6", false).Debug(ctx, "6")
	if err := iLogger.Sync(); err != nil {
		return
	}

}
