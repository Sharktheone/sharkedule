package environment

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/task/locations"
	ktypes "github.com/Sharktheone/sharkedule/kanban/types"
	"github.com/Sharktheone/sharkedule/utils"
	"log"
)

type Environment struct {
	*ktypes.Environment
	boardUUIDs      []*string
	tagUUIDs        []*string
	statusUUIDs     []*string
	priorityUUIDs   []*string
	columnUUIDs     []*string
	taskUUIDs       []*string
	memberUUIDs     []*string
	checklistUUIDs  []*string
	attachmentUUIDs []*string
	dateUUIDs       []*string
	actionUUIDs     []*string
	workspace       string
}

func (e *Environment) Index() {
	if e.Boards != nil || e.boardUUIDs != nil {
		e.IndexBoards()
	}
	if e.Columns != nil || e.columnUUIDs != nil {
		e.IndexColumns()
	}
	if e.Tasks != nil || e.taskUUIDs != nil {
		e.IndexTasks()
	}

	e.GetIndexed()
}

func (e *Environment) GetIndexed() {
	//TODO: parallelize

	for _, tag := range e.tagUUIDs {
		t, err := db.DB.GetTag(e.workspace, *tag)
		if err != nil {
			log.Printf("error getting tag: %v", err)
			continue
		}
		e.Tags = append(e.Tags, t)
	}

	for _, col := range e.columnUUIDs {
		c, err := db.DB.GetColumn(e.workspace, *col)
		if err != nil {
			log.Printf("error getting column: %v", err)
			continue
		}
		e.Columns = append(e.Columns, c)
	}

	for _, t := range e.taskUUIDs {
		t, err := db.DB.GetTask(e.workspace, *t)
		if err != nil {
			log.Printf("error getting task: %v", err)
			continue
		}
		e.Tasks = append(e.Tasks, t)

	}
}

func (e *Environment) IndexBoards() {
	for _, b := range e.boardUUIDs {
		var found = false
		for _, bor := range e.Boards {
			if bor.UUID == *b {
				found = true
				break
			}
		}
		if found {
			break
		}
		board, err := db.DB.GetBoard(e.workspace, *b)
		if err != nil {
			log.Printf("error getting board: %v", err)
			continue
		}
		e.Boards = append(e.Boards, board)
	}

	for _, b := range e.Boards {
		e.IndexBoard(b)
	}
}

func (e *Environment) IndexBoard(b *ktypes.Board) {
	e.columnUUIDs = utils.AppendSliceIfMissing(e.columnUUIDs, b.Columns...)
	e.tagUUIDs = utils.AppendSliceIfMissing(e.tagUUIDs, b.Tags...)
	e.memberUUIDs = utils.AppendSliceIfMissing(e.memberUUIDs, b.Members...)
	e.dateUUIDs = utils.AppendSliceIfMissing(e.dateUUIDs, b.Dates...)
	e.attachmentUUIDs = utils.AppendSliceIfMissing(e.attachmentUUIDs, b.Attachments...)
	e.checklistUUIDs = utils.AppendSliceIfMissing(e.checklistUUIDs, b.Checklists...)
	e.actionUUIDs = utils.AppendSliceIfMissing(e.actionUUIDs, b.Actions...)

	if b.Status != "" {
		e.statusUUIDs = utils.AppendIfMissing(e.statusUUIDs, &b.Status)
	}
	if b.Priority != "" {
		e.priorityUUIDs = utils.AppendIfMissing(e.priorityUUIDs, &b.Priority)
	}
	if b.DueDate != "" {
		e.dateUUIDs = utils.AppendIfMissing(e.dateUUIDs, &b.DueDate)
	}
}

func (e *Environment) IndexColumns() {
	for _, c := range e.columnUUIDs {
		var found = false
		for _, col := range e.Columns {
			if col.UUID == *c {
				found = true
				break
			}
		}
		if found {
			break
		}
		col, err := db.DB.GetColumn(e.workspace, *c)
		if err != nil {
			log.Printf("error getting column: %v", err)
			continue
		}
		e.Columns = append(e.Columns, col)
	}

	for _, col := range e.Columns {
		e.IndexColumn(col)
	}
}

func (e *Environment) IndexColumn(column *ktypes.Column) {
	e.taskUUIDs = utils.AppendSliceIfMissing(e.taskUUIDs, column.Tasks...)
	e.tagUUIDs = utils.AppendSliceIfMissing(e.tagUUIDs, column.Tags...)

	for _, b := range column.Boards {
		bor, err := db.DB.GetBoard(e.workspace, b)
		if err != nil {
			log.Printf("error getting board: %v", err)
			continue
		}
		e.BoardNames[b] = bor.Name
	}
}

func (e *Environment) IndexTasks() {
	for _, t := range e.taskUUIDs {
		var found = false
		for _, tsk := range e.Tasks {
			if tsk.UUID == *t {
				found = true
				break
			}
		}
		if found {
			break
		}
		tsk, err := db.DB.GetTask(e.workspace, *t)
		if err != nil {
			log.Printf("error getting task: %v", err)
			continue
		}
		e.Tasks = append(e.Tasks, tsk)
	}

	for _, t := range e.Tasks {
		e.IndexTask(e.workspace, t)
	}
}

func (e *Environment) IndexTask(workspace string, t *ktypes.Task) {
	e.tagUUIDs = utils.AppendSliceIfMissing(e.tagUUIDs, t.Tags...)
	e.memberUUIDs = utils.AppendSliceIfMissing(e.memberUUIDs, t.Members...)
	e.dateUUIDs = utils.AppendSliceIfMissing(e.dateUUIDs, t.Dates...)
	e.attachmentUUIDs = utils.AppendSliceIfMissing(e.attachmentUUIDs, t.Attachments...)
	e.checklistUUIDs = utils.AppendSliceIfMissing(e.checklistUUIDs, t.Checklists...)
	e.workspace = workspace

	for _, dep := range t.Dependencies {
		loc, err := locations.GetLocations(e.workspace, dep)
		if err != nil {
			log.Printf("error getting locations: %v", err)
			continue
		}
		e.DependentTasks[dep] = loc
	}

	for _, dep := range t.Dependents {
		loc, err := locations.GetLocations(e.workspace, dep)
		if err != nil {
			log.Printf("error getting locations: %v", err)
			continue
		}
		e.DependentTasks[dep] = loc
	}

	if t.Status != "" {
		e.statusUUIDs = utils.AppendIfMissing(e.statusUUIDs, &t.Status)
	}
	if t.Priority != "" {
		e.priorityUUIDs = utils.AppendIfMissing(e.priorityUUIDs, &t.Priority)
	}
	if t.DueDate != "" {
		e.dateUUIDs = utils.AppendIfMissing(e.dateUUIDs, &t.DueDate)
	}
}
