package kanbandb

func GetBoard(boards []*types2.Board, uuid string) (*types2.Board, error) {
	for _, b := range boards {
		if b.UUID == uuid {
			return b, nil
		}
	}
	return nil, fmt.Errorf("board with uuid %s does not exist", uuid)
}

func GetBoards(boards []*types2.Board) []*types2.Board {
	return boards
}

func GetBoardNames(boards []*types2.Board) []*namelist.NameList {
	var names []*namelist.NameList
	for _, b := range boards {
		names = append(names, &namelist.NameList{
			Name: b.Name,
			UUID: b.UUID,
		})
	}
	return names
}

func GetColumn(columns []*types2.Column, uuid string) (*types2.Column, error) {
	for _, c := range columns {
		if c.UUID == uuid {
			return c, nil
		}
	}
	return nil, fmt.Errorf("column with uuid %s does not exist", uuid)
}

func GetTask(tasks []*types2.Task, uuid string) (*types2.Task, error) {
	for _, t := range tasks {
		if t.UUID == uuid {
			return t, nil
		}
	}
	return nil, fmt.Errorf("task with uuid %s does not exist", uuid)
}
