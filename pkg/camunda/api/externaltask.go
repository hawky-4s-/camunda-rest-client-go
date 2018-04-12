package api

import (
	"fmt"
	"github.com/hawky-4s-/camunda-rest-client-go/pkg/camunda/util"
	"time"
)

const (
	// /external-task
	pathExternalTask              = "external-task"
	pathExternalTaskWithIdPattern = pathExternalTask + "/%s"
	pathExternalTaskFetchAndLock  = pathExternalTask + "/fetchAndLock"
	pathExternalTaskExtendLock    = pathExternalTaskWithIdPattern + "/extendLock"
	pathExternalTaskRetries       = pathExternalTaskWithIdPattern + "/retries"
	pathExternalTaskPriority      = pathExternalTaskWithIdPattern + "/priority"
	pathExternalTaskUnlock        = pathExternalTaskWithIdPattern + "/unlock"
	pathExternalTaskFailure       = pathExternalTaskWithIdPattern + "/failure"
	pathExternalTaskBpmnError     = pathExternalTaskWithIdPattern + "/bpmnError"
)

// ExternalTaskQueryRequest request DTO see https://docs.camunda.org/manual/develop/reference/rest/external-task/post-query/#request-body
type ExternalTaskQueryRequest struct {
	externalTaskId             string
	topicName                  string
	workerId                   string
	locked                     bool
	notLocked                  bool
	withRetriesLeft            bool
	noRetriesLeft              bool
	lockExpirationAfter        time.Time
	lockExpirationBefore       time.Time
	activityId                 string
	activityIdIn               []string
	executionId                string
	processInstanceId          string
	processDefinitionId        string
	tenantIdIn                 []string
	active                     bool
	suspended                  bool
	priorityHigherThanOrEquals int
	priorityLowerThanOrEquals  int
	// sorting
}

type ExternalTaskHandler func(externalTask ExternalTask, externalTaskService ExternalTaskService)

type ExternalTask struct {
	Id                   *string `json:"id,omitempty"`
	ActivityId           *string `json:"activityId,omitempty"`
	ActivityInstanceId   *string `json:"activityInstanceId,omitempty"`
	ErrorMessage         *string `json:"errorMessage,omitempty"`
	ExecutionId          *string `json:"executionId,omitempty"`
	LockExpirationTime   *string `json:"lockExpirationTime,omitempty"`
	ProcessDefinitionId  *string `json:"processDefinitionId,omitempty"`
	ProcessDefinitionKey *string `json:"processDefinitionKey,omitempty"`
	ProcessInstanceId    *string `json:"processInstanceId,omitempty"`
	TenantId             *string `json:"tenantId,omitempty"`
	Retries              *int    `json:"retries,omitempty"`
	Suspended            *bool   `json:"suspended,omitempty"`
	WorkerId             *string `json:"workerId,omitempty"`
	Priority             *int    `json:"priority,omitempty"`
	TopicName            *string `json:"topicName,omitempty"`
}

func (et ExternalTask) String() string {
	return util.Stringify(et)
}

type TopicSubscription struct {
	TopicName    string
	WorkerId     string
	LockDuration time.Time
	TaskHandler  *ExternalTaskHandler
}

type Topic struct {
	topicName         string
	lockDuration      time.Time // TODO: apply default value
	variables         map[string]interface{}
	businessKey       string
	processVariables  map[string]interface{}
	deserializeValues bool
}

type ExternalTaskService service

func (etc *ExternalTaskService) Subscribe(topic string, customOptions map[string]interface{}, taskHandler ExternalTaskHandler) (*TopicSubscription, error) {
	//_, ok := etc.topicSubscriptions[topic]
	//if ok {
	//	return nil, errors.New(fmt.Sprintf("TopicSubscription %s already exists.", topic))
	//}

	return nil, nil
}

func (etc *ExternalTaskService) Unsubscribe(topicSubscription *TopicSubscription) error {
	//delete(etc.topicSubscriptions, topicSubscription.TopicName)
	return nil
}

func (etc *ExternalTaskService) FetchAndLock(numOfTasks int, workerId string, topics []*Topic) error {
	return nil
}

func (etc *ExternalTaskService) Complete(taskId string, variables map[string]interface{}) error {
	return nil
}

func (etc *ExternalTaskService) ExtendLock(externalTask *ExternalTask, newDuration time.Time) error {
	return nil
}

func (etc *ExternalTaskService) Unlock(externalTask *ExternalTask) error {
	return nil
}

func (etc *ExternalTaskService) HandleFailure(externalTask *ExternalTask, errorMessage string, errorDetails string, retries int, retryTimeout time.Time) error {
	return nil
}

func (etc *ExternalTaskService) HandleBpmnError(externalTask *ExternalTask, errorCode string) error {
	return nil
}

func (etc *ExternalTaskService) ErrorDetails(externalTask *ExternalTask) (string, error) {
	return "", nil
}

func (etc *ExternalTaskService) Get(taskId string) (*ExternalTask, error) {
	var externalTask *ExternalTask
	path := fmt.Sprintf(pathExternalTaskWithIdPattern, taskId)
	_, err := etc.client.doGet(path, &externalTask, nil)
	if err != nil {
		return nil, err
	}
	return externalTask, nil
}

func (etc *ExternalTaskService) GetList(query ExternalTaskQueryRequest) ([]*ExternalTask, error) {
	//etc.doPostRequest(fmt.Sprintf(pathExternalTaskWithIdPattern, taskId))
	return nil, nil
}
