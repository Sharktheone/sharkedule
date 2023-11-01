package jsonfile

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/database"
	"github.com/Sharktheone/sharkedule/kanban/types"
)

func (J *JSONFile) NewTask(column, name string) (*types.Task, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	col, err := J.GetColumn(column)
	if err != nil {
		return nil, err
	}
	t := kanbandb.NewTask(&ws.Tasks, col, name)
	return t, nil
}

func (J *JSONFile) GetTask(uuid string) (*types.Task, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	return kanbandb.GetTask(ws.Tasks, uuid)
}

func (J *JSONFile) SaveTask(task *types.Task) error {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	if err := kanbandb.SaveTask(ws.Tasks, task); err != nil {
		return fmt.Errorf("failed saving task: %v", err)
	}
	return J.Save()
}

func (J *JSONFile) SaveTasks(tasks []*types.Task) error {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	kanbandb.SaveTasks(ws.Tasks, tasks)
	return J.Save()
}

func (J *JSONFile) MoveTask(column, uuid, toColumn string, toIndex int) error {
	var (
		col   *types.Column
		toCol *types.Column
		err   error
	)

	col, err = J.GetColumn(column)
	if err != nil {
		return err
	}

	if column == toColumn {
		toCol = col
	} else {
		toCol, err = J.GetColumn(toColumn)
		if err != nil {
			return err
		}
	}
	if err := kanbandb.MoveTask(col, toCol, uuid, toIndex); err != nil {
		return err
	}
	return J.Save()
}

func (J *JSONFile) DeleteTask(uuid string) error {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	if err := kanbandb.DeleteTask(ws.Tasks, uuid); err != nil {
		return fmt.Errorf("")
	}
	return J.Save()
}

func (J *JSONFile) AddTagToTask(task, tag string) error {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	if err := kanbandb.AddTagToTask(ws.Tasks, task, tag); err != nil {
		return err
	}
	return J.Save()
}

func (J *JSONFile) RemoveTagOnTask(task, tag string) error {
	t, err := J.GetTask(task)
	if err != nil {
		return err
	}
	if err := kanbandb.RemoveTagOnTask(t, tag); err != nil {
		return err
	}
	return J.Save()
}

func (J *JSONFile) SetTagsOnTask(task string, tags []string) error {
	t, err := J.GetTask(task)
	if err != nil {
		return err
	}
	kanbandb.SetTagsOnTask(t, tags)
	return J.Save()
}

func (J *JSONFile) SetTaskDescription(task, description string) error {
	t, err := J.GetTask(task)
	if err != nil {
		return err
	}
	kanbandb.SetTaskDescription(t, description)
	return J.Save()
}

func (J *JSONFile) DeleteTaskOnColumn(column, uuid string) error {
	col, err := J.GetColumn(column)
	if err != nil {
		return err
	}
	if err := kanbandb.DeleteTaskOnColumn(col, uuid); err != nil {
		return err
	}
	return J.Save()
}

func (J *JSONFile) RenameTask(task, name string) error {
	t, err := J.GetTask(task)
	if err != nil {
		return err
	}
	kanbandb.RenameTask(t, name)

	return J.Save()
}
