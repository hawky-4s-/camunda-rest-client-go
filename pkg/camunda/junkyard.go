package camunda

import (
	"encoding/base64"
	"net/http"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "Basic "+basicAuth("username1", "password123"))
	return nil
}

type userAgentTransport struct {
	userAgent string
	base      http.RoundTripper
}

func (t userAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	rt := t.base
	if rt == nil {
		return nil, nil
		//return nil, errors.New("transport: no Transport specified")
	}
	if t.userAgent == "" {
		return rt.RoundTrip(req)
	}
	newReq := *req
	newReq.Header = make(http.Header)
	for k, vv := range req.Header {
		newReq.Header[k] = vv
	}
	// TODO(cbro): append to existing User-Agent header?
	newReq.Header["User-Agent"] = []string{t.userAgent}
	return rt.RoundTrip(&newReq)
}
