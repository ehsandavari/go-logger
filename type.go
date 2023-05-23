package logger

import (
	"go.uber.org/zap/zapcore"
	"io"
	"net/http"
)

type httpRequest struct {
	request *http.Request
}

func newHttpRequest(request *http.Request) *httpRequest {
	return &httpRequest{
		request: request,
	}
}

func (r *httpRequest) MarshalLogObject(objectEncoder zapcore.ObjectEncoder) error {
	objectEncoder.AddString("url", r.request.URL.String())
	objectEncoder.AddString("method", r.request.Method)
	objectEncoder.AddInt64("contentLength", r.request.ContentLength)
	objectEncoder.AddString("version", r.request.Proto)
	if r.request.Body != nil {
		body, err := io.ReadAll(r.request.Body)
		if err != nil {
			return err
		}
		objectEncoder.AddString("body", string(body))
	}
	objectEncoder.OpenNamespace("header")
	for name, values := range r.request.Header {
		for _, value := range values {
			objectEncoder.AddString(name, value)
		}
	}
	return nil
}

type httpResponse struct {
	response *http.Response
}

func newHttpResponse(response *http.Response) *httpResponse {
	return &httpResponse{
		response: response,
	}
}

func (r *httpResponse) MarshalLogObject(objectEncoder zapcore.ObjectEncoder) error {
	objectEncoder.AddString("status", r.response.Status)
	objectEncoder.AddInt64("contentLength", r.response.ContentLength)
	objectEncoder.AddString("version", r.response.Proto)
	if r.response.Body != nil {
		body, err := io.ReadAll(r.response.Body)
		if err != nil {
			objectEncoder.AddString("ioReadAllError", err.Error())
		}
		if len(body) != 0 {
			objectEncoder.AddString("body", string(body))
		}
	}
	objectEncoder.OpenNamespace("header")
	for name, values := range r.response.Header {
		for _, value := range values {
			objectEncoder.AddString(name, value)
		}
	}
	return nil
}
