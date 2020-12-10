package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/go-service-register/logic"
	"github.com/linqiurong2021/go-service-register/util"
)

// Register 注册服务
type Register struct{}

var register *logic.Service

func init() {
	register = new(logic.Service)
}

// Register 注册服务
func (reg *Register) Register(c *gin.Context) {
	//
	//
	_, err := register.Create(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.BadRequest(err.Error(), ""))
	}
}

// UnRegister 取消服务
func (reg *Register) UnRegister(c *gin.Context) {
	//
	_, err := register.Delete(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.BadRequest(err.Error(), ""))
	}
}
