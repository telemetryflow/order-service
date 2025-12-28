// Package entity contains domain entities.
package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base contains common fields for all entities
type Base struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// NewBase creates a new Base with generated ID and timestamps
func NewBase() Base {
	now := time.Now()
	return Base{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// IsDeleted returns true if the entity has been soft-deleted
func (b *Base) IsDeleted() bool {
	return b.DeletedAt.Valid
}

// MarkUpdated updates the UpdatedAt timestamp
func (b *Base) MarkUpdated() {
	b.UpdatedAt = time.Now()
}

// MarkDeleted sets the DeletedAt timestamp for soft delete
func (b *Base) MarkDeleted() {
	now := time.Now()
	b.DeletedAt = gorm.DeletedAt{Time: now, Valid: true}
}

// Restore clears the DeletedAt timestamp
func (b *Base) Restore() {
	b.DeletedAt = gorm.DeletedAt{}
}

// BeforeCreate is a GORM hook that runs before creating a record
func (b *Base) BeforeCreate(_ *gorm.DB) error {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return nil
}
