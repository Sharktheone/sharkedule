package router

import (
	"github.com/Sharktheone/sharkedule/api/kanbanboard"
	"github.com/Sharktheone/sharkedule/api/kanbanboard/column"
	"github.com/Sharktheone/sharkedule/api/kanbanboard/column/task"
	"github.com/Sharktheone/sharkedule/kanban/tag"
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

				board := boards.Group(":kanbanboard") // /api/kanban/board/:kanbanboard
				{

					board.Get("", kanbanboard.Get)             // GET /api/kanban/board/:kanbanboard
					board.Delete("delete", kanbanboard.Delete) // DELETE /api/kanban/board/:kanbanboard/delete

					columns := board.Group("column") // /api/kanban/board/:kanban/board/column
					{
						columns.Put("new", column.Create) // PUT /api/kanban/board/:kanbanboard/column/new

						col := columns.Group(":column") // /api/kanban/board/:kanbanboard/column/:column
						{
							col.Get("", column.Get)                    // GET /api/kanban/board/:kanbanboard/column/:column
							col.Delete("delete", column.DeleteOnBoard) // DELETE /api/kanban/board/:kanbanboard/column/:column/delete
							col.Patch("move", column.Move)             // PATCH /api/kanban/board/:kanbanboard/column/:column/move

							tsk := col.Group("task") // /api/kanban/board/:kanbanboard/column/:column/task
							{
								tsk.Put("new", task.Create) // PUT /api/kanban/board/:kanbanboard/column/:column/task/new

								t := tsk.Group(":task") // /api/kanban/board/:kanbanboard/column/:column/task/:task
								{
									t.Patch("move", task.Move)              // PATCH /api/kanban/board/:kanbanboard/column/:column/task/:task/move
									t.Get("", task.Get)                     // GET /api/kanban/board/:kanbanboard/column/:column/task/:task
									t.Delete("delete", task.DeleteOnColumn) // DELETE /api/kanban/board/:kanbanboard/column/:column/task/:task/delete
								}
							}
						}
					}
				}
			}
			columns := kanban.Group("column") // /api/kanban/column
			{
				col := columns.Group(":column") // /api/kanban/column/:column
				{
					col.Delete("delete", column.Delete) // DELETE /api/kanban/column/delete
					col.Patch("rename", column.Rename)  // PATCH /api/kanban/column/rename
				}
			}
			tasks := kanban.Group("task") // /api/kanban/task
			{
				t := tasks.Group(":task") // /api/kanban/task/:task
				{
					t.Delete("delete", task.Delete) // DELETE /api/kanban/task/:task/delete
					t.Patch("rename", task.Rename)  // PATCH /api/kanban/task/:task/rename
					t.Put("tag", task.AddTag)       // PUT /api/kanban/task/:task/tag
					t.Delete("tag", task.RemoveTag) // DELETE /api/kanban/task/:task/tag
					t.Patch("tags", task.SetTags)   // PATCH /api/kanban/task/:task/tag
				}
			}
			tags := kanban.Group("tag") // /api/kanban/tag
			{
				tags.Put("new", tag.NewTag)          // PUT /api/kanban/tag/new
				tags.Delete("delete", tag.DeleteTag) // DELETE /api/kanban/tag/delete
				tags.Patch("rename", tag.Rename)     // PATCH /api/kanban/tag/rename
				tags.Patch("update", tag.Update)     // PATCH /api/kanban/tag/update
				tags.Get("list", tag.GetTags)        // GET /api/kanban/tag/list

			}
		}
	}

	web.Serve(r)

	if err := r.Listen(":8080"); err != nil {
		panic(err)
	}

}
