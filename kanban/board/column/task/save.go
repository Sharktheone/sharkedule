package task

import "sharkedule/kanban/board/column"

func (t *Task) Save() error {
	col, err := t.GetParentColumn()
	if err != nil {
		return err
	}
	updateTaskIndexes(col)
	col.Tasks[col.Index] = t

	return col.Save()
}

func updateTaskIndexes(col *column.Column) {
	for i, t := range col.Tasks {
		t.Index = i
	}
}
