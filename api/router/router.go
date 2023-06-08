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
		kanban := api.Group("kanban") // /api/kanban
		{
			boards := kanban.Group("board") // /api/kanban/board
			{
				boards.Get("list", kanbanboard.List)            // GET /api/kanban/board/list
				boards.Get("list/names", kanbanboard.ListNames) // GET /api/kanban/board/list/names
				boards.Put("new", kanbanboard.Create)           // PUT /api/kanban/board/new

				boards := boards.Group(":kanbanboard") // /api/kanban/board/:kanbanboard
				{

					boards.Get("", kanbanboard.Get)             // GET /api/kanban/board/:kanbanboard
					boards.Delete("delete", kanbanboard.Delete) // DELETE /api/kanban/board/:kanbanboard/delete

				}
			}
			columns := kanban.Group("column") // /api/kanban/column
			{
				columns.Put("new", column.Create) // PUT /api/kanban/column/new

				col := columns.Group(":column") // /api/kanban/column/:column
				{
					col.Get("", column.Get)             // GET /api/kanban/column/:column
					col.Delete("delete", column.Delete) // DELETE /api/kanban/column/:column/delete
					col.Patch("move", column.Move)      // PATCH /api/kanban/column/:column/move
				}
			}
			tsk := kanban.Group("task") // /api/kanban/task
			{
				tsk.Put("new", task.Create) // PUT /api/kanban/task/new

				t := tsk.Group(":task") // /api/kanban/task/:task
				{
					t.Patch("move", task.Move)      // PATCH /api/kanban/:task/move
					t.Get("", task.Get)             // GET /api/kanban/task/:task
					t.Delete("delete", task.Delete) // DELETE /api/kanban/task/:task/delete
				}
			}
		}

	}

	web.Serve(r)

	if err := r.Listen(":8080"); err != nil {
		panic(err)
	}

}
