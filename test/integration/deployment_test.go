// +build integration

package integration

import (
	"github.com/hawky-4s-/camunda-rest-client-go/pkg/camunda"
	"testing"
    "fmt"
    "github.com/hawky-4s-/camunda-rest-client-go/test/helpers"
    "github.com/hawky-4s-/camunda-rest-client-go/test/assert"
    "github.com/hawky-4s-/camunda-rest-client-go/pkg/camunda/util"
)

func Test_GetList(t *testing.T) {
	client := helpers.CreateClient(*helpers.BaseUrl)

	deployments, err := client.Deployments.GetList(nil)

	assert.Nil(err, t)
	assert.NotNil(deployments, t)
    for _, deployment := range deployments {
        fmt.Println(deployment)
    }
}

func Test_Delete(t *testing.T) {
	client := helpers.CreateClient(*helpers.BaseUrl)

	err := client.Deployments.Delete("myId", nil)

	assert.Nil(err, t)
}

func Test_Create(t *testing.T) {
    deploymentName := "test"
    deploymentSource := "process application"

    resources := make(map[string]interface{}, 0)
    resources["invoice.bpmn"] = helpers.ReadInvoiceProcess(t)

    deploymentCreateRequest := &camunda.DeploymentCreateRequest{
        DeploymentName:           util.String(deploymentName),
        EnableDuplicateFiltering: util.Bool(true),
        DeployChangedOnly:        util.Bool(false),
        DeploymentSource:         util.String(deploymentSource),
        Resources:                resources,
    }

    client := helpers.CreateClient(*helpers.BaseUrl)
    deployment, err := client.Deployments.Create(deploymentCreateRequest)

    assert.Nil(err, t)
    assert.NotNil(deployment, t)

    defer helpers.CleanupDeployment(client, deployment, t)

    assert.Equals(deployment.GetName(), deploymentName, t)
    assert.Equals(deployment.GetSource(), deploymentSource, t)
    fmt.Println(deployment)
}
