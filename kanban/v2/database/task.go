package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/v2/task"
)

func DeleteTask(tasks []*task.Task, uuid string) error {
	for index, t := range tasks {
		if t.UUID == uuid {
			tasks = append(tasks[:index], tasks[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while deleting task %s not found", uuid)
}

func RemoveTagFromTask(task *task.Task, tag string) error {
	for index, t := range task.Tags {
		if t == tag {
			task.Tags = append(task.Tags[:index], task.Tags[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while removing tag %s not found on task %s", tag, task.UUID)
}

func GetTask(tasks []*task.Task, uuid string) (*task.Task, error) {
	for _, t := range tasks {
		if t.UUID == uuid {
			return t, nil
		}
	}
	return nil, fmt.Errorf("task with uuid %s does not exist", uuid)
}

func SaveTask(tasks []*task.Task, task *task.Task) error {
	for i, t := range tasks {
		if t.UUID == task.UUID {
			tasks[i] = task
			return nil
		}
	}
	return fmt.Errorf("task with uuid %s does not exist", task.UUID)
}

func SaveTasks(tasks []*task.Task, tasksToSave []*task.Task) {
	tasks = tasksToSave
}
