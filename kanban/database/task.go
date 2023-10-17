package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/types"
)

func NewTask(tasks *[]*types.Task, column *types.Column, name string) *types.Task {
	task := types.NewTask(name)
	column.Tasks = append(column.Tasks, task.UUID)
	*tasks = append(*tasks, task)
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

func MoveTask(column, toColumn *types.Column, uuid string, toIndex int) error {
	var colFound bool

	for index, t := range column.Tasks {
		if t == uuid {
			column.Tasks = append(column.Tasks[:index], column.Tasks[index+1:]...)
			colFound = true
			break
		}
	}

	if !colFound {
		return fmt.Errorf("task %s not found in column %s", uuid, column)
	}

	if toIndex < len(toColumn.Tasks) {
		for i := range toColumn.Tasks {
			if i == toIndex {
				toColumn.Tasks = append(toColumn.Tasks[:i], append([]string{uuid}, toColumn.Tasks[i:]...)...)
			}
		}
	} else {
		toColumn.Tasks = append(toColumn.Tasks, uuid)
	}
	return nil
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

func AddTagToTask(tasks []*types.Task, task, tag string) error { // TODO: This should not be an task array
	for _, t := range tasks {
		if t.UUID == task {
			t.Tags = append(t.Tags, tag)
			return nil
		}
	}
	return fmt.Errorf("error while adding tag %s to task %s not found", tag, task)
}

func RemoveTagOnTask(task *types.Task, tag string) error {
	for index, t := range task.Tags {
		if t == tag {
			task.Tags = append(task.Tags[:index], task.Tags[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while deleting tag %s not found on task %s", tag, task.UUID)
}

func SetTagsOnTask(task *types.Task, tags []string) {
	task.Tags = tags
}

func SetTaskDescription(task *types.Task, description string) {
	task.Description = description
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

func RenameTask(task *types.Task, name string) {
	task.Name = name
}
