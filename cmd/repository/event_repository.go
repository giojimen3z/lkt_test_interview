package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"lkt_test_interview/cmd/models"
)

type EventRepository interface {
	Create(ctx context.Context, e *models.Event) error
	List(ctx context.Context) ([]models.Event, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Event, error)
}

type eventRepository struct{ db *gorm.DB }

func NewEventRepository(db *gorm.DB) EventRepository { return &eventRepository{db: db} }

func (r *eventRepository) Create(ctx context.Context, e *models.Event) error {
	return r.db.WithContext(ctx).Create(e).Error
}

func (r *eventRepository) List(ctx context.Context) ([]models.Event, error) {
	var out []models.Event
	if err := r.db.WithContext(ctx).
		Order("start_time ASC").
		Find(&out).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (r *eventRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Event, error) {
	var ev models.Event
	if err := r.db.WithContext(ctx).
		First(&ev, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &ev, nil
}
