package environment

import types2 "sharkedule/kanban/v2/types"

func GetBoardEnv(board *types2.Board) {

}

func GetColumnEnv(column *types2.Column) {

}

func GetTaskEnv(task *types2.Task) {

}

// THOUGHTS about this:

// Version 1
// board -> getEnv (all fields[columns, tags, members, etc.]) -> push to buffer
// -> getEnv for all columns in buffer -> push to buffer
// -> getEnv for all tasks in buffer -> push to buffer
// -> filter out duplicates -> update buffer

// Version 2
// board -> getEnv (all fields[columns, tags, members, etc.]) -> push uuids to buffer
// -> getEnv for all columns in buffer -> push non-duplicate uuids to buffer
// -> getEnv for all tasks in buffer -> push non-duplicate uuids to buffer
// -> get values for all uuids in buffer -> return env
