package services

import "github.com/railgorail/web-backend/internal/storage"

type Service struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) GetHelloMessage() string {
	return s.storage.GetMessage()
}
