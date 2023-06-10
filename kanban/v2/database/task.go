package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/v2/types"
)

func NewTask(column *types.Column, name string) *types.Task {
	task := types.NewTask(name)
	column.Tasks = append(column.Tasks, task.UUID)
	return task
}

func GetTask(tasks []*types.Task, uuid string) (*types.Task, error) {
	for _, t := range tasks {
		if t.UUID == uuid {
			return t, nil
		}
	}
	return nil, fmt.Errorf("task with uuid %s does not exist", uuid)
}

func SaveTask(tasks []*types.Task, task *types.Task) error {
	for i, t := range tasks {
		if t.UUID == task.UUID {
			tasks[i] = task
			return nil
		}
	}
	return fmt.Errorf("task with uuid %s does not exist", task.UUID)
}

func SaveTasks(tasks []*types.Task, tasksToSave []*types.Task) {
	tasks = tasksToSave
}

func MoveTask(column, uuid, toColumn string, toIndex int) error {
	var (
		col   *types.Column
		toCol *types.Column
		err   error
	)

	col, err = db.DB.GetColumn(column)
	if err != nil {
		return err
	}

	if column == toColumn {
		toCol = col
	} else {
		toCol, err = db.DB.GetColumn(toColumn)
		if err != nil {
			return err
		}
	}

	var colFound bool

	for index, t := range col.Tasks {
		if t == uuid {
			col.Tasks = append(col.Tasks[:index], col.Tasks[index+1:]...)
			colFound = true
			break
		}
	}

	if !colFound {
		return fmt.Errorf("task %s not found in column %s", uuid, column)
	}

	if toIndex > len(toCol.Tasks) {
		for i := range toCol.Tasks {
			if i == toIndex {
				toCol.Tasks = append(toCol.Tasks[:i], append([]string{uuid}, toCol.Tasks[i+1:]...)...)
			}
		}
	} else {
		toCol.Tasks = append(toCol.Tasks, uuid)
	}
	return nil
}

func RemoveTagFromTask(task *types.Task, tag string) error {
	for index, t := range task.Tags {
		if t == tag {
			task.Tags = append(task.Tags[:index], task.Tags[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while removing tag %s not found on task %s", tag, task.UUID)
}

func DeleteTask(tasks []*types.Task, uuid string) error {
	for index, t := range tasks {
		if t.UUID == uuid {
			tasks = append(tasks[:index], tasks[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while deleting task %s not found", uuid)
}

func AddTagToTask(tasks []*types.Task, task, tag string) error {
	for _, t := range tasks {
		if t.UUID == task {
			t.Tags = append(t.Tags, tag)
			return nil
		}
	}
	return fmt.Errorf("error while adding tag %s to task %s not found", tag, task)
}

func DeleteTaskOnColumn(column *types.Column, uuid string) error {
	for index, t := range column.Tasks {
		if t == uuid {
			column.Tasks = append(column.Tasks[:index], column.Tasks[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("failed to delete task on column: task %v not found on column %v", uuid, column.UUID)
}
