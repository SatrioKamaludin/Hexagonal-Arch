package models

import (
	"time"

	"github.com/google/uuid"
)

type Profiling struct {
	ID        uuid.UUID `json:"id"`
	APICall   string    `json:"method"`
	Duration  int64     `json:"duration"`
	Timestamp time.Time `son:"timestamp"`
}
