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

	api := r.Group("api")
	{
		task := api.Group("kanbanboard")
		{
			task.Get("list", tasks.ListKanbanBoards)
			task.Get("list/names", tasks.ListKanbanBoardNames)
			task.Put("new", tasks.CreateKanbanBoard)
			task.Get(":uuid", tasks.GetKanbanBoard)
		}
	}

	web.Serve(r)

	if err := r.Listen(":8080"); err != nil {
		panic(err)
	}

}
