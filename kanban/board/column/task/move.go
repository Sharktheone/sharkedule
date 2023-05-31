package task

import (
	"sharkedule/database/db"
	"sharkedule/database/types"
)

func (t *Task) Move(toIndex int, toColumn string) error {
	col, err := t.GetParentColumn()
	if err != nil {
		return err
	}
	updateTaskIndexes(col)
	col.Tasks = append(col.Tasks[:t.Index], col.Tasks[t.Index+1:]...)
	if toColumn == col.UUID {
		task, err := t.Convert()
		if err != nil {
			return err
		}
		col.Tasks = append(col.Tasks[:toIndex], append([]*types.Task{task}, col.Tasks[toIndex:]...)...)
		return db.DB.SaveColumn(col.Board, col)
	}
	if err := db.DB.SaveColumn(col.Board, col); err != nil {
		return err
	}

	toCol, err := db.DB.GetColumn(col.Board, toColumn)
	if err != nil {
		return err
	}
	task, err := t.Convert()
	if err != nil {
		return err
	}
	updateTaskIndexes(toCol)
	toCol.Tasks = append(toCol.Tasks[:toIndex], append([]*types.Task{task}, toCol.Tasks[toIndex:]...)...)

	return db.DB.SaveColumn(toCol.Board, toCol)
}
