package task

func (t *Task) Delete() error {
	column, err := t.GetParentColumn()
	if err != nil {
		return err
	}
	column.Tasks = append(column.Tasks[:t.Index], column.Tasks[t.Index+1:]...)

	return column.Save()
}