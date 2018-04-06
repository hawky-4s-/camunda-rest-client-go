package camunda

import (
	"time"
    "fmt"
    "github.com/hawky-4s-/camunda-rest-client-go/pkg/camunda/util"
)

const (
	// /incident
	pathIncident              = "incident"
	pathIncidentCount         = pathIncident + "/count"
	pathIncidentWithIdPattern = pathIncident + "/%s"
)

type Incident struct {
	Id                  *string    `json:"id,omitempty"`
	ProcessDefinitionId *string    `json:"processDefinitionId,omitempty"`
	ProcessInstanceId   *string    `json:"processInstanceId,omitempty"`
	ExecutionId         *string    `json:"executionId,omitempty"`
	IncidentTimestamp   *time.Time `json:"incidentTimestamp,omitempty"`
	IncidentType        *string    `json:"incidentType,omitempty"` // https://docs.camunda.org/manual/7.8/user-guide/process-engine/incidents/#incident-types
	ActivityId          *string    `json:"activityId,omitempty"`
	CauseIncidentId     *string    `json:"causeIncidentId,omitempty"`
	RootCauseIncidentId *string    `json:"rootCauseIncidentId,omitempty"`
	Configuration       *string    `json:"configuration,omitempty"`
	TenantId            *string    `json:"tenantId,omitempty"`
	IncidentMessage     *string    `json:"incidentMessage,omitempty"`
	JobDefinitionId     *string    `json:"jobDefinitionId,omitempty"`
}

func (i Incident) String() string {
	return util.Stringify(i)
}

type IncidentService service

func (ic *IncidentService) GetList() ([]*Incident, error) {
	var incidents []*Incident
	_, err := ic.client.doGet(pathIncident, &incidents, nil)
	if err != nil {
		return nil, err
	}
	return incidents, nil
}

func (ic *IncidentService) GetCount() (int, error) {
    var result CountResult
    _, err := ic.client.doGet(pathIncident, &result, nil)
    if err != nil {
        return 0, err
    }
    return result.Count, nil
}

func (ic *IncidentService) Delete(id string) error {
    _, err := ic.client.doDelete(fmt.Sprintf(pathIncidentWithIdPattern, id), nil)
    if err != nil {
        return err
    }
    return nil
}
