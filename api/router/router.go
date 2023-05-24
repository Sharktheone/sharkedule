package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"sharkedule/api/kanbanboard"
	"sharkedule/api/kanbanboard/column"
	"sharkedule/api/kanbanboard/column/task"
	"sharkedule/web"
)

func Start() {
	r := fiber.New()

	r.Use(cors.New())

	r.Use(logger.New())
	r.Use(recover.New())

	api := r.Group("api") // /api
	{
		kboard := api.Group("kanbanboard") // /api/kanbanboard
		{
			kboard.Get("list", kanbanboard.ListKanbanBoards)           // GET /api/kanbanboard/list
			kboard.Get("list/names", kanbanboard.ListKanbanBoardNames) // GET /api/kanbanboard/list/names
			kboard.Put("new", kanbanboard.CreateKanbanBoard)           // PUT /api/kanbanboard/new

			singleBoard := kboard.Group(":kanbanboard") // /api/kanbanboard/:kanbanboard
			{

				singleBoard.Get("", kanbanboard.GetKanbanBoard)             // GET /api/kanbanboard/:kanbanboard
				singleBoard.Delete("delete", kanbanboard.DeleteKanbanBoard) // DELETE /api/kanbanboard/:kanbanboard/delete

				col := singleBoard.Group("column") // /api/kanbanboard/:kanbanboard/column
				{
					col.Put("new", column.CreateKanbanBoardColumn) // PUT /api/kanbanboard/:kanbanboard/column/new

					singleCol := col.Group(":column") // /api/kanbanboard/:kanbanboard/column/:column
					{
						singleCol.Get("", column.GetKanbanBoardColumn)             // GET /api/kanbanboard/:kanbanboard/column/:column
						singleCol.Delete("delete", column.DeleteKanbanBoardColumn) // DELETE /api/kanbanboard/:kanbanboard/column/:column/delete
						singleCol.Patch("move", column.MoveKanbanBoardColumn)      // PATCH /api/kanbanboard/:kanbanboard/column/:column/move

						tsk := singleCol.Group("task") // /api/kanbanboard/:kanbanboard/column/:column/task
						{
							tsk.Put("new", task.CreateKanbanBoardColumnTask) // PUT /api/kanbanboard/:kanbanboard/column/:column/task/new

							singleTask := tsk.Group(":task") // /api/kanbanboard/:kanbanboard/column/:column/task/:task
							{
								singleTask.Patch("move", task.MoveKanbanBoardColumnTask)      // PATCH /api/kanbanboard/:kanbanboard/column/:column/task/:task/move
								singleTask.Get("", task.GetKanbanBoardColumnTask)             // GET /api/kanbanboard/:kanbanboard/column/:column/task/:task
								singleTask.Delete("delete", task.DeleteKanbanBoardColumnTask) // DELETE /api/kanbanboard/:kanbanboard/column/:column/task/:task/delete
							}
						}
					}
				}
			}
		}
	}

	web.Serve(r)

	if err := r.Listen(":8080"); err != nil {
		panic(err)
	}

}
