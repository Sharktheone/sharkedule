package workspace

import "github.com/Sharktheone/sharkedule/kanban/column"

// Column functions
func (w *Workspace) SaveColumn(column *column.Column) error {

}

func (w *Workspace) SaveColumns(columns []*column.Column) error {

}

func (w *Workspace) GetColumn(uuid string) (*column.Column, error) {

}

func (w *Workspace) DeleteColumnOnBoard(board, column string) error {

}

func (w *Workspace) RenameColumn(column, name string) error {

}

func (w *Workspace) DeleteColumn(uuid string) error {

}

func (w *Workspace) MoveColumn(board, uuid string, toIndex int) error {

}

func (w *Workspace) NewColumn(board, name string) (*column.Column, error) {

}
