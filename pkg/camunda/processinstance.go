package camunda

import "github.com/hawky-4s-/camunda-rest-client-go/pkg/camunda/util"

const (
	// REST Endpoint /process-instance
	pathProcessInstance              = "process-definition"
	pathProcessInstanceWithIdPattern = pathProcessInstance + "/%s"
)

type ProcessInstance struct {
	Id             *string `json:"id,omitempty"`
	DefinitionId   *string `json:"definitionId,omitempty"`
	BusinessKey    *string `json:"businessKey,omitempty"`
	CaseInstanceId *string `json:"caseInstanceId,omitempty"`
	Ended          *bool   `json:"ended,omitempty"`
	Suspended      *bool   `json:"suspended,omitempty"`
	TenantId       *string `json:"tenantId,omitempty"`
	Links          []*Link `json:"links,omitempty"`
}

func (pi ProcessInstance) String() string {
	return util.Stringify(pi)
}

type ProcessInstanceVariableQueryOperator string

const (
	EqualTo              ProcessInstanceVariableQueryOperator = "eq"
	NotEqualTo           ProcessInstanceVariableQueryOperator = "neq"
	GreaterThan          ProcessInstanceVariableQueryOperator = "gt"
	GreaterThanOrEqualTo ProcessInstanceVariableQueryOperator = "gteq"
	LowerThan            ProcessInstanceVariableQueryOperator = "lt"
	LowerThanOrEqualTo   ProcessInstanceVariableQueryOperator = "lteq"
	Like                 ProcessInstanceVariableQueryOperator = "like"
)

type ProcessInstanceVariableQuery struct {
	name     string
	operator ProcessInstanceVariableQueryOperator
	value    interface{}
}

type SortOption string
type ProcessInstanceSortOption SortOption

//var _, processInstanceSortOption = (ProcessInstanceSortOption)(nil)

const (
	BusinessKey   ProcessInstanceSortOption = "businessKey"
	InstanceId    ProcessInstanceSortOption = "instanceId"
	DefinitionId  ProcessInstanceSortOption = "definitionId"
	DefinitionKey ProcessInstanceSortOption = "definitionKey"
	TenantId      ProcessInstanceSortOption = "tenantId"
)

type ProcessInstanceSort struct {
	sortBy    ProcessInstanceSortOption
	sortOrder SortOrder
}

type ProcessInstanceQueryBuilder struct {
	processInstanceIds   []string
	businessKey          *string
	businessKeyLike      *string
	caseInstanceId       *string
	processDefinitionId  *string
	processDefinitionKey *string
	deploymentId         *string
	superProcessInstance *string
	subProcessInstance   *string
	superCaseInstance    *string
	subCaseInstance      *string
	active               *bool
	suspended            *bool
	incidentId           *string
	incidentType         *string
	incidentMessage      *string
	incidentMessageLike  *string
	tenantIdIn           []string
	withoutTenantId      *bool
	activityIdIn         []string

	variables *Variables
}

func (pdb *ProcessInstanceQueryBuilder) ById(id string) *ProcessInstanceQueryBuilder {
	return pdb
}

func (pdb *ProcessInstanceQueryBuilder) ByKey(id string) *ProcessInstanceQueryBuilder {
	return pdb
}

func (pdb *ProcessInstanceQueryBuilder) ByMessage(messageName string) *ProcessInstanceQueryBuilder {
	return pdb
}

func (pdb *ProcessInstanceQueryBuilder) WithBusinessKey(businessKey string) *ProcessInstanceQueryBuilder {
	return pdb
}

func (pdb *ProcessInstanceQueryBuilder) WithVariables(variables *Variables) *ProcessInstanceQueryBuilder {
	return pdb
}

type ProcessInstanceService service

func (pdc ProcessInstanceService) GetList() ([]*ProcessInstance, error) {
	return nil, nil
}

func (pdc ProcessInstanceService) GetListByQuery() *ProcessInstanceQueryBuilder {
	return nil
}

func (pdc *ProcessInstanceService) Delete(id string) error {
	//request, err := NewRequest(pdc.ctx, http.MethodDelete, pathProcessDefinition, nil)

	return nil
}
