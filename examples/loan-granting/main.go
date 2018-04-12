package loan_granting

import (
	"context"

	"github.com/hawky-4s-/camunda-rest-client-go/pkg/camunda"
	"github.com/hawky-4s-/camunda-rest-client-go/pkg/camunda/util"
	"os"
)

func main() {
	// initialize client
	client, err := camunda.NewClient(context.Background(), camunda.DefaultBaseEndpoint, nil)
	if err != nil {
		panic(err)
	}

	// deploy example process
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

	topicSubscription, err := client.ExternalTasks
	Subscribe(
		"creditScoreChecker",
		nil,
		func(et camunda.ExternalTask, ets camunda.ExternalTaskService) {

			defaultScore := et.GetVariable("defaultScore").(int)

			creditScores := []int{defaultScore, 9, 1, 4, 10}

			et.SetVariable("creditScores", creditScores)

			ets.Complete()
		})

	if err != nil {
		panic(err)
	}
}
