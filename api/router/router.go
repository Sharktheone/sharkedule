package router

import (
	"github.com/gin-gonic/gin"
	"sharkedule/web"
)

func Start() {
	r := gin.Default()

	web.Serve(r)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}
