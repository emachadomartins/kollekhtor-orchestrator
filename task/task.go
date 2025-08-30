package task

import (
	"encoding/json"
	"fmt"
)

type CreationTask struct {
	ExecutionID string `json:"execution_id"`
	Query       string `json:"query"`
}

type ExecutionTask struct {
}

func Build(raw []byte) (*CreationTask, error) {
	var dataMap map[string]any

	err := json.Unmarshal(raw, &dataMap)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling JSON: %v\n", err)
	}

	if _, ok := dataMap["execution_id"]; !ok {
		return nil, fmt.Errorf("missing 'execution_id' field")
	}

	if _, ok := dataMap["query"]; !ok {
		return nil, fmt.Errorf("missing 'query' field")
	}

	return &CreationTask{
		ExecutionID: dataMap["execution_id"].(string),
		Query:       dataMap["query"].(string),
	}, nil
}
