package routers

import (
	"github.com/gin-gonic/gin"
	"go-gorm/controllers/user"
)

func UserRouters(r *gin.Engine) {

	userRouter := r.Group("/user")
	//注册中间件
	//adminRouter.Use()
	{
		userRouter.GET("/index", user.UserController{}.Index)
		userRouter.GET("/add", user.UserController{}.Add)
		userRouter.GET("/edit", user.UserController{}.Edit)
	}

}
