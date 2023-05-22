package main

import (
	"context"
	"errors"
	"github.com/ehsandavari/go-logger"
	"net/http"
	"strings"
)

func main() {
	iLogger := logger.NewLogger(false, "debug", 1, "example", "", "uuid", "1.0.0", "development", "1e56443f5a73adf5f4e26bc0f592b10a4caa282f")
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestId", "valtest")
	ctx = context.WithValue(ctx, "traceId", "asdfasdf24321")

	url := "https://google.com?sad=asd&fsdg=asd"

	payload := strings.NewReader("{\n\t\"isPin\": true\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJFUzUxMiIsImtpZCI6IjllYzgzYmUxLWVmMjctNDA2OC1iYTUzLTk5MWY0MWI0OWYxYiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJXRUIiLCJhdXRob3JpdGllcyI6bnVsbCwiY2xpZW50IjoidW5raG93biIsImV4cCI6MTY4MjQzOTQ0MCwiaXNzIjoiIiwianRpIjoiNzEzOTE5MDAtNTA1Ny00N2E1LWI5ODktM2Q0MGI5MmRhOTdkIiwibW9iaWxlIjoiOTg5MjE1NTgwNjkwIiwibW9iaWxlX25vIjoiRWtyVlVHdmptRkhmb0MzYVhSQnZBNFE4WDBZbDNoNnJ1WWpGQ1E9PSIsIm5iZiI6MTY4MjQzMTY0MCwibm9uX290cCI6IiIsInBlcm1pc3Npb25zIjpudWxsLCJzY29wZSI6IndlYiIsInN1YiI6IjIxNDcwODJhLTA3OWEtNDA4Yy05OGE4LWFhYzExMmZiZjM3ZSJ9.ATg2NbeJoIOvUXts5n3sJwy-AAK_DN3tF02F7djtLopAIOQ3jxgfdcfXWJ4sTWo7zba4cNt0OjDgbMPOw6nDDMLlAYLYt6rOznMpUCCqrzxe5NJ32taa9fBhsOrk0NIGTMcqf0q5E4Ey90avuBCnD4iEUxzprvBr1To0xEzNHwGHdrw4")
	iLogger.WithBool("key", true).WithError(errors.New("sadadsadsasdasd")).WithHttpRequest(req).Debug(ctx, "asdad")

	res, _ := http.DefaultClient.Do(req)

	iLogger.WithBool("key", true).WithError(errors.New("sadadsadsasdasd")).WithHttpResponse(res).Debug(ctx, "asdad")
	if err := iLogger.Sync(); err != nil {
		return
	}

}
