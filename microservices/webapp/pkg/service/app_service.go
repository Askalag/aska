package service

import "github.com/askalag/aska/microservices/webapp/pkg"

type AppService struct {
	tcp pkg.ServicesTCP
}

func (app *AppService) StatusAll() map[string]interface{} {
	//TODO implement me
	panic("implement me")
}

func (app *AppService) Status() map[string]interface{} {
	return map[string]interface{}{
		"status": "ok",
	}
}

func NewAppService(tcp pkg.ServicesTCP) *AppService {
	return &AppService{
		tcp: tcp,
	}
}
