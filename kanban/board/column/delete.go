package column

func (c *Column) Delete() error {
	board, err := c.GetParentBoard()
	if err != nil {
		return err
	}

	board.Columns = append(board.Columns[:c.Index], board.Columns[c.Index+1:]...)

	return board.Save()
}
