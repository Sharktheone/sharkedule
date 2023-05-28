package task

import (
	"sharkedule/kanban"
	"sharkedule/kanban/column"
)

func (t *Task) GetParentBoard() (*kanban.Board, error) {
	return kanban.GetBoard(t.Board)
}

func (t *Task) GetParentColumn() (*column.Column, error) {
	board, err := t.GetParentBoard()
	if err != nil {
		return nil, err
	}
	return board.GetColumn(t.Column)
}
