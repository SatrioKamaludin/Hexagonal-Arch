package services

import (
	"CRUD-Go-Hexa-MongoDB/internal/domain/models"
	"CRUD-Go-Hexa-MongoDB/internal/ports"
	"time"
)

type ProfilingService struct {
	profilingRepo ports.IProfilingRepository
}

func NewProfilingService(profilingRepo ports.IProfilingRepository) *ProfilingService {
	return &ProfilingService{
		profilingRepo: profilingRepo,
	}
}

func (s *ProfilingService) Log(profiling models.Profiling) error {
	profiling.Timestamp = time.Now()
	return s.profilingRepo.Create(profiling)
}
