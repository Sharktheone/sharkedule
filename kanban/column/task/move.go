package task

func (t *Task) Move(toIndex int, toColumn string) error {
	col, err := t.GetParentColumn()
	if err != nil {
		return err
	}
	updateIndexes(col)
	col.Tasks = append(col.Tasks[:t.Index], col.Tasks[t.Index+1:]...)
	if toColumn == col.UUID {
		col.Tasks = append(col.Tasks[:toIndex], append([]*Task{t}, col.Tasks[toIndex:]...)...)
		return col.Save()
	}
	if err := col.Save(); err != nil {
		return err
	}

	board, err := col.GetParentBoard()
	if err != nil {
		return err
	}
	toCol, err := board.GetColumn(toColumn)
	if err != nil {
		return err
	}
	updateIndexes(toCol)
	toCol.Tasks = append(toCol.Tasks[:toIndex], append([]*Task{t}, toCol.Tasks[toIndex:]...)...)

	return toCol.Save()
}
