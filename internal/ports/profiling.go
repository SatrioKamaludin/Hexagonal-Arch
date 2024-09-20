package ports

import (
	"CRUD-Go-Hexa-MongoDB/internal/domain/models"
)

type ProfilingRepository interface {
	Create(profiling models.Profiling) error
}
