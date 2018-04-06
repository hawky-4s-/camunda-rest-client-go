// +build unit

package unit

import (
	"github.com/hawky-4s-/camunda-rest-client-go/pkg/camunda"
	"net/http"
	"os"
	"testing"
    "github.com/hawky-4s-/camunda-rest-client-go/test/helpers"
    "github.com/hawky-4s-/camunda-rest-client-go/test/assert"
    "github.com/hawky-4s-/camunda-rest-client-go/pkg/camunda/util"
)

func Test_GetList(t *testing.T) {
	mockServer := helpers.NewMockRequest("localhost:8080", t).
		WithPath("deployment").
		WithMethod(http.MethodGet).
		ThenRespondWith().
		StatusCode(http.StatusOK).
		ContentType(helpers.ContentTypeApplicationJson).
		BodyFromFile("../../testdata/deployments/get.json").
		Build()
	defer mockServer.Close()

	client := helpers.CreateClient(mockServer.URL)

	deployments, err := client.Deployments.GetList(nil)
	assert.Nil(err, t)
    assert.NotNil(deployments, t)
    assert.Equals(deployments[0].GetId(), "6f6ea462-32c8-11e8-b6fc-0242ac110002", t)
    assert.Equals(deployments[1].GetId(), "6faf7dfc-32c8-11e8-b6fc-0242ac110002", t)
}

func Test_Delete(t *testing.T) {
	mockServer := helpers.NewMockRequest("localhost:8080", t).
		WithPath("deployment/myId").
		WithMethod(http.MethodDelete).
		ThenRespondWith().
		StatusCode(http.StatusOK).
		Build()
	defer mockServer.Close()

	client := helpers.CreateClient(mockServer.URL)

	err := client.Deployments.Delete("myId", nil)
    assert.Nil(err, t)
}

func Test_Create(t *testing.T) {
	mockServer := helpers.NewMockRequest("localhost:8080", t).
		WithPath("deployment/create").
		WithMethod(http.MethodPost).
		ThenRespondWith().
		StatusCode(http.StatusOK).
		BodyFromFile("../../testdata/deployments/create.json").
		Build()
	defer mockServer.Close()

	client := helpers.CreateClient(mockServer.URL)

	resources := make(map[string]interface{}, 0)
	file, err := os.Open("../../testdata/deployments/invoice.bpmn")
	if err != nil {
		t.Error("Could not open file: ", err)
	}
	resources["invoice.bpmn"] = file

	deploymentCreateRequest := &camunda.DeploymentCreateRequest{
		DeploymentName:           util.String("test"),
		EnableDuplicateFiltering: util.Bool(true),
		DeployChangedOnly:        util.Bool(false),
		DeploymentSource:         util.String("process application"),
		Resources:                resources,
	}

	deployment, err := client.Deployments.Create(deploymentCreateRequest)

    assert.Nil(err, t)
    assert.NotNil(deployment, t)
    assert.Equals(deployment.GetId(), "8ddebc16-3786-11e8-b6fc-0242ac110002", t)
}
