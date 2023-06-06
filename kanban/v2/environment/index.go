package environment

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/v2/task"
	types2 "github.com/Sharktheone/sharkedule/kanban/v2/types"
	"log"
)

type Environment struct {
	types2.Environment
	tagUUIDs        []*string                      `json:"-"`
	statusUUIDs     []*string                      `json:"-"`
	priorityUUIDs   []*string                      `json:"-"`
	columnUUIDs     []*string                      `json:"-"`
	taskUUIDs       []*string                      `json:"-"`
	memberUUIDs     []*string                      `json:"-"`
	checklistUUIDs  []*string                      `json:"-"`
	attachmentUUIDs []*string                      `json:"-"`
	dateUUIDs       []*string                      `json:"-"`
	actionUUIDs     []*string                      `json:"-"`
	BoardNames      map[string]string              `json:"board_names,omitempty"`     // boardUUID -> name
	ColumnNames     map[string]map[string]string   `json:"column_names,omitempty"`    // columnUUID -> boardUUID -> name
	DependentTasks  map[string]map[string][]string `json:"dependent_tasks,omitempty"` // taskUUID -> boardUUID -> columnUUID
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

func (e *Environment) GetIndexed() {
	for _, tag := range e.tagUUIDs {
		t, err := db.DBV2.GetTag(*tag) // TODO: add db function
		if err != nil {
			log.Printf("error getting tag: %v", err)
			continue
		}
		e.Tags = append(e.Tags, t)
	}

	for _, status := range e.statusUUIDs {
		s, err := db.DBV2.GetStatus(*status) // TODO: add db function
		if err != nil {
			log.Printf("error getting status: %v", err)
			continue
		}
		e.Status = append(e.Status, s)
	}

	for _, priority := range e.priorityUUIDs {
		p, err := db.DBV2.GetPriority(*priority) // TODO: add db function
		if err != nil {
			log.Printf("error getting priority: %v", err)
			continue
		}
		e.Priority = append(e.Priority, p)
	}
	for _, column := range e.columnUUIDs {
		c, err := db.DBV2.GetColumn(*column)
		if err != nil {
			log.Printf("error getting column: %v", err)
			continue
		}
		e.Columns = append(e.Columns, c)
	}

	for _, task := range e.taskUUIDs {
		t, err := db.DBV2.GetTask(*task)
		if err != nil {
			log.Printf("error getting task: %v", err)
			continue
		}
		e.Tasks = append(e.Tasks, t)

	}
	for _, member := range e.memberUUIDs {
		m, err := db.DBV2.GetMember(*member) // TODO: add db function
		if err != nil {
			log.Printf("error getting member: %v", err)
			continue
		}
		e.Members = append(e.Members, m)

	}
	for _, checklist := range e.checklistUUIDs {
		c, err := db.DBV2.GetChecklist(*checklist) // TODO: add db function
		if err != nil {
			log.Printf("error getting checklist: %v", err)
			continue
		}
		e.Checklists = append(e.Checklists, c)
	}

	for _, attachment := range e.attachmentUUIDs {
		a, err := db.DBV2.GetAttachment(*attachment) // TODO: add db function
		if err != nil {
			log.Printf("error getting attachment: %v", err)
			continue
		}
		e.Attachments = append(e.Attachments, a)

	}
	for _, date := range e.dateUUIDs {
		d, err := db.DBV2.GetDate(*date) // TODO: add db function
		if err != nil {
			log.Printf("error getting date: %v", err)
			continue
		}
		e.Dates = append(e.Dates, d)

	}
	//for _, action := range e.actionUUIDs { // TODO: Think this through, do we need actions in the Environment?
	//	a, err := db.DBV2.GetAction(*action) // TODO: add db function
	//	if err != nil {
	//		log.Printf("error getting action: %v", err)
	//		continue
	//	}
	//	e.Actions = append(e.Actions, a)
	//
	//}
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
	e.checklistUUIDs = AppendSliceIfMissing(e.checklistUUIDs, board.Checklists...)
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
		tsk, err := db.DBV2.GetTask(*t)
		if err != nil {
			log.Printf("error getting task: %v", err)
			continue
		}
		e.Tasks = append(e.Tasks, tsk)
	}

	for _, t := range e.Tasks {
		e.IndexTask(t)
	}
}

func (e *Environment) IndexTask(t *types2.Task) {
	e.tagUUIDs = AppendSliceIfMissing(e.tagUUIDs, t.Tags...)
	e.memberUUIDs = AppendSliceIfMissing(e.memberUUIDs, t.Members...)
	e.dateUUIDs = AppendSliceIfMissing(e.dateUUIDs, t.Dates...)
	e.attachmentUUIDs = AppendSliceIfMissing(e.attachmentUUIDs, t.Attachments...)
	e.checklistUUIDs = AppendSliceIfMissing(e.checklistUUIDs, t.CheckList...)

	for _, dep := range t.Dependencies {
		locations, err := task.GetLocations(dep)
		if err != nil {
			log.Printf("error getting locations: %v", err)
			continue
		}
		e.DependentTasks[dep] = locations
	}

	for _, dep := range t.Dependents {
		locations, err := task.GetLocations(dep)
		if err != nil {
			log.Printf("error getting locations: %v", err)
			continue
		}
		e.DependentTasks[dep] = locations
	}

	if t.Status != "" {
		e.statusUUIDs = AppendIfMissing(e.statusUUIDs, &t.Status)
	}
	if t.Priority != "" {
		e.priorityUUIDs = AppendIfMissing(e.priorityUUIDs, &t.Priority)
	}
	if t.DueDate != "" {
		e.dateUUIDs = AppendIfMissing(e.dateUUIDs, &t.DueDate)
	}
}

func AppendIfMissing(slice []*string, s *string) []*string {
	for _, ele := range slice {
		if *ele == *s {
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
