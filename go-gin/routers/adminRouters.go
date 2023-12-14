package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin/controllers/admin"
	"go-gin/middlewares"
)

func AdminRouters(r *gin.Engine) {

	adminRouter := r.Group("/admin", middlewares.InitMiddleware)
	//注册中间件
	//adminRouter.Use()
	{
		adminRouter.GET("/user", admin.UserController{}.Index)
		adminRouter.GET("/user/add", admin.UserController{}.Add)
		adminRouter.GET("/user/edit", admin.UserController{}.Edit)
	}

}
