package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"lkt_test_interview/cmd/models"
	"lkt_test_interview/cmd/repository"
)

type EventService interface {
	Create(ctx context.Context, req models.CreateEventRequest) (*models.Event, error)
	List(ctx context.Context) ([]models.Event, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Event, error)
}

type eventService struct{ repo repository.EventRepository }

func NewEventService(repo repository.EventRepository) EventService {
	return &eventService{repo: repo}
}

func (s *eventService) Create(ctx context.Context, req models.CreateEventRequest) (*models.Event, error) {
	if req.Title == "" || len(req.Title) > 100 {
		return nil, errors.New("invalid title: must be non-empty and <= 100 chars")
	}
	if req.StartTime.IsZero() || req.EndTime.IsZero() || !req.StartTime.Before(req.EndTime) {
		return nil, errors.New("invalid time range: start_time must be before end_time")
	}

	ev := &models.Event{
		ID:          uuid.New(),
		Title:       req.Title,
		Description: req.Description,
		StartTime:   req.StartTime.UTC(),
		EndTime:     req.EndTime.UTC(),
		CreatedAt:   time.Now().UTC(),
	}
	if err := s.repo.Create(ctx, ev); err != nil {
		return nil, err
	}
	return ev, nil
}

func (s *eventService) List(ctx context.Context) ([]models.Event, error) {
	return s.repo.List(ctx)
}

func (s *eventService) GetByID(ctx context.Context, id uuid.UUID) (*models.Event, error) {
	ev, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return ev, nil
}
