package router

import (
	"github.com/Sharktheone/sharkedule/api/board"
	"github.com/Sharktheone/sharkedule/api/column"
	"github.com/Sharktheone/sharkedule/api/task"
	"github.com/Sharktheone/sharkedule/kanban/tag"
	"github.com/Sharktheone/sharkedule/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

/// # Auth implementation options:
///
/// - handle it at the root level (r.Use(auth))
///		- handle each route in a switch case
///		- Pros/Cons:
///			+ clean at router level
///			- need for each route to be defined in two places
/// - handle it at the route level (r.Get("route", auth, handler))
///		- Pros/Cons:
/// 		- bloated handler definition
///			- harder to see auth when looking at the handler
///			+ you can see every route's auth in one place
/// 		+ handler is clean from any auth code
/// - handle it at the handler level (handler(c *fiber.Ctx, auth bool))
///		- Pros/Cons:
///			- handler has auth code
///			+ handler definition is clean
/// - pass auth "decision" to user struct
/// 	- __Handler calls user.<func>() => user checks if it is allowed and returns an error if not__
///		- Pros/Cons:
///			- handler has more code
///			+ handler definition is clean
///		- Handler calls user.<func>(workspace) => user checks if it is allowed and calls workspace.<func>() if it is
///			- Pros/Cons:
///				+ clean handler
///				- (user has more code)
/// - something like user.workspace(workspace).<func>()
///

func Start() {
	r := fiber.New()

	r.Use(cors.New())

	r.Use(logger.New())
	r.Use(recover.New())

	api := r.Group("api") // /api
	{
		workspace := api.Group(":workspace") // /api/:workspace
		{
			kanban := workspace.Group("kanban") // /api/:workspace/kanban
			{
				boards := kanban.Group("board") // /api/:workspace/kanban/board
				{
					boards.Get("", board.List)           // GET /api/:workspace/kanban/board/
					boards.Get("names", board.ListNames) // GET /api/:workspace/kanban/board//names
					boards.Put("new", board.Create)      // PUT /api/:workspace/kanban/board/new

					brd := boards.Group(":board") // /api/:workspace/kanban/board/:board
					{

						brd.Get("", board.Get)             // GET /api/:workspace/kanban/board/:board
						brd.Delete("delete", board.Delete) // DELETE /api/:workspace/kanban/board/:board/delete

						columns := brd.Group("column") // /api/:workspace/kanban/board/:kanban/board/column
						{
							columns.Put("new", column.Create) // PUT /api/:workspace/kanban/board/:board/column/new

							col := columns.Group(":column") // /api/:workspace/kanban/board/:board/column/:column
							{
								col.Get("", column.Get)                    // GET /api/:workspace/kanban/board/:board/column/:column
								col.Delete("delete", column.DeleteOnBoard) // DELETE /api/:workspace/kanban/board/:board/column/:column/delete
								col.Patch("move", column.Move)             // PATCH /api/:workspace/kanban/board/:board/column/:column/move

								tsk := col.Group("task") // /api/:workspace/kanban/board/:board/column/:column/task
								{
									tsk.Put("new", task.Create) // PUT /api/:workspace/kanban/board/:board/column/:column/task/new

									t := tsk.Group(":task") // /api/:workspace/kanban/board/:board/column/:column/task/:task
									{
										t.Patch("move", task.Move)              // PATCH /api/:workspace/kanban/board/:board/column/:column/task/:task/move
										t.Get("", task.Get)                     // GET /api/:workspace/kanban/board/:board/column/:column/task/:task
										t.Delete("delete", task.DeleteOnColumn) // DELETE /api/:workspace/kanban/board/:board/column/:column/task/:task/delete
									}
								}
							}
						}
					}
				}
				columns := kanban.Group("column") // /api/:workspace/kanban/column
				{
					col := columns.Group(":column") // /api/:workspace/kanban/column/:column
					{
						col.Delete("delete", column.Delete) // DELETE /api/:workspace/kanban/column/delete
						col.Patch("rename", column.Rename)  // PATCH /api/:workspace/kanban/column/rename
					}
				}
				tasks := kanban.Group("task") // /api/:workspace/kanban/task
				{
					t := tasks.Group(":task") // /api/:workspace/kanban/task/:task
					{
						t.Delete("delete", task.Delete)             // DELETE /api/:workspace/kanban/task/:task/delete
						t.Patch("rename", task.Rename)              // PATCH /api/:workspace/kanban/task/:task/rename
						t.Put("tag", task.AddTag)                   // PUT /api/:workspace/kanban/task/:task/tag
						t.Delete("tag", task.RemoveTag)             // DELETE /api/:workspace/kanban/task/:task/tag
						t.Patch("tags", task.SetTags)               // PATCH /api/:workspace/kanban/task/:task/tag
						t.Patch("description", task.SetDescription) // PATCH /api/:workspace/kanban/task/:task/description
					}
				}
				tags := kanban.Group("tag") // /api/:workspace/kanban/tag
				{
					tags.Put("new", tag.NewTag)          // PUT /api/:workspace/kanban/tag/new
					tags.Delete("delete", tag.DeleteTag) // DELETE /api/:workspace/kanban/tag/delete
					tags.Patch("rename", tag.Rename)     // PATCH /api/:workspace/kanban/tag/rename
					tags.Patch("update", tag.Update)     // PATCH /api/:workspace/kanban/tag/update
					tags.Get("list", tag.GetTags)        // GET /api/:workspace/kanban/tag/list

				}
			}
		}
	}

	web.Serve(r)

	if err := r.Listen(":8080"); err != nil {
		panic(err)
	}

}
