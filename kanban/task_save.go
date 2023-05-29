package kanban

func (t *Task) Save() error {
	col, err := t.GetParentColumn()
	if err != nil {
		return err
	}
	updateTaskIndexes(col)
	col.Tasks[col.Index] = t

	return col.Save()
}

func updateTaskIndexes(col *Column) {
	for i, t := range col.Tasks {
		t.Index = i
	}
}
