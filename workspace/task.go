package workspace

import "github.com/Sharktheone/sharkedule/kanban/task"

// Task functions
func (w *Workspace) SaveTask(task *task.Task) error {

}

func (w *Workspace) SaveTasks(tasks []*task.Task) error {

}

func (w *Workspace) GetTask(uuid string) (*task.Task, error) {

}

func (w *Workspace) DeleteTaskOnColumn(column, uuid string) error {

}

func (w *Workspace) DeleteTask(uuid string) error {

}

func (w *Workspace) MoveTask(column, uuid, toColumn string, toIndex int) error {

}

func (w *Workspace) NewTask(column, name string) (*task.Task, error) {

}

func (w *Workspace) RenameTask(task, name string) error {

}

func (w *Workspace) RemoveTagOnTask(column, uuid string) error {

}

func (w *Workspace) SetTagsOnTask(task string, tags []string) error {

}

func (w *Workspace) SetTaskDescription(task, description string) error {

}
