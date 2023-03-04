package service

import "IntegrationLab1/configs"

type Service struct {
	DocumentWriter
}

func NewService(cfg *configs.Config) *Service {
	return &Service{
		DocumentWriter: NewWriteDocService(cfg),
	}
}
