package model

import "github.com/linqiurong2021/go-service-register/mysql"

// Create 创建
func (s *Service) Create(in *Service) (*Service, error) {
	if err := mysql.DB.Create(&in).Error; err != nil {
		return nil, err
	}
	return in, nil
}

// UpdateAlive 创建
func (s *Service) UpdateAlive(ID string, alive bool) (bool, error) {
	if err := mysql.DB.Where("ID = ?", ID).UpdateColumn("alive", alive).Error; err != nil {
		return false, err
	}
	return true, nil
}

// Update 创建
func (s *Service) Update(in *Service) (*Service, error) {
	if err := mysql.DB.Save(&in).Error; err != nil {
		return nil, err
	}
	return in, nil
}

// Delete 创建
func (s *Service) Delete(ID uint) (bool, error) {
	if err := mysql.DB.Where("ID = ?", ID).Delete(&Service{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

// GetByID 通过ID获取
func (s *Service) GetByID(ID string) (*Service, error) {
	var service = new(Service)
	if err := mysql.DB.Where("ID = ?", ID).Find(&service).Error; err != nil {
		return nil, err
	}
	return service, nil
}

// GetServiceByHostAndPort 通过Host与Port获取
func (s *Service) GetServiceByHostAndPort(host string, port uint) (*Service, error) {
	var service = new(Service)
	if err := mysql.DB.Where("host = ?", host).Where("port = ?", port).Find(&service).Error; err != nil {
		return nil, err
	}
	return service, nil
}

// GetServiceByHostAndPortAndURL 通过Host与Port与URL获取
func (s *Service) GetServiceByHostAndPortAndURL(host string, port uint, url string) (*Service, error) {
	var service = new(Service)
	if err := mysql.DB.Where("host = ?", host).Where("port = ?", port).Where("url = ?", url).Find(&service).Error; err != nil {
		return nil, err
	}
	return service, nil
}

// GetByURL 通过URL获取
func (s *Service) GetByURL(URL string) (*Service, error) {
	var service = new(Service)
	if err := mysql.DB.Where("url = ?", URL).Find(&service).Error; err != nil {
		return nil, err
	}
	return service, nil
}

// GetAllService 获取所有服务
func (s *Service) GetAllService() ([]*Service, error) {
	var services []*Service
	if err := mysql.DB.Find(&services).Error; err != nil {
		return nil, err
	}
	return services, nil
}
