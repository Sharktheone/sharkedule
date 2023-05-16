package router

import (
	"github.com/gin-gonic/gin"
	"sharkedule/api/tasks"
	"sharkedule/web"
)

func Start() {
	r := gin.Default()

	r.Use(cors())

	web.Serve(r)

	api := r.Group("/api")
	{
		task := api.Group("/task")
		{
			task.GET("/:uuid", tasks.GetKanbanBoard)
			task.GET("/list", tasks.ListKanbanBoards)
			task.GET("/list/names", tasks.ListKanbanBoardNames)
			task.PUT("/", tasks.CreateKanbanBoard)
		}
	}

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}
