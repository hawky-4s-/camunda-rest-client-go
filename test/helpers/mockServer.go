package helpers

import (
	"fmt"
	"github.com/hawky-4s-/camunda-rest-client-go/test/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"
)

type MockRequest interface {
	WithBaseUrl(baseUrl string) MockRequest
	WithPath(path string) MockRequest
	WithMethod(method string) MockRequest
	WithAcceptHeader(accept string) MockRequest
	WithContentType(contentType string) MockRequest
	WithBody(body interface{}) MockRequest
	ThenRespondWith() MockResponse
}

type MockResponse interface {
	StatusCode(statusCode int) MockResponse
	ContentType(contentType string) MockResponse
	Body(body interface{}) MockResponse
	BodyFromFile(fileName string) MockResponse
	Error(err error) MockResponse
	Build() *httptest.Server
}

type mockRequest struct {
	baseUrl     string
	path        string
	method      string
	accept      string
	contentType string
	body        interface{}
}

type mockResponse struct {
	statusCode  int
	contentType string
	body        interface{}
	err         error
}

type MockBuilder struct {
	t        *testing.T
	request  mockRequest
	response mockResponse
}

func NewMockRequest(baseUrl string, t *testing.T) MockRequest {
	return &MockBuilder{
		t: t,
		request: mockRequest{
			baseUrl:     baseUrl,
			path:        "/",
			method:      http.MethodGet,
			accept:      ContentTypeApplicationJson,
			contentType: "",
			body:        nil,
		},
		response: mockResponse{
			statusCode:  http.StatusOK,
			contentType: ContentTypeApplicationJson,
			body:        nil,
			err:         nil,
		},
	}
}

func (mb *MockBuilder) WithBaseUrl(baseUrl string) MockRequest {
	mb.request.baseUrl = baseUrl
	return mb
}

func (mb *MockBuilder) WithPath(path string) MockRequest {
	mb.request.path = path
	return mb
}

func (mb *MockBuilder) WithMethod(method string) MockRequest {
	mb.request.method = method
	return mb
}

func (mb *MockBuilder) WithAcceptHeader(accept string) MockRequest {
	mb.request.accept = accept
	return mb
}

func (mb *MockBuilder) WithContentType(contentType string) MockRequest {
	mb.request.contentType = contentType
	return mb
}

func (mb *MockBuilder) WithBody(body interface{}) MockRequest {
	mb.request.body = body
	return mb
}

func (mb *MockBuilder) ThenRespondWith() MockResponse {
	return mb
}

func (mb *MockBuilder) StatusCode(statusCode int) MockResponse {
	mb.response.statusCode = statusCode
	return mb
}

func (mb *MockBuilder) ContentType(contentType string) MockResponse {
	mb.response.contentType = contentType
	return mb
}

func (mb *MockBuilder) Body(body interface{}) MockResponse {
	mb.response.body = body
	return mb
}

func (mb *MockBuilder) BodyFromFile(fileName string) MockResponse {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		mb.t.Fatalf("Unable to read file: %s. Error: %s", fileName, err)
	}

	mb.response.body = string(content)
	return mb
}

func (mb *MockBuilder) Error(err error) MockResponse {
	mb.response.err = err
	return mb
}

func (mb *MockBuilder) Build() *httptest.Server {
	request := func(r *http.Request) {
		bytes, err := httputil.DumpRequest(r, true)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", bytes)

		assert.HttpMethodIs(r, mb.request.method, mb.t)
		assert.UrlPathIs(r, mb.request.path, mb.t)
		assert.AcceptHeaderIs(r, mb.request.contentType, mb.t)
		assert.ContentTypeIs(r, mb.request.contentType, mb.t)
		//assert.HTTPBodyContains(r, mb.request.body, mb.t)
	}
	mockServer := newMockServer(
		mb.response.statusCode,
		mb.response.contentType,
		mb.response.body,
		request,
	)

	return mockServer
}

//
// Test Helpers
//
func newMockServer(statusCode int, contentType string, body interface{}, requestHandler func(r *http.Request)) *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {

		if requestHandler != nil {
			requestHandler(r)
		}

		w.WriteHeader(statusCode)
		w.Header().Set("Content-Type", contentType)
		fmt.Fprint(w, body)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}
