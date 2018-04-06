// +build unit

package unit

import (
	"net/http"
	"testing"
    "github.com/hawky-4s-/camunda-rest-client-go/test/assert"
    "github.com/hawky-4s-/camunda-rest-client-go/test/helpers"
)

func Test_GetExternalTaskById(t *testing.T) {
	mockServer := helpers.NewMockRequest("localhost:8080", t).
		WithPath("external-task/uuid").
		WithMethod(http.MethodGet).
		WithContentType(helpers.ContentTypeApplicationJson).
		WithBody(nil).
		ThenRespondWith().
		StatusCode(http.StatusOK).
		ContentType(helpers.ContentTypeApplicationJson).
		BodyFromFile("../../testdata/external_tasks/get_by_id.json").
		Build()
	defer mockServer.Close()

	client := helpers.CreateClient(mockServer.URL)
	externalTask, err := client.ExternalTasks.Get("uuid")
    assert.Nil(err, t)
    assert.NotNil(externalTask, t)
}
