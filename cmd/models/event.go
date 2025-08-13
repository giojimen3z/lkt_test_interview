package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Title       string    `json:"title" gorm:"size:100;not null;index"`
	Description *string   `json:"description,omitempty"`
	StartTime   time.Time `json:"start_time" gorm:"not null;index"`
	EndTime     time.Time `json:"end_time" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null;autoCreateTime"`
}
