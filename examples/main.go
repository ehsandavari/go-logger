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
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJFUzUxMiIsImtpZCI6IjllYzgzYmUxLWVmMjctNDA2OC1iYTUzLTk5MWY0MWI0OWYxYiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJXRUIiLCJhdXRob3JpdGllcyI6bnVsbCwiY2xpZW50IjoidW5raG93biIsImV4cCI6MTY4MjQzOTQ0MCwiaXNzIjoiIiwianRpIjoiNzEzOTE5MDAtNTA1Ny00N2E1LWI5ODktM2Q0MGI5MmRhOTdkIiwibW9iaWxlIjoiOTg5MjE1NTgwNjkwIiwibW9iaWxlX25vIjoiRWtyVlVHdmptRkhmb0MzYVhSQnZBNFE4WDBZbDNoNnJ1WWpGQ1E9PSIsIm5iZiI6MTY4MjQzMTY0MCwibm9uX290cCI6IiIsInBlcm1pc3Npb25zIjpudWxsLCJzY29wZSI6IndlYiIsInN1YiI6IjIxNDcwODJhLTA3OWEtNDA4Yy05OGE4LWFhYzExMmZiZjM3ZSJ9.ATg2NbeJoIOvUXts5n3sJwy-AAK_DN3tF02F7djtLopAIOQ3jxgfdcfXWJ4sTWo7zba4cNt0OjDgbMPOw6nDDMLlAYLYt6rOznMpUCCqrzxe5NJ32taa9fBhsOrk0NIGTMcqf0q5E4Ey90avuBCnD4iEUxzprvBr1To0xEzNHwGHdrw4")

	iLogger.WithBool("key", true).WithError(errors.New("sadadsadsasdasd")).WithHttpRequest(req).Error(ctx, "asdad")

	res, _ := http.DefaultClient.Do(req)

	iLogger.WithBool("key", true).WithError(errors.New("sadadsadsasdasd")).WithHttpRequest(req).Error(ctx, "asdad")
	iLogger.WithBool("key", true).WithError(errors.New("sadadsadsasdasd")).WithHttpResponse(res).Error(ctx, "asdad")
	iLogger.WithBool("key", true).WithError(errors.New("sadadsadsasdasd")).WithHttpRequest(req).WithHttpResponse(res).Error(ctx, "asdad")
	name := NewName("1", "2", "3")
	iLogger.WithAny("Sname", name).WithBool("key", true).WithError(errors.New("sadadsadsasdasd")).WithHttpResponse(res).Debug(ctx, "asdad")
	iLogger.Warn(ctx, "asdasdasdasd")

	if err := iLogger.Sync(); err != nil {
	}
}
