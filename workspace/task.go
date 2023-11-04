package workspace

import "github.com/Sharktheone/sharkedule/kanban/task"

// Task functions
func (w *Workspace) SaveTask(workspace string, task *task.Task) error {

}

func (w *Workspace) SaveTasks(workspace string, tasks []*task.Task) error {

}

func (w *Workspace) GetTask(workspace, uuid string) (*task.Task, error) {

}

func (w *Workspace) DeleteTaskOnColumn(workspace, column, uuid string) error {

}

func (w *Workspace) DeleteTask(workspace, uuid string) error {

}

func (w *Workspace) MoveTask(workspace, column, uuid, toColumn string, toIndex int) error {

}

func (w *Workspace) NewTask(workspace, column, name string) (*task.Task, error) {

}

func (w *Workspace) RenameTask(workspace, task, name string) error {

}

func (w *Workspace) RemoveTagOnTask(workspace, column, uuid string) error {

}

func (w *Workspace) SetTagsOnTask(workspace, task string, tags []string) error {

}

func (w *Workspace) SetTaskDescription(workspace, task, description string) error {

}
