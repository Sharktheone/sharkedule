package kanbandb

import "github.com/Sharktheone/sharkedule/kanban/v2/board"

func DeleteBoard(boards []*board.Board, uuid string) {
	for index, b := range boards {
		if b.UUID == uuid {
			boards = append(boards[:index], boards[index+1:]...)
		}
	}
}
