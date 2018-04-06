package camunda

import "github.com/hawky-4s-/camunda-rest-client-go/pkg/camunda/util"

const (
	// /engine
	pathEngine = "engine"
)

type Engine struct {
	Name *string `json:"name,omitempty"`
}

func (e Engine) String() string {
	return util.Stringify(e)
}

type EngineService service

func (ec *EngineService) GetList() ([]*Engine, error) {
	var engines []*Engine
	_, err := ec.client.doGet(pathEngine, &engines, nil)
	if err != nil {
		return nil, err
	}
	return engines, nil
}
