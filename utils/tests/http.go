package tests

import (
	"net/http"
	"testing"
)

type Response struct {
	t            *testing.T
	expectedCode int
	expectedBody string
}

func NewResponse(t *testing.T, expectedCode int, expectedBody string) Response {
	return Response{
		t:            t,
		expectedCode: expectedCode,
		expectedBody: expectedBody,
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
	}

	return 0, nil
}
