// Package dto contains DTOs for Order.
package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/telemetryflow/order-service/internal/domain/entity"
)

// OrderResponse represents the order API response
type OrderResponse struct {
	ID         uuid.UUID `json:"id"`
	CustomerID uuid.UUID `json:"customer_id"`
	Total      float64   `json:"total"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// FromOrder converts entity to response DTO
func FromOrder(e *entity.Order) OrderResponse {
	return OrderResponse{
		ID:         e.ID,
		CustomerID: e.CustomerID,
		Total:      e.Total,
		Status:     e.Status,
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
	}
}

// FromOrders converts entities to response DTOs
func FromOrders(entities []entity.Order) []OrderResponse {
	responses := make([]OrderResponse, len(entities))
	for i, e := range entities {
		responses[i] = FromOrder(&e)
	}
	return responses
}

// CreateOrderRequest represents the create order request
type CreateOrderRequest struct {
	CustomerID uuid.UUID `json:"customer_id" validate:"required"`
	Total      float64   `json:"total" validate:"required"`
	Status     string    `json:"status" validate:"required"`
}

// UpdateOrderRequest represents the update order request
type UpdateOrderRequest struct {
	CustomerID uuid.UUID `json:"customer_id" validate:"required"`
	Total      float64   `json:"total" validate:"required"`
	Status     string    `json:"status" validate:"required"`
}

// OrderToResponse converts entity pointer to response DTO pointer
func OrderToResponse(e *entity.Order) *OrderResponse {
	if e == nil {
		return nil
	}
	resp := FromOrder(e)
	return &resp
}

// OrderListResponse represents the list order API response
type OrderListResponse struct {
	Data   []*OrderResponse `json:"data"`
	Total  int              `json:"total"`
	Offset int              `json:"offset"`
	Limit  int              `json:"limit"`
}
