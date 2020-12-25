package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/go-service-register/config"
	"github.com/linqiurong2021/go-service-register/etcd"
	"github.com/linqiurong2021/go-service-register/logic"
)

// InitRouter 注册路由
func InitRouter(r *gin.Engine) {
	//
	service := r.Group("/service")
	service.POST("/register", func(c *gin.Context) {
		//
		var newConf *etcd.EtcdProxyConfItem
		err := c.BindJSON(&newConf)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}
		fmt.Printf("获取到的参数:%#v\n", newConf)
		// 已有的配置项
		hasConfList, err := logic.GetHasConfItemList()
		if err != nil {
			fmt.Printf("GetHasConfItemList err %s\n", err.Error())
			panic(err)
		}
		newConfStr, err := logic.AddConfItem(hasConfList, newConf)
		if err != nil {
			fmt.Printf("AddConfItem err %s\n", err.Error())
			panic(err)
		}
		//
		_, err = etcd.Set(config.Conf.EtcdConfig.Key, newConfStr)

		if err != nil {
			fmt.Println("etcd set failure , err : ", err.Error())
			return
		}
		c.JSON(http.StatusOK, "ok")
	})
}
