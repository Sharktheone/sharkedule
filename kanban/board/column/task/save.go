package task

import (
	"sharkedule/database/db"
	"sharkedule/database/types"
)

func (t *Task) Save() error {
	col, err := t.GetParentColumn()
	if err != nil {
		return err
	}
	updateTaskIndexes(col)
	task, err := t.Convert()
	if err != nil {
		return err
	}
	col.Tasks[col.Index] = task

	return db.DB.SaveColumn(col.Board, col)
}

func updateTaskIndexes(col *types.Column) {
	for i, t := range col.Tasks {
		t.Index = i
	}
}
