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
		api.GET("/task/:uuid", tasks.GetKanbanBoard)
	}

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}
