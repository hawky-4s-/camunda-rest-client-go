package camunda

import "github.com/hawky-4s-/camunda-rest-client-go/pkg/camunda/util"

const (
	// REST Endpoint /process-definition
	pathProcessDefinition                            = "process-definition"
	pathProcessDefinitionWithIdPattern               = pathProcessDefinition + "/%s"
	pathProcessDefinitionWithIdPatternStart          = pathProcessDefinitionWithIdPattern + "/start"
	pathProcessDefinitionKeyWithKeyPattern           = pathProcessDefinition + "/key/%s"
	pathProcessDefinitionKeyWithKeyPatternStart      = pathProcessDefinitionKeyWithKeyPattern + "/start"
	pathProcessDefinitionKeyWithTenantIdPattern      = pathProcessDefinitionKeyWithKeyPattern + "/tenant-id/%s"
	pathProcessDefinitionKeyWithTenantIdPatternStart = pathProcessDefinitionKeyWithTenantIdPattern + "/start"
)

type ProcessDefinition struct {
	Id                *string `json:"id,omitempty"`
	Key               *string `json:"key,omitempty"`
	Category          *string `json:"category,omitempty"`
	Description       *string `json:"description,omitempty"`
	Name              *string `json:"name,omitempty"`
	Version           *int    `json:"Version,omitempty"`
	Resource          *string `json:"resource,omitempty"`
	DeploymentId      *string `json:"deploymentId,omitempty"`
	Diagram           *string `json:"diagram,omitempty"`
	Suspended         *bool   `json:"suspended,omitempty"`
	TenantId          *string `json:"tenantId,omitempty"`
	VersionTag        *string `json:"versionTag,omitempty"`
	HistoryTimeToLive *int    `json:"historyTimeToLive,omitempty"` // days
}

func (p ProcessDefinition) String() string {
	return util.Stringify(p)
}

type StartProcessDefinitionBuilder struct {
	processDefinitionService ProcessDefinitionService
	id                       *string
	key                      *string
	messageName              *string
	businessKey              *string
	variables                *Variables
}

func (pdb *StartProcessDefinitionBuilder) ById(id string) *StartProcessDefinitionBuilder {
	return pdb
}

func (pdb *StartProcessDefinitionBuilder) ByKey(id string) *StartProcessDefinitionBuilder {
	return pdb
}

func (pdb *StartProcessDefinitionBuilder) ByMessage(messageName string) *StartProcessDefinitionBuilder {
	return pdb
}

func (pdb *StartProcessDefinitionBuilder) WithBusinessKey(businessKey string) *StartProcessDefinitionBuilder {
	return pdb
}

func (pdb *StartProcessDefinitionBuilder) WithVariables(variables *Variables) *StartProcessDefinitionBuilder {
	return pdb
}

func (pdb *StartProcessDefinitionBuilder) Now() *ProcessInstance {
	// TODO: call here
	return nil
}

type ProcessDefinitionService service

func (pdc ProcessDefinitionService) GetList() ([]*ProcessDefinition, error) {
	return nil, nil
}

func (pdc ProcessDefinitionService) Start() *StartProcessDefinitionBuilder {
	return nil
}

func (pdc ProcessDefinitionService) StartProcessInstanceById(id string) *ProcessDefinition {
	pdc.Start().ById("").WithBusinessKey("").WithVariables(&Variables{})

	return nil
}

func (pdc ProcessDefinitionService) StartProcessInstanceByKey(key string) *ProcessDefinition {
	return nil
}

func (pdc *ProcessDefinitionService) Delete(id string) error {
	//request, err := NewRequest(pdc.ctx, http.MethodDelete, pathProcessDefinition, nil)

	return nil
}
