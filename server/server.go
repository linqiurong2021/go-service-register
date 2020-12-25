package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/go-service-register/config"
	"github.com/linqiurong2021/go-service-register/etcd"
	"github.com/linqiurong2021/go-service-register/logic"
	"github.com/linqiurong2021/go-service-register/router"
)

// Server Server
type Server struct{}

// Start Start
func (s *Server) Start() {
	r := gin.Default()
	router.InitRouter(r)
	//
	s.RegisterToServer()
	port := fmt.Sprintf(":%d", config.Conf.Port)
	fmt.Printf("server start at port: %s\n", port)
	r.Run(port)
}

// RegisterToServer 注册服务
func (s *Server) RegisterToServer() {

	// 初始化
	etcd.Init(config.Conf.EtcdConfig)

	var localServiceConf = &etcd.EtcdProxyConfItem{
		Host:  config.Conf.Host,
		Port:  config.Conf.Port,
		Name:  "服务注册",
		URL:   "/service/register",
		ID:    "/service/register",
		Alive: true,
	}
	hasConfList, err := logic.GetHasConfItemList()
	if err != nil {
		fmt.Printf("GetHasConfItemList err %s\n", err.Error())
		panic(err)
	}
	newConf, err := logic.AddConfItem(hasConfList, localServiceConf)
	if err != nil {
		fmt.Printf("AddConfItem err %s\n", err.Error())
		panic(err)
	}
	//
	_, err = etcd.Set(config.Conf.EtcdConfig.Key, newConf)
	if err != nil {
		fmt.Println("etcd set failure , err : ", err.Error())
		return
	}
}
