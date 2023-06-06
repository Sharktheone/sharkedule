package environment

import (
	"github.com/Sharktheone/sharkedule/database/db"
	types2 "github.com/Sharktheone/sharkedule/kanban/v2/types"
	"log"
)

type Environment struct {
	types2.Environment
	tagUUIDs        []*string                    `json:"-"`
	statusUUIDs     []*string                    `json:"-"`
	priorityUUIDs   []*string                    `json:"-"`
	columnUUIDs     []*string                    `json:"-"`
	taskUUIDs       []*string                    `json:"-"`
	memberUUIDs     []*string                    `json:"-"`
	checklistUUIDs  []*string                    `json:"-"`
	attachmentUUIDs []*string                    `json:"-"`
	dateUUIDs       []*string                    `json:"-"`
	actionUUIDs     []*string                    `json:"-"`
	BoardNames      map[string]string            `json:"board_names,omitempty"`     // boardUUID -> name
	ColumnNames     map[string]map[string]string `json:"column_names,omitempty"`    // columnUUID -> boardUUID -> name
	DependentTasks  map[string]map[string]string `json:"dependent_tasks,omitempty"` // taskUUID -> boardUUID -> columnUUID
}

func (e *Environment) Index() {
	if e.Boards != nil {
		e.IndexBoards()
	} else if e.Columns != nil || e.columnUUIDs != nil {
		e.IndexColumns()
	} else if e.Tasks != nil || e.tagUUIDs != nil {
		e.IndexTasks()
	}

	// TODO: get tags, members, etc. from db
}

func (e *Environment) IndexBoards() {
	for _, board := range e.Boards {
		e.IndexBoard(board)
	}
}

func (e *Environment) IndexBoard(board *types2.Board) {
	e.columnUUIDs = AppendSliceIfMissing(e.columnUUIDs, board.Columns...)
	e.tagUUIDs = AppendSliceIfMissing(e.tagUUIDs, board.Tags...)
	e.memberUUIDs = AppendSliceIfMissing(e.memberUUIDs, board.Members...)
	e.dateUUIDs = AppendSliceIfMissing(e.dateUUIDs, board.Dates...)
	e.attachmentUUIDs = AppendSliceIfMissing(e.attachmentUUIDs, board.Attachments...)
	e.checklistUUIDs = AppendSliceIfMissing(e.checklistUUIDs, board.CheckList...)
	e.actionUUIDs = AppendSliceIfMissing(e.actionUUIDs, board.Actions...)

	if board.Status != "" {
		e.statusUUIDs = AppendIfMissing(e.statusUUIDs, &board.Status)
	}
	if board.Priority != "" {
		e.priorityUUIDs = AppendIfMissing(e.priorityUUIDs, &board.Priority)
	}
	if board.DueDate != "" {
		e.dateUUIDs = AppendIfMissing(e.dateUUIDs, &board.DueDate)
	}
	e.IndexColumns()
}

func (e *Environment) IndexColumns() {
	for _, c := range e.columnUUIDs {
		column, err := db.DBV2.GetColumn(*c)
		if err != nil {
			log.Printf("error getting column: %v", err)
			continue
		}
		e.Columns = append(e.Columns, column)
	}

	for _, column := range e.Columns {
		e.IndexColumn(column)
	}
}

func (e *Environment) IndexColumn(column *types2.Column) {
	e.taskUUIDs = AppendSliceIfMissing(e.taskUUIDs, column.Tasks...)
	e.tagUUIDs = AppendSliceIfMissing(e.tagUUIDs, column.Tags...)

	for _, b := range column.Boards {
		board, err := db.DB.GetBoard(b)
		if err != nil {
			log.Printf("error getting board: %v", err)
			continue
		}
		e.BoardNames[b] = board.Name
	}

	e.IndexTasks()
}

func (e *Environment) IndexTasks() {
	for _, t := range e.taskUUIDs {
		task, err := db.DBV2.GetTask(*t)
		if err != nil {
			log.Printf("error getting task: %v", err)
			continue
		}
		e.Tasks = append(e.Tasks, task)
	}

	for _, task := range e.Tasks {
		e.IndexTask(task)
	}
}

func (e *Environment) IndexTask(task *types2.Task) {
	e.tagUUIDs = AppendSliceIfMissing(e.tagUUIDs, task.Tags...)
	e.memberUUIDs = AppendSliceIfMissing(e.memberUUIDs, task.Members...)
	e.dateUUIDs = AppendSliceIfMissing(e.dateUUIDs, task.Dates...)
	e.attachmentUUIDs = AppendSliceIfMissing(e.attachmentUUIDs, task.Attachments...)
	e.checklistUUIDs = AppendSliceIfMissing(e.checklistUUIDs, task.CheckList...)

	//TODO: implement dependent tasks and dependencies

	if task.Status != "" {
		e.statusUUIDs = AppendIfMissing(e.statusUUIDs, &task.Status)
	}
	if task.Priority != "" {
		e.priorityUUIDs = AppendIfMissing(e.priorityUUIDs, &task.Priority)
	}
	if task.DueDate != "" {
		e.dateUUIDs = AppendIfMissing(e.dateUUIDs, &task.DueDate)
	}
}
