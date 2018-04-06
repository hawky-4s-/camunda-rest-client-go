package camunda

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
    "github.com/hawky-4s-/camunda-rest-client-go/pkg/camunda/util"
)

const (
	// /deployment
	pathDeployment              = "deployment"
	pathDeploymentCreate        = pathDeployment + "/create"
	pathDeploymentWithIdPattern = pathDeployment + "/%s"
	pathDeploymentDelete        = pathDeploymentWithIdPattern
)

type Deployment struct {
	Id                                      *string                                    `json:"id,omitempty"`
	Name                                    *string                                    `json:"name,omitempty"`
	Source                                  *string                                    `json:"source,omitempty"`
	TenantId                                *string                                    `json:"tenant_id,omitempty"`
	DeploymentTime                          *time.Time                                 `json:"deployment_time,omitempty"`
	Links                                   []*Link                                    `json:"links,omitempty"`
	DeployedProcessDefinitions              map[string]*ProcessDefinition              `json:"deployedProcessDefinitions,omitempty"`
	DeployedCaseDefinitions                 map[string]*CaseDefinition                 `json:"deployedCaseDefinitions,omitempty"`
	DeployedDecisionDefinitions             map[string]*DecisionDefinition             `json:"deployedDecisionDefinitions,omitempty"`
	DeployedDecisionRequirementsDefinitions map[string]*DecisionRequirementsDefinition `json:"deployedDecisionRequirementsDefinitions,omitempty"`
}

func (d Deployment) String() string {
	return util.Stringify(d)
}

type DeploymentGetRequest struct {
	Id                                *string
	Name                              *string
	NameLike                          *string
	Source                            *string
	WithoutSource                     *string
	TenantIdIn                        []*string
	WithoutTentantId                  *bool
	IncludeDeploymentsWithoutTenantId *bool
	After                             *time.Time
	Before                            *time.Time
	SortBy                            *string // id / name / deploymentTime / tenantId
	SortOrder                         *string // asc / desc
	FirstResult                       *int
	MaxResults                        *int
}

type DeploymentCreateRequest struct {
	DeploymentName           *string
	EnableDuplicateFiltering *bool
	DeployChangedOnly        *bool
	DeploymentSource         *string
	TenantId                 *string
	Resources                map[string]interface{}
}

type DeploymentService service

func (dc *DeploymentService) GetList(queryParams *DeploymentGetRequest) ([]*Deployment, error) {
	var deployments []*Deployment
	_, err := dc.client.doGet(pathDeployment, &deployments, nil)
	if err != nil {
		return nil, err
	}
	return deployments, nil
}

func (dc *DeploymentService) Create(createDeployment *DeploymentCreateRequest) (*Deployment, error) {
	// prepare stuff
	// Prepare a form that you will submit to that URL.

	var err error
	var body bytes.Buffer

	w := multipart.NewWriter(&body)

	w.WriteField("deployment-name", *createDeployment.DeploymentName)
	w.WriteField("enable-duplicate-filtering", strconv.FormatBool(*createDeployment.EnableDuplicateFiltering))
	w.WriteField("deploy-changed-only", strconv.FormatBool(*createDeployment.DeployChangedOnly))
	w.WriteField("deployment-source", *createDeployment.DeploymentSource)
	w.WriteField("tenant-id", *createDeployment.TenantId)

	for key, resource := range createDeployment.Resources {

		var fw io.Writer

		if x, ok := resource.(io.Closer); ok {
			defer x.Close()
		}
		// Add a file
		if x, ok := resource.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return nil, err
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return nil, err
			}
		}

		if r, ok := resource.(io.Reader); ok {
			if _, err = io.Copy(fw, r); err != nil {
				return nil, err
			}
		}
	}
	//Close writer to get terminating boundary.
	if err := w.Close(); err != nil {
		return nil, err
	}

	var deployment *Deployment
	request, err := newMultiPartUploadRequest(dc.client.config.Endpoint, dc.client.config.UserAgent, http.MethodPost, pathDeploymentCreate, body, w.FormDataContentType())
	if err != nil {
		return nil, err
	}
	_, err = doRequest(dc.client.ctx, dc.client.httpClient, dc.client.requestInterceptors, request, &deployment)
	if err != nil {
		return nil, err
	}

	return deployment, nil
}

/*
* Deletes Deployment with ID
* @param id
* @param optional (nil or map[string]interface{}) with one or more of:
*   @param "cascade" (bool)
*   @param "skipCustomListeners" (bool)
*   @param "skipIoMappings" (bool)
* @return error (nil or error)
 */
func (dc *DeploymentService) Delete(id string, optional map[string]interface{}) error {
	var request *http.Request
	var err error

	if request, err = newJsonRequest(dc.client.config.Endpoint, dc.client.config.UserAgent, http.MethodDelete, fmt.Sprintf(pathDeploymentDelete, id), nil); err != nil {
		return err
	}

	request.URL.RawQuery, err = addQueryParams(request.URL, optional)
	if err != nil {
		return err
	}

    _, err = doRequest(dc.client.ctx, dc.client.httpClient, dc.client.requestInterceptors, request, nil)
    if err != nil {
        return err
    }

    return nil
}
