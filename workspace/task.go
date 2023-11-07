package workspace

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/task"
	"github.com/Sharktheone/sharkedule/kanban/types"
)

// Task functions
func (w *Workspace) SaveTask(task *task.Task) error {
	return db.DB.SaveTask(w.UUID, task.Task)

}

func (w *Workspace) SaveTasks(tasks []*task.Task) error {
	var t []*types.Task
	for _, tsk := range tasks {
		t = append(t, tsk.Task)
	}
	return db.DB.SaveTasks(w.UUID, t)
}

func (w *Workspace) GetTask(uuid string) (*task.Task, error) {
	return task.Get(w.UUID, uuid)
}

func (w *Workspace) DeleteTaskOnColumn(column, uuid string) error {
	return db.DB.DeleteTaskOnColumn(w.UUID, column, uuid)
}

func (w *Workspace) DeleteTask(uuid string) error {
	return db.DB.DeleteTask(w.UUID, uuid)
}

func (w *Workspace) MoveTask(column, uuid, toColumn string, toIndex int) error {
	return db.DB.MoveTask(w.UUID, column, uuid, toColumn, toIndex)
}

func (w *Workspace) NewTask(column, name string) (*task.Task, error) {
	t, err := db.DB.NewTask(w.UUID, column, name)
	if err != nil {
		return nil, err
	}
	return &task.Task{Task: t, Workspace: w.UUID}, nil
}

func (w *Workspace) RenameTask(task, name string) error {
	return db.DB.RenameTask(w.UUID, task, name)
}

func (w *Workspace) RemoveTagOnTask(column, uuid string) error {
	return db.DB.RemoveTagOnTask(w.UUID, column, uuid)
}

func (w *Workspace) SetTagsOnTask(task string, tags []string) error {
	return db.DB.SetTagsOnTask(w.UUID, task, tags)
}

func (w *Workspace) SetTaskDescription(task, description string) error {
	return db.DB.SetTaskDescription(w.UUID, task, description)
}
