package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/go-service-register/config"
	"github.com/linqiurong2021/go-service-register/controller"
	"github.com/linqiurong2021/go-service-register/logic"
	"github.com/linqiurong2021/go-service-register/model"
)

//
var register *controller.Register

// Router Router
type Router struct{}

// RegisterRouter 注册路由
func (r *Router) RegisterRouter(engine *gin.Engine) {
	service := engine.Group("/service")
	// 跨域问题
	service.POST("/register", register.Register)     // 新增
	service.POST("/unregister", register.UnRegister) // 删除
}

// RegisterToGateway 注册到网关
func (r *Router) RegisterToGateway() {
	//
	var logic = new(logic.Service)

	host := config.Conf.Host
	port := config.Conf.Port

	fmt.Printf("%s:%d", host, port)
	//

	serviceData := &model.Service{
		Name:  "服务注册服务",
		Host:  host,
		Port:  port,
		URL:   "/service/register",
		Alive: true,
	}
	// 有判断如果已存在则不再添加
	_, err := logic.CreateData(serviceData)
	if err != nil {
		fmt.Println("Register Service Error: ", err.Error())
	}
	serviceData.ID = serviceData.ID + 1
	serviceData.URL = "/service/unregister"
	_, err = logic.CreateData(serviceData)
	if err != nil {
		fmt.Println("Register Service Error: ", err.Error())
	}
}
