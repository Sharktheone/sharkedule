package environment

import types2 "github.com/Sharktheone/sharkedule/kanban/v2/types"

func GetBoardEnv(board *types2.Board) {
	env := Environment{
		Environment: types2.Environment{
			Boards: []*types2.Board{board},
		},
	}
	env.Index()

}

func GetColumnEnv(column *types2.Column) {
	env := Environment{
		Environment: types2.Environment{
			Columns: []*types2.Column{column},
		},
	}
	env.Index()

}

func GetTaskEnv(task *types2.Task) {
	env := Environment{
		Environment: types2.Environment{
			Tasks: []*types2.Task{task},
		},
	}
	env.Index()
}

func AppendIfMissing(slice []*string, s *string) []*string {
	for _, ele := range slice {
		if ele == s {
			return slice
		}
	}
	return append(slice, s)
}

func AppendMultipleIfMissing(slice []*string, s []*string) []*string {
	for _, ele := range s {
		slice = AppendIfMissing(slice, ele)
	}
	return slice
}

func AppendSliceIfMissing(slice []*string, s ...string) []*string {
	for _, ele := range s {
		slice = AppendIfMissing(slice, &ele)
	}
	return slice
}

// THOUGHTS about this:

// Version 1 xxxxxx
// board -> getEnv (all fields[columns, tags, members, etc.]) -> push to buffer
// -> getEnv for all columns in buffer -> push to buffer
// -> getEnv for all tasks in buffer -> push to buffer
// -> filter out duplicates -> update buffer

// Version 2 <<<<<
// board -> getEnv (all fields[columns, tags, members, etc.]) -> push uuids to buffer
// -> getEnv for all columns in buffer -> push non-duplicate uuids to buffer
// -> getEnv for all tasks in buffer -> push non-duplicate uuids to buffer
// -> get values for all uuids in buffer -> return env
