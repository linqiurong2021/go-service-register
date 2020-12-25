package main

import (
	"fmt"

	"github.com/linqiurong2021/go-service-register/config"
	"github.com/linqiurong2021/go-service-register/etcd"
	"github.com/linqiurong2021/go-service-register/server"
)

func main() {
	//
	if err := config.Init("./config/config.ini"); err != nil {
		fmt.Printf("load config from file falure !, err:%v\n", err)
		return
	}
	// 初始化etcd 配置
	etcd.Init(config.Conf.EtcdConfig)
	// 服务实例化
	var server = new(server.Server)
	// 开启服务
	server.Start()
}
