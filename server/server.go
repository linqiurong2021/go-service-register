package server

import (
	"fmt"

	"github.com/linqiurong2021/go-service-register/config"
	"github.com/linqiurong2021/go-service-register/model"
	"github.com/linqiurong2021/go-service-register/router"
	"github.com/linqiurong2021/go-service-register/service"

	"github.com/gin-gonic/gin"
)

// Server Server
type Server struct{}

// Start Start
func (s *Server) Start() {
	//
	enginer := gin.Default()
	// 注册路由
	router.RegisterRouter(enginer)
	//
	addr := fmt.Sprintf(":%d", config.Conf.Port)
	enginer.Run(addr)
}

// RegisterToGateway 注册到网关
func (s *Server) RegisterToGateway() {
	//
	var service = new(service.Service)
	//
	serviceData := &model.Service{
		Name:  "服务注册服务",
		Host:  config.Conf.Host,
		Port:  config.Conf.Port,
		URL:   "/service/register",
		Alive: true,
	}
	_, err := service.Create(serviceData)
	if err != nil {
		fmt.Println("Register Service Error: ", err.Error())
	}

	serviceData = &model.Service{
		Name:  "服务注册服务",
		Host:  config.Conf.Host,
		Port:  config.Conf.Port,
		URL:   "/service/unregister",
		Alive: true,
	}
	_, err = service.Create(serviceData)
	if err != nil {
		fmt.Println("Register Service Error: ", err.Error())
	}
}
