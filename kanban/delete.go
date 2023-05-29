package kanban

import "sharkedule/database/db"

func (b *Board) Delete() error {
	boards, err := db.DB.GetBoards()
	if err != nil {
		return err
	}
	for index, board := range boards {
		if board.UUID == b.UUID {
			boards = append(boards[:index], boards[index+1:]...)
		}
	}
	return db.DB.SaveBoards(boards)
}
