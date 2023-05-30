package task

import (
	"sharkedule/kanban/board"
	"sharkedule/kanban/board/column"
)

func (t *Task) GetParentBoard() (*board.Board, error) {
	return board.GetBoard(t.Board)
}

func (t *Task) GetParentColumn() (*column.Column, error) {
	board, err := t.GetParentBoard()
	if err != nil {
		return nil, err
	}
	return board.GetColumn(t.Column)
}
