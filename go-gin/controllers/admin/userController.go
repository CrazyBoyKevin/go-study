package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/models"
	"net/http"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	time := models.UnixToTime(1702297840)
	fmt.Printf(time)
	c.String(http.StatusOK, models.GetDate())
	//con.success(c)
}

func (con UserController) Add(c *gin.Context) {
	username, _ := c.Get("username")
	//类型断言 空接口类型-->string
	if v, err := username.(string); !err {
		fmt.Printf("err: Error")
	} else {
		fmt.Printf("username: %s", v)
	}
	c.String(http.StatusOK, "添加用户")
}

func (con UserController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "修改用户")
}
