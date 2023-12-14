package routers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string `json:"title" form:"title"`
	Desc    string `json:"desc" form:"desc"`
	Content string `json:"content" form:"content"`
}

func JsonRouters(r *gin.Engine) {
	//路由分组
	jsonRouters := r.Group("/json")
	{
		//获取post过来的xml
		r.POST("/getRaw", func(c *gin.Context) {
			data, err := c.GetRawData()
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			} else {
				a := &Article{}
				if err := json.Unmarshal(data, &a); err == nil {
					c.JSON(http.StatusOK, a)
				} else {
					c.JSON(http.StatusBadRequest, err)
				}
			}
		})

		jsonRouters.GET("/", func(c *gin.Context) {
			c.JSON(200, map[string]interface{}{
				"success": true,
				"msg":     "hello Gin",
			})
		})

		jsonRouters.GET("/1", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"success": true,
				"msg":     "hello Gin H",
			})
		})

		jsonRouters.GET("/2", func(c *gin.Context) {
			a := Article{
				Title:   "title",
				Desc:    "desc",
				Content: "content",
			}
			c.JSON(200, a)
		})

		//响应JSONP请求 callback=xxx
		//xxx({"title":"title-jsonp","desc":"desc","content":"content"});
		jsonRouters.GET("/p", func(c *gin.Context) {
			a := Article{
				Title:   "title-jsonp",
				Desc:    "desc",
				Content: "content",
			}
			c.JSONP(200, a)
		})
	}

}
