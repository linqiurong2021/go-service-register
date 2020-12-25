package logic

import (
	"encoding/json"
	"fmt"

	"github.com/linqiurong2021/go-service-register/config"
	"github.com/linqiurong2021/go-service-register/etcd"
)

// AddConfItem AddConfItem
func AddConfItem(hasConfList []*etcd.EtcdProxyConfItem, newConf *etcd.EtcdProxyConfItem) (string, error) {
	// 把当前 service/register 添加到配置列表
	var newConfList = append(hasConfList[0:], newConf)
	// 获取当前服务器地址
	bytes, err := json.Marshal(newConfList)
	if err != nil {
		return "", err
	}
	value := string(bytes)
	return value, nil
}

// GetHasConfItemList 获取已有配置项
func GetHasConfItemList() (hasConfList []*etcd.EtcdProxyConfItem, err error) {

	// 获取已有的服务
	resp, err := etcd.Get(config.Conf.EtcdConfig.Key)
	if err != nil {
		fmt.Println("etcd set failure , err : ", err.Error())
		return nil, err
	}

	// 未获取到数据
	if resp.Kvs != nil {
		//
		err := json.Unmarshal(resp.Kvs[0].Value, &hasConfList)
		if err != nil {
			return nil, err
		}
	}
	return hasConfList, nil
}
