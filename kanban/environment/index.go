package environment

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/task/locations"
	"github.com/Sharktheone/sharkedule/kanban/types"
	"log"
)

type Environment struct {
	*types.Environment
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
}

func (e *Environment) Index() {
	if e.Boards != nil || e.boardUUIDs != nil {
		e.IndexBoards()
	} else if e.Columns != nil || e.columnUUIDs != nil {
		e.IndexColumns()
	} else if e.Tasks != nil || e.tagUUIDs != nil {
		e.IndexTasks()
	}

	e.GetIndexed()
}

func (e *Environment) GetIndexed() {
	for _, tag := range e.tagUUIDs {
		t, err := db.DB.GetTag(*tag)
		if err != nil {
			log.Printf("error getting tag: %v", err)
			continue
		}
		e.Tags = append(e.Tags, t)
	}

	for _, status := range e.statusUUIDs {
		s, err := db.DB.GetStatus(*status)
		if err != nil {
			log.Printf("error getting status: %v", err)
			continue
		}
		e.Status = append(e.Status, s)
	}

	for _, priority := range e.priorityUUIDs {
		p, err := db.DB.GetPriority(*priority)
		if err != nil {
			log.Printf("error getting priority: %v", err)
			continue
		}
		e.Priority = append(e.Priority, p)
	}
	for _, col := range e.columnUUIDs {
		c, err := db.DB.GetColumn(*col)
		if err != nil {
			log.Printf("error getting column: %v", err)
			continue
		}
		e.Columns = append(e.Columns, c)
	}

	for _, t := range e.taskUUIDs {
		t, err := db.DB.GetTask(*t)
		if err != nil {
			log.Printf("error getting task: %v", err)
			continue
		}
		e.Tasks = append(e.Tasks, t)

	}
	for _, member := range e.memberUUIDs {
		m, err := db.DB.GetMember(*member)
		if err != nil {
			log.Printf("error getting member: %v", err)
			continue
		}
		e.Members = append(e.Members, m)

	}
	for _, checklist := range e.checklistUUIDs {
		c, err := db.DB.GetChecklist(*checklist)
		if err != nil {
			log.Printf("error getting checklist: %v", err)
			continue
		}
		e.Checklists = append(e.Checklists, c)
	}

	for _, attachment := range e.attachmentUUIDs {
		a, err := db.DB.GetAttachment(*attachment)
		if err != nil {
			log.Printf("error getting attachment: %v", err)
			continue
		}
		e.Attachments = append(e.Attachments, a)

	}
	for _, date := range e.dateUUIDs {
		d, err := db.DB.GetDate(*date)
		if err != nil {
			log.Printf("error getting date: %v", err)
			continue
		}
		e.Dates = append(e.Dates, d)

	}
	//for _, action := range e.actionUUIDs { // TODO: Think this through, do we need actions in the Environment? - YES
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
	for _, b := range e.boardUUIDs {
		board, err := db.DB.GetBoard(*b)
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

func (e *Environment) IndexBoard(b *types.Board) {
	e.columnUUIDs = AppendSliceIfMissing(e.columnUUIDs, b.Columns...)
	e.tagUUIDs = AppendSliceIfMissing(e.tagUUIDs, b.Tags...)
	e.memberUUIDs = AppendSliceIfMissing(e.memberUUIDs, b.Members...)
	e.dateUUIDs = AppendSliceIfMissing(e.dateUUIDs, b.Dates...)
	e.attachmentUUIDs = AppendSliceIfMissing(e.attachmentUUIDs, b.Attachments...)
	e.checklistUUIDs = AppendSliceIfMissing(e.checklistUUIDs, b.Checklists...)
	e.actionUUIDs = AppendSliceIfMissing(e.actionUUIDs, b.Actions...)

	if b.Status != "" {
		e.statusUUIDs = AppendIfMissing(e.statusUUIDs, &b.Status)
	}
	if b.Priority != "" {
		e.priorityUUIDs = AppendIfMissing(e.priorityUUIDs, &b.Priority)
	}
	if b.DueDate != "" {
		e.dateUUIDs = AppendIfMissing(e.dateUUIDs, &b.DueDate)
	}
	e.IndexColumns()
}

func (e *Environment) IndexColumns() {
	for _, c := range e.columnUUIDs {
		col, err := db.DB.GetColumn(*c)
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

func (e *Environment) IndexColumn(column *types.Column) {
	e.taskUUIDs = AppendSliceIfMissing(e.taskUUIDs, column.Tasks...)
	e.tagUUIDs = AppendSliceIfMissing(e.tagUUIDs, column.Tags...)

	for _, b := range column.Boards {
		bor, err := db.DB.GetBoard(b)
		if err != nil {
			log.Printf("error getting board: %v", err)
			continue
		}
		e.BoardNames[b] = bor.Name
	}

	e.IndexTasks()
}

func (e *Environment) IndexTasks() {
	for _, t := range e.taskUUIDs {
		tsk, err := db.DB.GetTask(*t)
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

func (e *Environment) IndexTask(t *types.Task) {
	e.tagUUIDs = AppendSliceIfMissing(e.tagUUIDs, t.Tags...)
	e.memberUUIDs = AppendSliceIfMissing(e.memberUUIDs, t.Members...)
	e.dateUUIDs = AppendSliceIfMissing(e.dateUUIDs, t.Dates...)
	e.attachmentUUIDs = AppendSliceIfMissing(e.attachmentUUIDs, t.Attachments...)
	e.checklistUUIDs = AppendSliceIfMissing(e.checklistUUIDs, t.CheckList...)

	for _, dep := range t.Dependencies {
		loc, err := locations.GetLocations(dep)
		if err != nil {
			log.Printf("error getting locations: %v", err)
			continue
		}
		e.DependentTasks[dep] = loc
	}

	for _, dep := range t.Dependents {
		loc, err := locations.GetLocations(dep)
		if err != nil {
			log.Printf("error getting locations: %v", err)
			continue
		}
		e.DependentTasks[dep] = loc
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

//func AppendMultipleIfMissing(slice []*string, s []*string) []*string {
//	for _, ele := range s {
//		slice = AppendIfMissing(slice, ele)
//	}
//	return slice
//}

func AppendSliceIfMissing(slice []*string, s ...string) []*string {
	for _, ele := range s {
		slice = AppendIfMissing(slice, &ele)
	}
	return slice
}
