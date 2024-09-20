package models

import "github.com/google/uuid"

type Product struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Stock int       `json:"stock"`
}
