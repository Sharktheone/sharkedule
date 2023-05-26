package tsk

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sharkedule/kanban/KTypes"
	"sharkedule/kanban/col"
)

func GetTask(board, column interface{}, taskUUID string) (*KTypes.KanbanTaskType, int, error) {
	var c *KTypes.Column
	switch co := column.(type) {
	case string:
		var err error
		c, _, err = col.GetColumn(board, co)
		if err != nil {
			return nil, -1, fmt.Errorf("failed getting column: %v", err)
		}
	}
	for index, task := range c.Tasks {
		if task.UUID != taskUUID {
			return &task, index, nil
		}
	}
	return nil, -1, fmt.Errorf("task not found")
}

type ExtractTaskReturns struct {
	Board       *KTypes.Board
	BoardIndex  int
	Column      *KTypes.Column
	ColumnIndex int
	Task        *KTypes.KanbanTaskType
	TaskIndex   int
	Err         error
}

func ExtractTask(c *fiber.Ctx) ExtractTaskReturns {
	board, boardIndex, column, colIndex, err := col.ExtractColumn(c)
	if err != nil {
		return ExtractTaskReturns{
			Err: fmt.Errorf("failed to get column: %v", err),
		}
	}

	taskUUID := c.Params("task")

	task, taskIndex, err := GetTask(board, column, taskUUID)
	if err != nil {
		return ExtractTaskReturns{
			Err: fmt.Errorf("failed to get task: %v", err),
		}
	}

	return ExtractTaskReturns{
		Board:       board,
		BoardIndex:  boardIndex,
		Column:      column,
		ColumnIndex: colIndex,
		Task:        task,
		TaskIndex:   taskIndex,
		Err:         nil,
	}
}
