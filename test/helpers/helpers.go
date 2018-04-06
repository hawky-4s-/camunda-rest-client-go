package helpers

import (
	"context"
	"github.com/hawky-4s-/camunda-rest-client-go/pkg/camunda"
    "github.com/hawky-4s-/camunda-rest-client-go/pkg/camunda/util"
	"testing"
    "os"
    "flag"
    "fmt"
    "github.com/hawky-4s-/camunda-rest-client-go/test/assert"
)

var BaseUrl = flag.String("baseUrl", "http://192.168.99.100:8080/engine-rest", "Endpoint of Camunda BPM Platform Engine-REST")

const (
	ContentTypeApplicationJson = "application/json"
)

func CreateClient(endpoint string) *camunda.Client {
	client, _ := camunda.NewClient(context.Background(), endpoint)
	return client
}

func ReadInvoiceProcess(t *testing.T) (*os.File) {
    return ReadFile("../fixtures/deployments/invoice.bpmn", t)
}

func ReadFile(fileName string, t *testing.T) (*os.File) {
    file, err := os.Open(fileName)
    if err != nil {
        t.Fatalf("Could not open file: %s", err)
        t.FailNow()
    }
    return file
}

func DeployInvoiceProcess(client *camunda.Client, t *testing.T) (*camunda.Deployment, error) {
    return DeployProcess(client, t)
}

func DeployProcess(client *camunda.Client, t *testing.T) (*camunda.Deployment, error) {
    resources := make(map[string]interface{}, 0)
    resources["invoice.bpmn"] = ReadInvoiceProcess(t)

    deploymentCreateRequest := &camunda.DeploymentCreateRequest{
        DeploymentName:           util.String("test"),
        EnableDuplicateFiltering: util.Bool(true),
        DeployChangedOnly:        util.Bool(false),
        DeploymentSource:         util.String("process application"),
        Resources:                resources,
    }

    deployment, err := client.Deployments.Create(deploymentCreateRequest)
    if err != nil {
        return nil, err
    }

    return deployment, err
}

func CleanupDeployment(client *camunda.Client, deployment *camunda.Deployment, t *testing.T) {
    deletionError := client.Deployments.Delete(deployment.GetId(), nil)

    assert.NoError(deletionError, t, fmt.Sprintf("Deployment %s was not properly deleted.", deployment.GetId()))

    deployments, _ := client.Deployments.GetList(nil)
    for _, d := range deployments {
        if d.GetId() == deployment.GetId() {
            t.FailNow()
        }
    }

}
