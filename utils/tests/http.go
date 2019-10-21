package tests

import (
	"net/http"
	"testing"
)

type Response struct {
	t            *testing.T
	expectedCode int
	expectedBody string
	handleBody   *func([]byte)
}

func NewResponse(t *testing.T, expectedCode int, expectedBody string) Response {
	return Response{
		t:            t,
		expectedCode: expectedCode,
		expectedBody: expectedBody,
		handleBody:   nil,
	}
}

func NewResponseWithHandleBody(t *testing.T, expectedCode int, handleBody func([]byte)) Response {
	return Response{
		t:            t,
		expectedCode: expectedCode,
		handleBody:   &handleBody,
	}
}

func (r Response) WriteHeader(statusCode int) {
	if r.expectedCode != statusCode {
		r.t.Errorf("Expected status code %d got %d instead", r.expectedCode, statusCode)
	}
}

func (r Response) Header() http.Header {
	return http.Header{}
}

func (r Response) Write(data []byte) (int, error) {
	if r.expectedBody != "" && r.expectedBody != string(data) {
		r.t.Error()

		return 0, nil
	} else if r.handleBody != nil {
		fun := *r.handleBody
		fun(data)
	}

	return 0, nil
}
