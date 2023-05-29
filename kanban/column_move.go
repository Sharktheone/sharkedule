package kanban

func (c *Column) Move(toIndex int) error {
	board, err := c.GetParentBoard()
	if err != nil {
		return err
	}

	board.Columns = append(board.Columns[:c.Index], board.Columns[c.Index+1:]...)

	board.Columns = append(board.Columns[:toIndex], append([]*Column{c}, board.Columns[toIndex:]...)...)

	return board.Save()
}
