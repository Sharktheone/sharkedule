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

	api := r.Group("api")
	{
		kboard := api.Group("kanbanboard")
		{
			kboard.Get("list", kanbanboard.ListKanbanBoards)
			kboard.Get("list/names", kanbanboard.ListKanbanBoardNames)
			kboard.Put("new", kanbanboard.CreateKanbanBoard)

			singleBoard := kboard.Group(":kanbanboard")
			{

				singleBoard.Get("", kanbanboard.GetKanbanBoard)
				singleBoard.Delete("delete", kanbanboard.DeleteKanbanBoard)

				col := singleBoard.Group("column")
				{
					col.Put("new", column.CreateKanbanBoardColumn)

					singleCol := col.Group(":column")
					{
						singleCol.Get("", column.GetKanbanBoardColumn)
						singleCol.Delete("delete", column.DeleteKanbanBoardColumn)

						tsk := singleCol.Group("task")
						{
							tsk.Put("new", task.CreateKanbanBoardColumnTask)

							singleTask := tsk.Group(":task")
							{
								singleTask.Patch("move", task.MoveKanbanBoardColumnTask)
								singleTask.Get("", task.GetKanbanBoardColumnTask)
								singleTask.Delete("delete", task.DeleteKanbanBoardColumnTask)
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
