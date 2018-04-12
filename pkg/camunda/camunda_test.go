package camunda

import (
	"context"
	"fmt"
	"github.com/hawky-4s-/camunda-rest-client-go/test/assert"
	"testing"
	"time"
)

func Test_DefaultDateTimeFormat(t *testing.T) {
	fmt.Println(time.Now().Format(DefaultDateTimeFormat))
}

func Test_NewClient(t *testing.T) {
	endpoint := "localhost:8080/engine-rest"

	client, err := NewClient(context.Background(), endpoint)
	assert.Nil(err, t)

	assert.NotNil(client, t)

	assert.NotNil(client.httpClient, t)

	assert.NotNil(client.config, t)
	assert.Equals(client.config.UserAgent, DefaultUserAgent, t)
	assert.Equals(client.config.Endpoint.String(), endpoint, t)
}
