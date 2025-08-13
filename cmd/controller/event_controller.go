package controller

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"lkt_test_interview/cmd/models"
	"lkt_test_interview/cmd/services"
	"lkt_test_interview/cmd/utils"
)

const reqTimeout = 3 * time.Second

type EventController struct{ svc services.EventService }

func NewEventController(s services.EventService) *EventController { return &EventController{svc: s} }

var PingController = &pingCtl{}

type pingCtl struct{}

func (*pingCtl) Ping(c *gin.Context) { c.String(http.StatusOK, "pong") }

// Create
// @Summary      Create Event
// @Description  Creates a new event with title, optional description, start_time < end_time. Returns the created event.
// @Tags         events
// @Accept       json
// @Produce      json
// @Param        payload  body      models.CreateEventRequest  true  "Event payload"
// @Success      201      {object}  models.Event
// @Failure      400      {object}  utils.APIError
// @Failure      500      {object}  utils.APIError
// @Router       /api/Events [post]
func (h *EventController) Create(c *gin.Context) {
	var req models.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.APIError{Error: "invalid JSON", Details: err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), reqTimeout)
	defer cancel()

	ev, err := h.svc.Create(ctx, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.APIError{Error: "validation/db error", Details: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ev)
}

// List
// @Summary      List Events
// @Description  Returns all events ordered by start_time ascending.
// @Tags         events
// @Produce      json
// @Success      200  {array}   models.Event
// @Failure      500  {object}  utils.APIError
// @Router       /api/Events [get]
func (h *EventController) List(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), reqTimeout)
	defer cancel()

	events, err := h.svc.List(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.APIError{Error: "db query failed", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

// GetByID
// @Summary      Get Event by ID
// @Description  Returns the event for the given UUID or 404 if not found.
// @Tags         events
// @Produce      json
// @Param        id   path      string  true  "Event UUID"
// @Success      200  {object}  models.Event
// @Failure      400  {object}  utils.APIError
// @Failure      404  {object}  utils.APIError
// @Failure      500  {object}  utils.APIError
// @Router       /api/Events/{id} [get]
func (h *EventController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.APIError{Error: "invalid id", Details: "must be a UUID"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), reqTimeout)
	defer cancel()

	ev, err := h.svc.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, utils.APIError{Error: "not found", Details: id.String()})
			return
		}
		c.JSON(http.StatusInternalServerError, utils.APIError{Error: "db query failed", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, ev)
}
