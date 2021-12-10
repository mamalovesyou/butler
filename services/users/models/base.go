package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID      `json=id`
	CreatedAt time.Time      `json=createdAt`
	UpdatedAt time.Time      `json=updatedAt`
	DeletedAt gorm.DeletedAt `json=deletedAt`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return nil
}
