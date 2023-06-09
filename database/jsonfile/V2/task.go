package jsonfileV2

import (
	"fmt"
	kanbandb "github.com/Sharktheone/sharkedule/kanban/v2/database"
	"github.com/Sharktheone/sharkedule/kanban/v2/types"
)

func (J *JSONFile) GetTask(uuid string) (*types.Task, error) {
	return kanbandb.GetTask(J.db.Tasks, uuid)
}

func (J *JSONFile) SaveTask(task *types.Task) error {
	if err := kanbandb.SaveTask(J.db.Tasks, task); err != nil {
		return fmt.Errorf("failed saving task: %v", err)
	}
	return J.Save()
}

func (J *JSONFile) SaveTasks(tasks []*types.Task) error {
	kanbandb.SaveTasks(J.db.Tasks, tasks)
	return J.Save()
}

func (J *JSONFile) MoveTask(column, uuid, toColumn string, toIndex int) error {
	if err := kanbandb.MoveTask(column, uuid, toColumn, toIndex); err != nil {
		return err
	}
	return J.Save()
}

func (J *JSONFile) DeleteTask(uuid string) error {
	if err := kanbandb.DeleteTask(J.db.Tasks, uuid); err != nil {
		return fmt.Errorf("")
	}
	return J.Save()
}

func (J *JSONFile) AddTagToTask(task, tag string) error {
	if err := kanbandb.AddTagToTask(J.db.Tasks, task, tag); err != nil {
		return err
	}
	return J.Save()
}

func (J *JSONFile) DeleteTaskFromColumn(column, uuid string) error {
	if err := kanbandb.DeleteTaskFromColumn(J.db.Tasks, column, uuid); err != nil {
		return err
	}
	return J.Save()
}
