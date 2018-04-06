// +build performance

package performance

import (
	"github.com/hawky-4s-/camunda-rest-client-go/camunda"
	"testing"
    "fmt"
    "github.com/hawky-4s-/camunda-rest-client-go/test/helpers"
    "github.com/hawky-4s-/camunda-rest-client-go/test/assert"
)

func Benchmark_GetList(b *testing.B) {
	client := helpers.CreateClient(*helpers.BaseUrl)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        client.Deployments.GetList(nil)
    }
}

func Test_Create(t *testing.T) {
    deploymentName := "test"
    deploymentSource := "process application"
    tenantId := "aTenantId"

    resources := make(map[string]interface{}, 0)
    resources["invoice.bpmn"] = helpers.ReadInvoiceProcess(t)

    deploymentCreateRequest := &camunda.DeploymentCreateRequest{
        DeploymentName:           camunda.String(deploymentName),
        EnableDuplicateFiltering: camunda.Bool(true),
        DeployChangedOnly:        camunda.Bool(false),
        DeploymentSource:         camunda.String(deploymentSource),
        TenantId:                 camunda.String(tenantId),
        Resources:                resources,
    }

    client := helpers.CreateClient(*helpers.BaseUrl)
    deployment, err := client.Deployments.Create(deploymentCreateRequest)

    assert.Nil(err, t)
    assert.NotNil(deployment, t)

    defer cleanupDeployment(client, deployment, t)

    assert.Equals(deployment.GetName(), deploymentName, t)
    assert.Equals(deployment.GetSource(), deploymentSource, t)
    fmt.Println(deployment)
}

func cleanupDeployment(client *camunda.Client, deployment *camunda.Deployment, t *testing.T) {
    deletionError := client.Deployments.Delete(deployment.GetId(), nil)

    assert.NoError(deletionError, t, fmt.Sprintf("Deployment %s was not properly deleted.", deployment.GetId()))

    deployments, _ := client.Deployments.GetList(nil)
    for _, d := range deployments {
        if d.GetId() == deployment.GetId() {
            t.FailNow()
        }
    }

}
