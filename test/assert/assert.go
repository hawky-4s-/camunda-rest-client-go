package assert

import (
    "net/http"
    "testing"
    "strings"
    "io/ioutil"
    "reflect"
)

func UrlPathIs(req *http.Request, expectedPath string, t *testing.T) {
    normalizedPath := strings.TrimPrefix(req.URL.Path, "/")
    if req.URL.RawQuery != "" {
        // if url has query params, add them
        normalizedPath += "?" + req.URL.RawQuery
    }
    if normalizedPath != expectedPath {
        t.Errorf("Expected path '%s', got '%s'.", expectedPath, normalizedPath)
    }
}

func HttpMethodIs(req *http.Request, expectedMethod string, t *testing.T) {
    if req.Method != expectedMethod {
        t.Errorf("Expected method '%s', got '%s'.", expectedMethod, req.Method)
    }
}

func AcceptHeaderIs(req *http.Request, expectedAcceptHeader string, t *testing.T) {
    if expectedAcceptHeader != "" {
        accept := req.Header.Get("Accept")
        if accept != expectedAcceptHeader {
            t.Errorf("Expected accept header to be '%s', got '%s'.", expectedAcceptHeader, accept)
        }
    }
}

func ContentTypeIs(req *http.Request, expectedContentType string, t *testing.T) {
    if expectedContentType != "" {
        contentType := req.Header.Get("Content-Type")
        if contentType != expectedContentType {
            t.Errorf("Expected content type to be '%s', got '%s'.", expectedContentType, contentType)
        }
    }
}

func HTTPBodyContains(req *http.Request, expectedBody string, t *testing.T) {
    bytes, err := ioutil.ReadAll(req.Body)
    if err != nil {
        t.Errorf("Unexpected error while reading body %s", err)
    }
    defer req.Body.Close()

    if string(bytes) != expectedBody {
        t.Errorf("Expected body to be '%s', got '%s'.", expectedBody, req.Method)
    }
}

func NoError(err error, t *testing.T, msg string) {
    if err != nil {
        t.Errorf("Error while retrieving %s. %s", msg, err)
    }
}

func Equals(given interface{}, expected interface{}, t *testing.T) {
    givenType := reflect.TypeOf(given)
    expectedType := reflect.TypeOf(expected)
    if givenType != expectedType {
        t.Errorf("Types not equal. Expected %v, got %v", expectedType, givenType)
    }
    if given != expected {
        t.Errorf("Values not equal. Expected %v, got %v", expected, given)
    }
}

func NotNil(i interface{}, t *testing.T) {
    if i == nil {
        t.Errorf("%v is nil", i)
    }
}

func Nil(i interface{}, t *testing.T) {
    if i != nil {
        t.Errorf("%v is not nil", i)
    }
}
