package main

import (
	"github.com/gin-gonic/gin"
	"go-gin/routers"
)

func main() {
	r := gin.Default()

	//配置全局中间件（可以多个）
	//r.Use()

	routers.ApiRoutersInit(r)
	routers.JsonRouters(r)
	routers.AdminRouters(r)

	r.Run(":8000")
}
