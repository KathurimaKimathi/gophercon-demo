package mock

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

// FakeExtensionMock is a fake implementation of the Extension interface
type FakeExtensionMock struct {
	MockMakeRequestFn func(ctx context.Context, method string, path string, body interface{}) (*http.Response, error)
}

// NewFakeExtensionMock initializes the fake extension mock
func NewFakeExtensionMock() *FakeExtensionMock {
	return &FakeExtensionMock{
		MockMakeRequestFn: func(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
			msg := struct {
				Message string `json:"message"`
			}{
				Message: "success",
			}

			payload, _ := json.Marshal(msg)

			return &http.Response{
				StatusCode: http.StatusOK,
				Status:     "200 OK",
				Body:       io.NopCloser(bytes.NewBuffer(payload)),
			}, nil
		},
	}
}

// MakeRequest is a mock implementation of the method used to make http requests
func (f *FakeExtensionMock) MakeRequest(ctx context.Context, method string, path string, body interface{}) (*http.Response, error) {
	return f.MockMakeRequestFn(ctx, method, path, body)
}
