package router

import (
	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/go-service-register/controller"
)

//
var register *controller.Register

// RegisterRouter 注册路由
func RegisterRouter(r *gin.Engine) {
	service := r.Group("/service")
	// 跨域问题
	service.POST("/register", register.Register)     // 新增
	service.POST("/unregister", register.UnRegister) // 删除
}
