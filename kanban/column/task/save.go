package task

import "sharkedule/kanban/column"

func (t *Task) Save() error {
	col, err := t.GetParentColumn()
	if err != nil {
		return err
	}
	updateIndexes(col)
	col.Tasks[col.Index] = t

	return col.Save()
}

func updateIndexes(col *column.Column) {
	for i, t := range col.Tasks {
		t.Index = i
	}
}
