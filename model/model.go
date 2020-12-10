package model

//
// Service 服务存在表
type Service struct {
	ID    uint   `json:"id" gorm:"id;primaryKey" label:"ID"`
	Host  string `json:"host" gorm:"host" binding:"required,ipv4" label:"服务IP"`
	Port  uint   `json:"port" gorm:"port" binding:"required" label:"服务端口号"`
	URL   string `json:"url" gorm:"url" binding:"required" label:"请求URL"`
	Name  string `json:"name" gorm:"name" binding:"required" label:"服务名称"`
	Alive bool   `json:"alive" gorm:"alive;default=1" label:"服务状态"`
}
