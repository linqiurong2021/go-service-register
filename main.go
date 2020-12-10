package main

import (
	"fmt"

	"github.com/linqiurong2021/go-gateway/config"
	"github.com/linqiurong2021/go-gateway/mysql"
	"github.com/linqiurong2021/go-service-register/model"
	"github.com/linqiurong2021/go-service-register/server"
	"github.com/linqiurong2021/go-service-register/validator"
)

func main() {
	//
	if err := config.Init("./config/config.ini"); err != nil {
		fmt.Printf("load config from file falure !, err:%v\n", err)
		return
	}

	if err := mysql.InitMySQL(config.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	// 模型绑定
	mysql.DB.AutoMigrate(&model.Service{})

	//
	validator := new(validator.Validator)

	// 开启校验转换
	if err := validator.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
	// 服务实例化
	var server = new(server.Server)
	// 注册服务
	server.RegisterToGateway()
	// 开启服务
	server.Start()
}
