package etcd

import (
	"context"
	"fmt"
	"time"

	"github.com/linqiurong2021/go-service-register/config"
	"go.etcd.io/etcd/clientv3"
)

// EtcdProxyConfItem 配置项
type EtcdProxyConfItem struct {
	ID    string `binding:"required" json:"id"`
	Host  string `binding:"required" json:"host"`
	Port  uint   `binding:"required" json:"port"`
	URL   string `binding:"required" json:"url"`
	Name  string `binding:"required" json:"name"`
	Alive bool   `binding:"required" json:"alive"`
}

// EtcdClient EtcdClient
var EtcdClient *clientv3.Client

// Init Init
func Init(cfg *config.EtcdConfig) (client *clientv3.Client, err error) {
	//
	EtcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: cfg.DialTimeout * time.Second,
	})
	if err != nil {
		fmt.Printf("\n init etcd failure, err: %s\n", err.Error())
		return
	}
	return EtcdClient, err
}

// Get 设置
func Get(key string) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := EtcdClient.Get(ctx, key)
	cancel()
	if err != nil {
		// handle error!
	}
	return resp, err
}

// Set 设置
func Set(key string, value string) (*clientv3.PutResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := EtcdClient.Put(ctx, key, value)
	cancel()
	if err != nil {
		// handle error!
	}
	return resp, err
}
