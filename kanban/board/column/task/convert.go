package task

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/database/types"
	"github.com/mitchellh/mapstructure"
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
