package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gorm/models"
	"net/http"
)

type UserController struct {
}

// Index 查询
func (con UserController) Index(c *gin.Context) {
	users := []models.User{}
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"msg":    "success",
		"result": users,
	})
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

func (con UserController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "删除用户")
}
