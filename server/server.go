package server

import (
	"fmt"

	"github.com/linqiurong2021/go-service-register/config"
	"github.com/linqiurong2021/go-service-register/router"

	"github.com/gin-gonic/gin"
)

// Server Server
type Server struct{}

// Start Start
func (s *Server) Start() {
	//
	enginer := gin.Default()
	// 注册路由
	router := new(router.Router)
	router.RegisterRouter(enginer)
	router.RegisterToGateway()
	//
	addr := fmt.Sprintf(":%d", config.Conf.Port)
	enginer.Run(addr)
}
