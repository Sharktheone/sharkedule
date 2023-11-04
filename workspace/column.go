package workspace

import "github.com/Sharktheone/sharkedule/kanban/column"

// Column functions
func (w *Workspace) SaveColumn(workspace string, column *column.Column) error {

}

func (w *Workspace) SaveColumns(workspace string, columns []*column.Column) error {

}

func (w *Workspace) GetColumn(workspace, uuid string) (*column.Column, error) {

}

func (w *Workspace) DeleteColumnOnBoard(workspace, board, column string) error {

}

func (w *Workspace) RenameColumn(workspace, column, name string) error {

}

func (w *Workspace) DeleteColumn(workspace, uuid string) error {

}

func (w *Workspace) MoveColumn(workspace, board, uuid string, toIndex int) error {

}

func (w *Workspace) NewColumn(workspace, board, name string) (*column.Column, error) {

}
