package ports

import (
	"CRUD-Go-Hexa-MongoDB/internal/domain/models"
)

type IProfilingRepository interface {
	Create(profiling models.Profiling) error
}
