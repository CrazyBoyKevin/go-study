package main

import (
	"github.com/gin-gonic/gin"
	"go-gorm/routers"
)

func main() {
	r := gin.Default()

	//配置全局中间件（可以多个）
	//r.Use()
	routers.UserRouters(r)

	r.Run(":8000")
}
