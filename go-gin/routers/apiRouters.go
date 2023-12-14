package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiRoutersInit(r *gin.Engine) {

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "值：%v", "你好Gin")
	})

	r.GET("/news", func(c *gin.Context) {
		c.String(200, "你好, 我是新闻页面")
	})

	r.GET("/test", func(c *gin.Context) {
		c.String(200, "你好, 我是来测试的")
	})

	r.GET("/xml", func(c *gin.Context) {
		a := Article{
			Title:   "title",
			Desc:    "desc",
			Content: "content",
		}
		c.XML(200, a)
	})

	r.GET("/querytest", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		page := c.DefaultQuery("page", "1")
		fmt.Println("username: " + username + ", password: " + password + ", page: " + page)
	})

	//form获取数据
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		fmt.Printf("username: %s, password: %s\n", username, password)
	})

	//获取数据绑定到结构体
	r.POST("/article2bind", func(c *gin.Context) {
		a := &Article{}
		if err := c.ShouldBind(&a); err == nil {
			c.JSON(http.StatusOK, a)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
		fmt.Printf("%#v", a)
	})

	//路由传值
	r.GET("/age/:age", func(c *gin.Context) {
		age := c.Param("age")
		fmt.Printf("age: %s", age)
		c.JSON(http.StatusOK, "OK")
	})

}
