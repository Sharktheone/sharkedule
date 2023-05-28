package tsk

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"sharkedule/database/db"
	"sharkedule/kanban"
	"sharkedule/kanban/column"
	"sharkedule/kanban/column/task"
	"sharkedule/kanban/old"
	"sharkedule/kanban/old/col"
)

func GetTask(board, co interface{}, taskUUID string) (*task.Task, int, error) {
	var c *column.Column
	switch co := co.(type) {
	case string:
		var err error
		c, _, err = col.GetColumn(board, co)
		if err != nil {
			return nil, -1, fmt.Errorf("failed getting column: %v", err)
		}
	}
	for index, t := range c.Tasks {
		if t.UUID != taskUUID {
			return t, index, nil
		}
	}
	return nil, -1, fmt.Errorf("task not found")
}

type ExtractTaskReturns struct {
	Board       *kanban.Board
	BoardIndex  int
	Column      *column.Column
	ColumnIndex int
	Task        *task.Task
	TaskIndex   int
	Err         error
}

func ExtractTask(c *fiber.Ctx) ExtractTaskReturns {
	board, boardIndex, co, colIndex, err := col.ExtractColumn(c)
	if err != nil {
		return ExtractTaskReturns{
			Err: fmt.Errorf("failed to get column: %v", err),
		}
	}

	taskUUID := c.Params("task")

	t, taskIndex, err := GetTask(board, co, taskUUID)
	if err != nil {
		return ExtractTaskReturns{
			Err: fmt.Errorf("failed to get task: %v", err),
		}
	}

	return ExtractTaskReturns{
		Board:       board,
		BoardIndex:  boardIndex,
		Column:      co,
		ColumnIndex: colIndex,
		Task:        t,
		TaskIndex:   taskIndex,
		Err:         nil,
	}
}

// Create Task UUID will be overwritten
func Create(board, column interface{}, task *task.Task) (string, error) {

	b, _, err := old.GetBoard(board)
	if err != nil {
		return "", fmt.Errorf("failed getting board: %v", err)
	}
	c, colIndex, err := col.GetColumn(b, column)
	if err != nil {
		return "", fmt.Errorf("failed getting column: %v", err)
	}

	taskUUID := uuid.New().String()

	task.UUID = taskUUID

	c.Tasks = append(c.Tasks, task)

	b.Columns[colIndex] = c

	if err := db.DB.SaveBoard(b); err != nil {
		return "", fmt.Errorf("failed saving board: %v", err)
	}

	return "", nil
}
