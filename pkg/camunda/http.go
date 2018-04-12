package camunda

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func newHTTPClient(ctx context.Context, options *HttpConfiguration) (*http.Client, error) {
	// TODO: create meaningful HTTPClient object
	if options.HTTPClient != nil {
		return options.HTTPClient, nil
	}

	client := &http.Client{}
	return client, nil
}

type RequestInterceptor func(req *http.Request) error

func applyRequestInterceptors(req *http.Request, interceptors []RequestInterceptor) error {
	for _, interceptor := range interceptors {
		if err := interceptor(req); err != nil {
			return err
		}
	}
	return nil
}

func closeBody(res *http.Response) {
	if res == nil || res.Body == nil {
		return
	}
	res.Body.Close()
}

// checkResponse returns an error (of type *Error) if the response
// status code is not 2xx.
func checkResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	_, err := ioutil.ReadAll(res.Body)
	if err == nil {
		return fmt.Errorf("response error - statuscode: %d, error: %s", res.StatusCode, err)
	}
	//jerr := new(errorReply)
	//err = json.Unmarshal(slurp, jerr)
	//if err == nil && jerr.Error != nil {
	//    if jerr.Error.Code == 0 {
	//        jerr.Error.Code = res.StatusCode
	//    }
	//    jerr.Error.Body = string(slurp)
	//    return jerr.Error
	//    }
	//}
	//return &Error{
	//    Code:   res.StatusCode,
	//    Body:   string(slurp),
	//    Header: res.Header,
	//}
	return nil
}

// do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
//
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
func doRequest(ctx context.Context, httpClient *http.Client, requestInterceptors []RequestInterceptor, request *http.Request, v interface{}) (*http.Response, error) {
	request = request.WithContext(ctx)

	// allow user to modify request with interceptors
	applyRequestInterceptors(request, requestInterceptors)

	resp, err := httpClient.Do(request)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// If the error type is *url.Error, sanitize its URL before returning.
		//if e, ok := err.(*url.Error); ok {
		//	if url, err := url.Parse(e.URL); err == nil {
		//		e.URL = sanitizeURL(url).String()
		//		return nil, e
		//	}
		//}

		return nil, err
	}
	defer resp.Body.Close()

	err = checkResponse(resp)
	if err != nil {
		// Even though there was an error, we still return the response
		// in case the caller wants to inspect it further.
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			//if c.debug {
			//    httputil.DumpResponse(resp, true)
			//}
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return resp, err
}

func newJsonRequest(endpoint *url.URL, userAgent string, method string, path string, body interface{}) (*http.Request, error) {
	return newRequest(endpoint, userAgent, method, path, body, ContentTypeJson)
}

func newMultiPartUploadRequest(endpoint *url.URL, userAgent string, method string, path string, body interface{}, contentType string) (*http.Request, error) {
	return newRequest(endpoint, userAgent, method, path, body, contentType)
}

func newRequest(endpoint *url.URL, userAgent string, method string, path string, body interface{}, contentType string) (*http.Request, error) {
	if !strings.HasSuffix(endpoint.Path, "/") {
		return nil, fmt.Errorf("endpoint must have a trailing slash, but %q does not", endpoint)
	}
	requestUrl, err := endpoint.Parse(path)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	// do not encode body, just use it
	if buffer, ok := body.(io.ReadWriter); ok {
		buf = buffer
	} else if contentType == "application/json" && body != nil {
		// encode body to json
		buf = new(bytes.Buffer)
		encoder := json.NewEncoder(buf)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(method, requestUrl.String(), buf)
	if err != nil {
		return nil, err
	}

	//if body != nil {
	request.Header.Set("Content-Type", contentType)
	//}

	request.Header.Set("Accept", "application/json")

	if userAgent != "" {
		request.Header.Set("User-Agent", userAgent)
	}

	return request, nil
}

func addQueryParams(url *url.URL, params map[string]interface{}) (string, error) {
	q := url.Query()
	for name, v := range params {
		var value string
		switch v.(type) {
		case string:
			value = v.(string)
		case bool:
			value = strconv.FormatBool(v.(bool))
		case int:
			value = strconv.Itoa(v.(int))
		case float32:
			value = strconv.FormatFloat(v.(float64), 'E', -1, 32)
		case float64:
			value = strconv.FormatFloat(v.(float64), 'E', -1, 64)
		default:
			return "", fmt.Errorf("Unable to determine type of %s", v)
		}

		q.Add(name, value)
	}

	return q.Encode(), nil
}
