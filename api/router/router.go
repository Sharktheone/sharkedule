package router

import (
	"github.com/Sharktheone/sharkedule/api/kanbanboard"
	"github.com/Sharktheone/sharkedule/api/kanbanboard/column"
	"github.com/Sharktheone/sharkedule/api/kanbanboard/column/task"
	"github.com/Sharktheone/sharkedule/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
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
			kboard.Get("list", kanbanboard.List)            // GET /api/kanbanboard/list
			kboard.Get("list/names", kanbanboard.ListNames) // GET /api/kanbanboard/list/names
			kboard.Put("new", kanbanboard.Create)           // PUT /api/kanbanboard/new

			singleBoard := kboard.Group(":kanbanboard") // /api/kanbanboard/:kanbanboard
			{

				singleBoard.Get("", kanbanboard.Get)             // GET /api/kanbanboard/:kanbanboard
				singleBoard.Delete("delete", kanbanboard.Delete) // DELETE /api/kanbanboard/:kanbanboard/delete

				col := singleBoard.Group("column") // /api/kanbanboard/:kanbanboard/column
				{
					col.Put("new", column.Create) // PUT /api/kanbanboard/:kanbanboard/column/new

					singleCol := col.Group(":column") // /api/kanbanboard/:kanbanboard/column/:column
					{
						singleCol.Get("", column.Get)             // GET /api/kanbanboard/:kanbanboard/column/:column
						singleCol.Delete("delete", column.Delete) // DELETE /api/kanbanboard/:kanbanboard/column/:column/delete
						singleCol.Patch("move", column.Move)      // PATCH /api/kanbanboard/:kanbanboard/column/:column/move

						tsk := singleCol.Group("task") // /api/kanbanboard/:kanbanboard/column/:column/task
						{
							tsk.Put("new", task.Create) // PUT /api/kanbanboard/:kanbanboard/column/:column/task/new

							singleTask := tsk.Group(":task") // /api/kanbanboard/:kanbanboard/column/:column/task/:task
							{
								singleTask.Patch("move", task.Move)      // PATCH /api/kanbanboard/:kanbanboard/column/:column/task/:task/move
								singleTask.Get("", task.Get)             // GET /api/kanbanboard/:kanbanboard/column/:column/task/:task
								singleTask.Delete("delete", task.Delete) // DELETE /api/kanbanboard/:kanbanboard/column/:column/task/:task/delete
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
