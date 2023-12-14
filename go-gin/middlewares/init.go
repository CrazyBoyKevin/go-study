package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleware(c *gin.Context) {
	//可以判断用户是否登录
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)
	c.Set("username", "Kevin")

	//在中间件中或handler中启动新的goroutine的时候，不能直接用c.context 需要copy使用副本
	//定义一个goroutine统计日志
	cCp := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done! in path" + cCp.Request.URL.Path)
	}()
}
