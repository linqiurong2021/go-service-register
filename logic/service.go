package logic

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/go-service-register/model"
	"github.com/linqiurong2021/go-service-register/service"
	"github.com/linqiurong2021/go-service-register/util"
	"gorm.io/gorm"
)

var serve *service.Service

// Service 服务
type Service struct {
}

// Delete 删除
type Delete struct {
	ID uint `json:"id" binding:"required"`
}

func init() {
	serve = new(service.Service)
}

// GetAllService 获取所有服务
func (s *Service) GetAllService() ([]*model.Service, error) {
	return serve.GetAllService()
}

// Create 获取所有服务
func (s *Service) Create(c *gin.Context) (*model.Service, error) {
	var inService *model.Service
	err := c.BindJSON(&inService)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ValidateFailure(err))
		return nil, nil
	}
	// 判断是否存在 如果已存在则直接返回
	service, err := serve.GetServiceByHostAndPortAndURL(inService.Host, inService.Port, inService.URL)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	// 如果已有端口号
	if service.ID > 0 {
		host := fmt.Sprintf("%s:%d:%s", inService.Host, inService.Port, inService.URL)
		return nil, errors.New(host + " has exists")
	}

	service, err = serve.Create(inService)
	if err != nil {
		return nil, err
	}
	c.JSON(http.StatusOK, util.Success("create success", ""))
	return service, nil
}

// Update 更新服务
func (s *Service) Update(inService *model.Service) (*model.Service, error) {

	return serve.Update(inService)
}

// UpdateAlive 更新健康状态
func (s *Service) UpdateAlive(ID string, alive bool) (bool, error) {
	return serve.UpdateAlive(ID, alive)
}

// Delete 删除服务
func (s *Service) Delete(c *gin.Context) (bool, error) {
	//
	var delete Delete
	err := c.BindJSON(&delete)
	if err != nil {
		return false, err
	}
	ok, err := serve.Delete(delete.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.BadRequest(err.Error(), ""))
	}
	c.JSON(http.StatusOK, util.Success("delete success", ""))
	return ok, nil
}

// GetByID 获取服务
func (s *Service) GetByID(ID string) (*model.Service, error) {

	return serve.GetByID(ID)
}

// GetByURL 获取服务
func (s *Service) GetByURL(URL string) (*model.Service, error) {

	return serve.GetByURL(URL)
}
