package service

import "github.com/linqiurong2021/go-service-register/model"

// Service 服务
type Service struct{}

var service *model.Service

func init() {
	service = new(model.Service)
}

// GetAllService 获取所有服务
func (s *Service) GetAllService() ([]*model.Service, error) {
	return service.GetAllService()
}

// Create 获取所有服务
func (s *Service) Create(inService *model.Service) (*model.Service, error) {

	return service.Create(inService)
}

// Update 获取所有服务
func (s *Service) Update(inService *model.Service) (*model.Service, error) {

	return service.Update(inService)
}

// Delete 删除服务
func (s *Service) Delete(ID uint) (bool, error) {

	return service.Delete(ID)
}

// UpdateAlive 获取所有服务
func (s *Service) UpdateAlive(ID string, alive bool) (bool, error) {
	return service.UpdateAlive(ID, alive)
}

// GetByID 获取所有服务
func (s *Service) GetByID(ID string) (*model.Service, error) {

	return service.GetByID(ID)
}

// GetServiceByHostAndPort 获取所有服务
func (s *Service) GetServiceByHostAndPort(host string, port uint) (*model.Service, error) {

	return service.GetServiceByHostAndPort(host, port)
}

// GetServiceByHostAndPortAndURL 获取所有服务
func (s *Service) GetServiceByHostAndPortAndURL(host string, port uint, url string) (*model.Service, error) {
	//
	return service.GetServiceByHostAndPortAndURL(host, port, url)
}

// GetByURL 获取所有服务
func (s *Service) GetByURL(URL string) (*model.Service, error) {

	return service.GetByURL(URL)
}
