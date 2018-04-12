//go:generate go run gen-accessors.go -v

package camunda

import (
	"context"
	"fmt"
	"github.com/hawky-4s-/camunda-rest-client-go/pkg/camunda/api"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	Version = "0.1.0-dev"

	ContentTypeJson     = "application/json"
	DefaultUserAgent    = "camunda-rest-httpClient-go/" + Version
	DefaultBaseEndpoint = "http://localhost:8080/engine-rest"
	// DefaultDateTimeFormat: see https://docs.camunda.org/manual/7.8/reference/rest/overview/date-format/
	DefaultDateTimeFormat = "2006-01-02T15:04:05.000-0700"

	DefaultMaxTasks        = 10
	DefaultPollingDelay    = 30 * time.Second
	DefaultPollingDuration = 60 * time.Second
	DefaultRetryAttempts   = 3
	DefaultRetryDuration   = 30 * time.Second
)

var log = newLogger()

type Client struct {
	httpClient          *http.Client
	config              *HttpConfiguration
	requestInterceptors []RequestInterceptor
	topicSubscriptions  map[string]*TopicSubscription

	common service

	Deployments        *DeploymentService
	ExternalTasks      *api.ExternalTaskService
	Incidents          *IncidentService
	Metrics            *MetricsService
	ProcessDefinitions *ProcessDefinitionService
	ProcessInstances   *ProcessInstanceService
}

type service struct {
	client *Client
}

func NewClient(ctx context.Context, endpoint string, opts ...ClientOption) (*Client, error) {
	var options = new(HttpConfiguration)

	// TODO: remove or improve
	//o := []ClientOption{
	//	//		WithEndpoint(DefaultBaseEndpoint),
	//	//		WithUserAgent(DefaultUserAgent),
	//}
	//if opts != nil {
	//	o = append(o, opts...)
	//}
	//
	//for _, option := range o {
	//	option.Apply(config)
	//}

	if endpoint == "" {
		endpoint = DefaultBaseEndpoint
	}
	baseEndpoint, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	if !strings.HasSuffix(baseEndpoint.Path, "/") {
		baseEndpoint.Path += "/"
	}
	options.Endpoint = baseEndpoint

	options.UserAgent = DefaultUserAgent

	httpClient, err := newHTTPClient(ctx, options)
	if err != nil {
		return nil, fmt.Errorf("dialing: %v", err)
	}

	client := &Client{
		httpClient:          httpClient,
		config:              options,
		requestInterceptors: make([]RequestInterceptor, 0),
	}

	client.common.client = client
	client.Deployments = (*DeploymentService)(&client.common)
	client.ExternalTasks = (*ExternalTaskService)(&client.common)
	client.Incidents = (*IncidentService)(&client.common)
	client.Metrics = (*MetricsService)(&client.common)
	client.ProcessDefinitions = (*ProcessDefinitionService)(&client.common)
	client.ProcessDefinitions = (*ProcessDefinitionService)(&client.common)

	return client, nil
}

// httpClient := &http.Client{
// Jar: cookieJar,
// CheckRedirect: redirectPolicyFunc,
// }
//
//req, err := http.NewRequest("GET", "http://localhost/", nil)
//req.Header.Add("Authorization","Basic " + basicAuth("username1","password123"))

//resp, err := httpClient.do(req)
//}

// config config
//type withClient *http.Client
//func (w withClient) Apply(o *HttpConfiguration) {
//    o.HTTPClient = w
//}
//
//func WithClient(httpClient *http.Client) ClientOption {
//    return withClient(httpClient)
//}
//
//type withEndpoint string
//
//func (w withEndpoint) Apply(o *HttpConfiguration) {
//	o.Endpoint = string(w)
//}
//
//func WithEndpoint(endpoint string) ClientOption {
//
//
//	return withEndpoint(url)
//}
//
//type withUserAgent string
//
//func (w withUserAgent) Apply(o *HttpConfiguration) {
//	o.UserAgent = string(w)
//}
//
//func WithUserAgent(userAgent string) ClientOption {
//	return withUserAgent(userAgent)
//}
//
//type withBasicAuth struct {
//	username string
//	password string
//}
//
//func (w withBasicAuth) Apply(o *HttpConfiguration) {
//	o.BasicAuth = basicAuth(w.username, w.password)
//}
//
//func WithBasicAuth(username, password string) ClientOption {
//	return &withBasicAuth{username, password}
//}

// CLIENT IMPLEMENTATION //

// HTTP STUFF
func (c *Client) doGet(ctx context.Context, path string, v interface{}, queryParams map[string]interface{}) (*http.Response, error) {
	var err error

	request, err := newJsonRequest(c.config.Endpoint, c.config.UserAgent, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	request.URL.RawQuery, err = addQueryParams(request.URL, queryParams)

	response, err := doRequest(ctx, c.httpClient, c.requestInterceptors, request, v)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) doPost(ctx context.Context, path string, body interface{}, v interface{}) (*http.Response, error) {
	request, err := newJsonRequest(c.config.Endpoint, c.config.UserAgent, http.MethodPost, path, body)
	if err != nil {
		return nil, err
	}

	response, err := doRequest(ctx, c.httpClient, c.requestInterceptors, request, v)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) doPut(ctx context.Context, path string, body interface{}, v interface{}) (*http.Response, error) {
	request, err := newJsonRequest(c.config.Endpoint, c.config.UserAgent, http.MethodPut, path, body)
	if err != nil {
		return nil, err
	}

	response, err := doRequest(ctx, c.httpClient, c.requestInterceptors, request, v)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) doDelete(ctx context.Context, path string, v interface{}) (*http.Response, error) {
	request, err := newJsonRequest(c.config.Endpoint, c.config.UserAgent, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	response, err := doRequest(ctx, c.httpClient, c.requestInterceptors, request, v)
	if err != nil {
		return response, err
	}

	return response, nil
}
