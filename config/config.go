package config

import "gopkg.in/ini.v1"

// AppConfig App配置项
type AppConfig struct {
	Release        bool   `ini:"release"`
	Host           string `ini:"host"`
	Port           uint   `ini:"port"`
	*MySQLConfig   `ini:"mysql"`
	*RedisConfig   `ini:"redis"`
	*GatewayConfig `ini:"gateway"`
}

// MySQLConfig 数据库配置项
type MySQLConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	DB       string `ini:"db"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Charset  string `ini:"charset"`
}

// RedisConfig redis配置文件
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     uint   `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

// GatewayConfig 网关配置
type GatewayConfig struct {
	Host string `ini:"host"`
	Port uint   `ini:"prot"`
}

// Conf 配置
var Conf = new(AppConfig)

// Init 初始化
func Init(file string) error {
	return ini.MapTo(Conf, file)
}
