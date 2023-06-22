package main

import (
	"context"
	"errors"
	"github.com/ehsandavari/go-logger"
	"net/http"
	"strings"
)

type Name struct {
	Name  string
	Name1 string
	Name2 string
}

func NewName(name string, name1 string, name2 string) *Name {
	return &Name{Name: name, Name1: name1, Name2: name2}
}

func main() {
	iLogger := logger.NewLogger(false, "debug", 1, "example", "test", "uuid", "1.0.0", "development", "1e56443f5a73adf5f4e26bc0f592b10a4caa282f", false, logger.WithElk("localhost:50000", 5), logger.WithConsole())
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestId", "valtest")
	ctx = context.WithValue(ctx, "traceId", "asdfasdf24321")

	url := "https://github.com/ehsandavari"

	payload := strings.NewReader("{\n\t\"isPin\": true\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer asdkdsaijojfaspdjfasd;f")

	iLogger.WithBool("key", true).WithError(errors.New("sadadsadsasdasd")).WithHttpRequest(req).Error(ctx, "asdad")

	res, _ := http.DefaultClient.Do(req)

	iLogger.WithBool("key", true).WithError(errors.New("sadadsadsasdasd")).WithHttpResponse(res).Error(ctx, "asdad")
	iLogger.WithAny("Sname", NewName("1", "2", "3")).WithBool("key", true).WithError(errors.New("sadadsadsasdasd")).WithHttpResponse(res).Debug(ctx, "asdad")
	iLogger.Warn(ctx, "asdasdasdasd")

	if err := iLogger.Sync(); err != nil {
	}
}
