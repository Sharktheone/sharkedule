package access

import (
	"errors"
	"fmt"
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/element"
	"github.com/Sharktheone/sharkedule/field"
	"github.com/Sharktheone/sharkedule/kanban/board"
	"github.com/Sharktheone/sharkedule/kanban/column"
	"github.com/Sharktheone/sharkedule/kanban/tag"
	"github.com/Sharktheone/sharkedule/kanban/task"
	"github.com/Sharktheone/sharkedule/kanban/types"
	"github.com/Sharktheone/sharkedule/user/access/workspaceaccess"
	"github.com/Sharktheone/sharkedule/workspace"
)

func (a *Access) GetWorkspace(uuid string) (*workspace.Workspace, error) {
	for _, w := range a.Workspaces {
		if w.UUID == uuid {
			ws, error := db.DB.GetWorkspace(uuid)
			if error != nil {
				return nil, error
			}
			return &workspace.Workspace{Workspace: ws}, nil
		}
	}
	return nil, errors.New("workspace not found")
}

// Board functions
func (a *Access) CreateBoard(workspace, name string) (*board.Board, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.Permissions.CreateBoards {
		return nil, fmt.Errorf("no permissions to create board in workspace %s", workspace)
	}

	brd, err := db.DB.CreateBoard(workspace, name)
	if err != nil {
		return nil, err
	}
	return &board.Board{Board: brd, Workspace: workspace}, nil
}

func (a *Access) SaveBoard(workspace string, board *board.Board) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.UpdateBoards {
		return fmt.Errorf("no permissions to update board in workspace %s", workspace)
	}

	brd, err := ws.Board(board.UUID)
	if err == nil {
		if !brd.Permissions.Update {
			return fmt.Errorf("no permissions to update board %s", board.UUID)
		}
	}

	return db.DB.SaveBoard(workspace, board.Board)

}

func (a *Access) SaveBoards(workspace string, boards []*board.Board) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.UpdateBoards {
		return fmt.Errorf("no permissions to update board in workspace %s", workspace)
	}

	for _, b := range boards {
		brd, err := ws.Board(b.UUID)
		if err == nil { //error is nil, when we have
			if !brd.Permissions.Update {
				return fmt.Errorf("no permissions to update board %s", b.UUID)
			}
		}

	}

	var b []*types.Board
	for _, brd := range boards {
		b = append(b, brd.Board)
	}

	return db.DB.SaveBoards(workspace, b)

}

func (a *Access) GetBoard(workspace, uuid string) (*board.Board, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllBoards {
		_, err := ws.Board(uuid) //when it is in the slice, the person has access to it
		if err != nil {
			return nil, err
		}
	}

	b, err := db.DB.GetBoard(workspace, uuid)
	if err != nil {
		return nil, err
	}
	return &board.Board{Board: b, Workspace: workspace}, nil

}

func (a *Access) GetAllBoards(workspace string) ([]*board.Board, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllBoards {
		var uuids []string
		for _, b := range ws.Boards {
			uuids = append(uuids, b.UUID)
		}
		b, err := db.DB.GetBoards(workspace, uuids)
		if err != nil {
			return nil, err
		}

		var boards []*board.Board
		for _, brd := range b {
			boards = append(boards, &board.Board{Board: brd, Workspace: workspace})
		}
		return boards, nil
	}

	b, err := db.DB.GetAllBoards(workspace)
	if err != nil {
		return nil, err
	}

	var boards []*board.Board
	for _, brd := range b {
		boards = append(boards, &board.Board{Board: brd, Workspace: workspace})
	}
	return boards, nil

}

func (a *Access) GetBoards(workspace string, uuids []string) ([]*board.Board, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllBoards {
		for _, uuid := range uuids {
			_, err := ws.Board(uuid) //when the board is NOT in the slice / the user can't access them, this returns an error
			if err != nil {
				return nil, err
			}
		}
	}

	b, err := db.DB.GetBoards(workspace, uuids)
	if err != nil {
		return nil, err
	}

	var boards []*board.Board
	for _, brd := range b {
		boards = append(boards, &board.Board{Board: brd, Workspace: workspace})
	}
	return boards, nil

}

func (a *Access) GetAllBoardNames(workspace string) ([]*types.NameList, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllBoards {
		var uuids []string
		for _, b := range ws.Boards {
			uuids = append(uuids, b.UUID)
		}
		return db.DB.GetBoardNames(workspace, uuids)
	}

	return db.DB.GetAllBoardNames(workspace)

}

func (a *Access) GetBoardNames(workspace string, uuids []string) (names []*types.NameList, err error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllBoards {
		for _, uuid := range uuids {
			_, err := ws.Board(uuid) //when the board is NOT in the slice / the user can't access them, this returns an error
			if err != nil {
				return nil, err
			}
		}
	}

	return db.DB.GetBoardNames(workspace, uuids)
}

func (a *Access) DeleteBoard(workspace, uuid string) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.DeleteBoards {
		return fmt.Errorf("no permissions to delete board in workspace %s", workspace)
	}

	brd, err := ws.Board(uuid)
	if err == nil {
		if !brd.Permissions.Delete {
			return fmt.Errorf("no permissions to delete board %s", uuid)
		}
	}

	return db.DB.DeleteBoard(workspace, uuid)

}

// Column functions
func (a *Access) SaveColumn(workspace string, column *column.Column) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.UpdateColumns { //TODO check for override permissions
		return fmt.Errorf("no permissions to update column in workspace %s", workspace)
	}

	col, err := ws.Column(column.UUID)
	if err == nil {
		if !col.Permissions.Update {
			return fmt.Errorf("no permissions to update column %s", column.UUID)
		}
	}

	return db.DB.SaveColumn(workspace, column.Column)
}

func (a *Access) SaveColumns(workspace string, columns []*column.Column) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.UpdateColumns {
		return fmt.Errorf("no permissions to update column in workspace %s", workspace)
	}

	for _, c := range columns {
		col, err := ws.Column(c.UUID)
		if err == nil {
			if !col.Permissions.Update {
				return fmt.Errorf("no permissions to update column %s", c.UUID)
			}
		}
	}

	var c []*types.Column
	for _, col := range columns {
		c = append(c, col.Column)
	}

	return db.DB.SaveColumns(workspace, c)
}

func (a *Access) GetColumn(workspace, uuid string) (*column.Column, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllColumns {
		_, err := ws.Column(uuid) //when it is in the slice, the person has access to it
		if err != nil {
			return nil, err
		}
	}

	c, err := db.DB.GetColumn(workspace, uuid)
	if err != nil {
		return nil, err
	}
	return &column.Column{Column: c, Workspace: workspace}, nil

}

func (a *Access) DeleteColumnOnBoard(workspace, board, column string) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.RemoveColumnsOnBoard {
		return fmt.Errorf("no permissions to remove column in workspace %s", workspace)
	}

	brd, err := ws.Board(board)
	if err == nil {
		if !brd.Permissions.RemoveColumns {
			return fmt.Errorf("no permissions to remove column %s", column)
		}
	}

	col, err := ws.Column(column)
	if err == nil {
		if !col.Permissions.RemoveFromBoard {
			return fmt.Errorf("no permissions to remove column %s", column)
		}
	}

	return db.DB.DeleteColumnOnBoard(workspace, board, column)

}

func (a *Access) RenameColumn(workspace, column, name string) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.RenameColumns {
		return fmt.Errorf("no permissions to rename column in workspace %s", workspace)
	}

	col, err := ws.Column(column)
	if err == nil {
		if !col.Permissions.Rename {
			return fmt.Errorf("no permissions to rename column %s", column)
		}
	}

	return db.DB.RenameColumn(workspace, column, name)

}

func (a *Access) DeleteColumn(workspace, uuid string) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.DeleteColumns {
		return fmt.Errorf("no permissions to delete column in workspace %s", workspace)
	}

	col, err := ws.Column(uuid)
	if err == nil {
		if !col.Permissions.Delete {
			return fmt.Errorf("no permissions to delete column %s", uuid)
		}
	}

	return db.DB.DeleteColumn(workspace, uuid)

}

func (a *Access) MoveColumn(workspace, board, uuid string, toIndex int) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.MoveColumns {
		return fmt.Errorf("no permissions to move column in workspace %s", workspace)
	}

	brd, err := ws.Board(board)
	if err == nil {
		if !brd.Permissions.MoveColumns {
			return fmt.Errorf("no permissions to move column %s", uuid)
		}
	}

	col, err := ws.Column(uuid)
	if err == nil {
		if !col.Permissions.Move {
			return fmt.Errorf("no permissions to move column %s", uuid)
		}
	}

	return db.DB.MoveColumn(workspace, board, uuid, toIndex)

}

func (a *Access) NewColumn(workspace, board, name string) (*column.Column, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.Permissions.CreateColumns {
		return nil, fmt.Errorf("no permissions to create column in workspace %s", workspace)
	}

	brd, err := ws.Board(board)
	if err == nil {
		if !brd.Permissions.CreateColumns {
			return nil, fmt.Errorf("no permissions to create column %s", board)
		}
	}

	c, err := db.DB.NewColumn(workspace, board, name)
	if err != nil {
		return nil, err
	}
	return &column.Column{Column: c, Workspace: workspace}, nil

}

// Task functions
func (a *Access) SaveTask(workspace string, task *task.Task) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.UpdateTasks {
		return fmt.Errorf("no permissions to update task in workspace %s", workspace)
	}

	t, err := ws.Task(task.UUID)
	if err == nil {
		if !t.Permissions.Update {
			return fmt.Errorf("no permissions to update task %s", task.UUID)
		}
	}

	return db.DB.SaveTask(workspace, task.Task)

}

func (a *Access) SaveTasks(workspace string, tasks []*task.Task) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.UpdateTasks {
		return fmt.Errorf("no permissions to update task in workspace %s", workspace)
	}

	for _, t := range tasks {
		tsk, err := ws.Task(t.UUID)
		if err == nil {
			if !tsk.Permissions.Update {
				return fmt.Errorf("no permissions to update task %s", t.UUID)
			}
		}
	}

	var t []*types.Task
	for _, tsk := range tasks {
		t = append(t, tsk.Task)
	}

	return db.DB.SaveTasks(workspace, t)

}

func (a *Access) GetTask(workspace, uuid string) (*task.Task, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllTasks {
		_, err := ws.Task(uuid) //when it is in the slice, the person has access to it
		if err != nil {
			return nil, err
		}
	}

	t, err := db.DB.GetTask(workspace, uuid)
	if err != nil {
		return nil, err
	}
	return &task.Task{Task: t, Workspace: workspace}, nil

}

func (a *Access) DeleteTaskOnColumn(workspace, column, uuid string) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.RemoveTasksOnColumn {
		return fmt.Errorf("no permissions to remove task in workspace %s", workspace)
	}

	col, err := ws.Column(column)
	if err == nil {
		if !col.Permissions.RemoveTasks {
			return fmt.Errorf("no permissions to remove task %s", uuid)
		}
	}

	t, err := ws.Task(uuid)
	if err == nil {
		if !t.Permissions.RemoveFromColumn {
			return fmt.Errorf("no permissions to remove task %s", uuid)
		}
	}

	return db.DB.DeleteTaskOnColumn(workspace, column, uuid)

}

func (a *Access) DeleteTask(workspace, uuid string) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.DeleteTasks {
		return fmt.Errorf("no permissions to delete task in workspace %s", workspace)
	}

	t, err := ws.Task(uuid)
	if err == nil {
		if !t.Permissions.Delete {
			return fmt.Errorf("no permissions to delete task %s", uuid)
		}
	}

	return db.DB.DeleteTask(workspace, uuid)

}

func (a *Access) MoveTask(workspace, column, uuid, toColumn string, toIndex int) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.MoveTasks {
		return fmt.Errorf("no permissions to move task in workspace %s", workspace)
	}

	col, err := ws.Column(column)
	if err == nil {
		if !col.Permissions.MoveTasks {
			return fmt.Errorf("no permissions to move task %s", uuid)
		}
	}

	t, err := ws.Task(uuid)
	if err == nil {
		if !t.Permissions.Move {
			return fmt.Errorf("no permissions to move task %s", uuid)
		}
	}

	return db.DB.MoveTask(workspace, column, uuid, toColumn, toIndex)

}

func (a *Access) NewTask(workspace, column, name string) (*task.Task, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.Permissions.CreateTasks {
		return nil, fmt.Errorf("no permissions to create task in workspace %s", workspace)
	}

	col, err := ws.Column(column)
	if err == nil {
		if !col.Permissions.CreateTasks {
			return nil, fmt.Errorf("no permissions to create task %s", column)
		}
	}

	t, err := db.DB.NewTask(workspace, column, name)
	if err != nil {
		return nil, err
	}
	return &task.Task{Task: t, Workspace: workspace}, nil

}

func (a *Access) RenameTask(workspace, task, name string) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.RenameTasks {
		return fmt.Errorf("no permissions to rename task in workspace %s", workspace)
	}

	t, err := ws.Task(task)
	if err == nil {
		if !t.Permissions.Rename {
			return fmt.Errorf("no permissions to rename task %s", task)
		}
	}

	return db.DB.RenameTask(workspace, task, name)

}

func (a *Access) RemoveTagOnTask(workspace, task, uuid string) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.UpdateTagsOnTask {
		return fmt.Errorf("no permissions to remove tag on task in workspace %s", workspace)
	}

	tsk, err := ws.Task(task)
	if err == nil {
		if !tsk.Permissions.UpdateTags {
			return fmt.Errorf("no permissions to remove tag on task %s", task)
		}
	}

	t, err := ws.Tag(uuid)
	if err == nil {
		if !t.Permissions.UpdateOn {
			return fmt.Errorf("no permissions to remove tag on task %s", task)
		}
	}

	return db.DB.RemoveTagOnTask(workspace, task, uuid)
}

func (a *Access) SetTagsOnTask(workspace, task string, tags []string) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.UpdateTagsOnTask {
		return fmt.Errorf("no permissions to set tags on task in workspace %s", workspace)
	}

	tsk, err := ws.Task(task)
	if err == nil {
		if !tsk.Permissions.UpdateTags {
			return fmt.Errorf("no permissions to set tags on task %s", task)
		}
	}

	return db.DB.SetTagsOnTask(workspace, task, tags)
}

func (a *Access) SetTaskDescription(workspace, task, description string) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.UpdateTaskDescription {
		return fmt.Errorf("no permissions to set description on task in workspace %s", workspace)
	}

	tsk, err := ws.Task(task)
	if err == nil {
		if !tsk.Permissions.UpdateDescription {
			return fmt.Errorf("no permissions to set description on task %s", task)
		}
	}

	return db.DB.SetTaskDescription(workspace, task, description)
}

// Tag functions
func (a *Access) GetAllTags(workspace string) ([]*tag.Tag, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllTags {
		var tags []*tag.Tag
		for _, tsk := range ws.Tags {
			t, err := db.DB.GetTag(workspace, tsk.UUID)
			if err != nil {
				return nil, err
			}

			tags = append(tags, &tag.Tag{Tag: t, Workspace: workspace})
		}

		return tags, nil
	}

	t, err := db.DB.GetAllTags(workspace)
	if err != nil {
		return nil, err
	}

	var tags []*tag.Tag
	for _, tsk := range t {
		tags = append(tags, &tag.Tag{Tag: tsk, Workspace: workspace})
	}
	return tags, nil

}

func (a *Access) GetTag(workspace, uuid string) (*tag.Tag, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllTags {
		_, err := ws.Tag(uuid) //when it is in the slice, the person has access to it
		if err != nil {
			return nil, err
		}
	}

	t, err := db.DB.GetTag(workspace, uuid)
	if err != nil {
		return nil, err
	}
	return &tag.Tag{Tag: t, Workspace: workspace}, nil

}

func (a *Access) AddTagToTask(workspace, task, tag string) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.UpdateTagsOnTask {
		return fmt.Errorf("no permissions to add tag to task in workspace %s", workspace)
	}

	tsk, err := ws.Task(task)
	if err == nil {
		if !tsk.Permissions.UpdateTags {
			return fmt.Errorf("no permissions to add tag to task %s", task)
		}
	}

	t, err := ws.Tag(tag)
	if err == nil {
		if !t.Permissions.UpdateOn {
			return fmt.Errorf("no permissions to add tag to task %s", task)
		}
	}

	return db.DB.AddTagToTask(workspace, task, tag)

}

//Other functions

func (a *Access) GetStatus(workspace, uuid string) (*types.Status, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllStatuses {
		_, err := ws.Status(uuid) //when it is in the slice, the person has access to it
		if err != nil {
			return nil, err
		}
	}

	return db.DB.GetStatus(workspace, uuid)
}

func (a *Access) GetPriority(workspace, uuid string) (*types.Priority, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllPriorities {
		_, err := ws.Priority(uuid) //when it is in the slice, the person has access to it
		if err != nil {
			return nil, err
		}
	}

	return db.DB.GetPriority(workspace, uuid)

}

func (a *Access) GetChecklist(workspace, uuid string) (*types.Checklist, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllChecklists {
		_, err := ws.Checklist(uuid) //when it is in the slice, the person has access to it
		if err != nil {
			return nil, err
		}
	}

	return db.DB.GetChecklist(workspace, uuid)

}

func (a *Access) GetAttachment(workspace, uuid string) (*types.Attachment, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllAttachments {
		_, err := ws.Attachment(uuid) //when it is in the slice, the person has access to it
		if err != nil {
			return nil, err
		}
	}

	return db.DB.GetAttachment(workspace, uuid)
}

func (a *Access) GetDate(workspace, uuid string) (*types.Date, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllDates {
		_, err := ws.Date(uuid) //when it is in the slice, the person has access to it
		if err != nil {
			return nil, err
		}
	}

	return db.DB.GetDate(workspace, uuid)
}

func (a *Access) ListWorkspaces() ([]*workspace.List, error) {
	var list []*workspace.List

	for _, w := range a.Workspaces {
		var ws, err = db.DB.GetWorkspace(w.UUID) // we don't need to check for permissions, because we already have them => saves time
		if err != nil {
			return nil, err //TODO: this could be problematic, because when we haven't synced the database and so maybe not removed the workspace from the user but from the database
		}

		list = append(list, &workspace.List{
			UUID:        ws.UUID,
			Name:        ws.Name,
			Description: ws.Description,
			Cover:       ws.Cover,
			Archived:    ws.Archived,
			Color:       ws.Color,
		})
	}

	return list, nil
}

func (a *Access) WorkspaceInfo() ([]*workspace.Info, error) {
	var info []*workspace.Info

	for _, w := range a.Workspaces {
		var ws, err = db.DB.GetWorkspace(w.UUID) // we don't need to check for permissions, because we already have them => saves time
		if err != nil {
			return nil, err //TODO: this could be problematic, because when we haven't synced the database and so maybe not removed the workspace from the user but from the database
		}

		var boards []*types.NameList
		if w.AllBoards {
			var b, err = db.DB.GetAllBoardNames(w.UUID)
			if err != nil {
				return nil, err
			}
			boards = b
		} else {
			var brds []string
			for _, b := range w.Boards {
				brds = append(brds, b.UUID)
			}
			var b, err = db.DB.GetBoardNames(w.UUID, brds)
			if err != nil {
				return nil, err
			}
			boards = b

		}
		if err != nil {
			return nil, err
		}

		info = append(info, &workspace.Info{
			List: &workspace.List{
				UUID:        ws.UUID,
				Name:        ws.Name,
				Description: ws.Description,
				Cover:       ws.Cover,
				Archived:    ws.Archived,
				Color:       ws.Color,
			},
			Boards: boards,
		})
	}

	return info, nil

}

//GetUser(uuid string) (*types.Member, error) TODO

func (a *Access) workspace(uuid string) (*workspaceaccess.WorkspaceAccess, error) {
	for _, w := range a.Workspaces {
		if w.UUID == uuid {
			return &w, nil
		}
	}
	return nil, errors.New("workspace not found")
}

//func (a *Access) board(workspace, uuid string) (*BoardAccess, error) {
//	ws, err := a.workspace(workspace)
//	if err != nil {
//		return nil, err
//	}
//
//	return ws.board(uuid)
//}

func (a *Access) DeleteWorkspace(uuid string) error {
	ws, err := a.workspace(uuid)
	if err != nil {
		return err
	}

	if !ws.Permissions.DeleteWorkspace {
		return fmt.Errorf("no permissions to delete workspace %s", uuid)
	}

	return db.DB.DeleteWorkspace(uuid)
}

func (a *Access) GetElement(workspace string, uuid string) (*element.Element, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllElements {
		_, err := ws.Element(uuid) //when it is in the slice, the person has access to it
		if err != nil {
			return nil, err
		}
	}

	return db.DB.GetElement(workspace, uuid) //TODO
}

func (a *Access) GetField(workspace string, uuid string) (*field.Field, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllFields {
		_, err := ws.Field(uuid) //when it is in the slice, the person has access to it
		if err != nil {
			return nil, err
		}
	}

	//return db.DB.GetField(workspace, uuid) //TODO
	return nil, nil
}

func (a *Access) DeleteElement(workspace, uuid string) error {
	_, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	//TODO

	//if !ws.Permissions.DeleteElements {
	//	return fmt.Errorf("no permissions to delete element in workspace %s", workspace)
	//}
	//
	//elem, err := ws.Element(uuid)
	//if err == nil {
	//	if !elem.Permissions.Delete {
	//		return fmt.Errorf("no permissions to delete element %s", uuid)
	//	}
	//}
	//
	//return db.DB.DeleteElement(workspace, uuid)
	return nil
}
