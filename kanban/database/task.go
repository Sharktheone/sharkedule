package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/database/db"
	types2 "github.com/Sharktheone/sharkedule/kanban/types"
)

func NewTask(column *types2.Column, name string) *types2.Task {
	task := types2.NewTask(name)
	column.Tasks = append(column.Tasks, task.UUID)
	return task
}

func GetTask(tasks []*types2.Task, uuid string) (*types2.Task, error) {
	for _, t := range tasks {
		if t.UUID == uuid {
			return t, nil
		}
	}
	return nil, fmt.Errorf("task with uuid %s does not exist", uuid)
}

func SaveTask(tasks []*types2.Task, task *types2.Task) error {
	for i, t := range tasks {
		if t.UUID == task.UUID {
			tasks[i] = task
			return nil
		}
	}
	return fmt.Errorf("task with uuid %s does not exist", task.UUID)
}

func SaveTasks(tasks []*types2.Task, tasksToSave []*types2.Task) {
	tasks = tasksToSave
}

func MoveTask(column, uuid, toColumn string, toIndex int) error {
	var (
		col   *types2.Column
		toCol *types2.Column
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

func RemoveTagFromTask(task *types2.Task, tag string) error {
	for index, t := range task.Tags {
		if t == tag {
			task.Tags = append(task.Tags[:index], task.Tags[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while removing tag %s not found on task %s", tag, task.UUID)
}

func DeleteTask(tasks []*types2.Task, uuid string) error {
	for index, t := range tasks {
		if t.UUID == uuid {
			tasks = append(tasks[:index], tasks[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while deleting task %s not found", uuid)
}

func AddTagToTask(tasks []*types2.Task, task, tag string) error { // TODO: This should not be an task array
	for _, t := range tasks {
		if t.UUID == task {
			t.Tags = append(t.Tags, tag)
			return nil
		}
	}
	return fmt.Errorf("error while adding tag %s to task %s not found", tag, task)
}

func DeleteTagOnTask(task *types2.Task, tag string) error {
	for index, t := range task.Tags {
		if t == tag {
			task.Tags = append(task.Tags[:index], task.Tags[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while deleting tag %s not found on task %s", tag, task.UUID)
}

func DeleteTaskOnColumn(column *types2.Column, uuid string) error {
	for index, t := range column.Tasks {
		if t == uuid {
			column.Tasks = append(column.Tasks[:index], column.Tasks[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("failed to delete task on column: task %v not found on column %v", uuid, column.UUID)
}

func RenameTask(task *types2.Task, name string) {
	task.Name = name
}
