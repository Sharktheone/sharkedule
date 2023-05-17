package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"sharkedule/api/tasks"
	"sharkedule/web"
)

func Start() {
	//r := gin.Default()
	r := fiber.New()

	r.Use(cors.New())

	web.Serve(r)

	api := r.Group("/api")
	{
		task := api.Group("/task")
		{
			task.Get("/:uuid", tasks.GetKanbanBoard)
			task.Get("/list", tasks.ListKanbanBoards)
			task.Get("/list/names", tasks.ListKanbanBoardNames)
			task.Get("/new", tasks.CreateKanbanBoard)
		}
	}

	if err := r.Listen(":8080"); err != nil {
		panic(err)
	}

}
