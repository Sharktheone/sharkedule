package task

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"sharkedule/database/types"
)

func (t *Task) Convert() (*types.Task, error) {
	var task *types.Task
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  &task,
		TagName: "json",
	})
	if err != nil {
		return nil, fmt.Errorf("failed creating task decoder: %v", err)
	}
	if err := decoder.Decode(t); err != nil {
		return nil, fmt.Errorf("failed decoding task: %v", err)
	}
	return task, nil
}
