// Package dto contains DTOs for Orderitem.
package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/telemetryflow/order-service/internal/domain/entity"
)

// OrderitemResponse represents the orderitem API response
type OrderitemResponse struct {
	ID        uuid.UUID `json:"id"`
	OrderID   uuid.UUID `json:"order_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// FromOrderitem converts entity to response DTO
func FromOrderitem(e *entity.Orderitem) OrderitemResponse {
	return OrderitemResponse{
		ID:        e.ID,
		OrderID:   e.OrderID,
		ProductID: e.ProductID,
		Quantity:  e.Quantity,
		Price:     e.Price,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

// FromOrderitems converts entities to response DTOs
func FromOrderitems(entities []entity.Orderitem) []OrderitemResponse {
	responses := make([]OrderitemResponse, len(entities))
	for i, e := range entities {
		responses[i] = FromOrderitem(&e)
	}
	return responses
}

// CreateOrderitemRequest represents the create orderitem request
type CreateOrderitemRequest struct {
	OrderID   uuid.UUID `json:"order_id" validate:"required"`
	ProductID uuid.UUID `json:"product_id" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required"`
	Price     float64   `json:"price" validate:"required"`
}

// UpdateOrderitemRequest represents the update orderitem request
type UpdateOrderitemRequest struct {
	OrderID   uuid.UUID `json:"order_id" validate:"required"`
	ProductID uuid.UUID `json:"product_id" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required"`
	Price     float64   `json:"price" validate:"required"`
}

// OrderitemToResponse converts entity pointer to response DTO pointer
func OrderitemToResponse(e *entity.Orderitem) *OrderitemResponse {
	if e == nil {
		return nil
	}
	resp := FromOrderitem(e)
	return &resp
}

// OrderitemListResponse represents the list orderitem API response
type OrderitemListResponse struct {
	Data   []*OrderitemResponse `json:"data"`
	Total  int                  `json:"total"`
	Offset int                  `json:"offset"`
	Limit  int                  `json:"limit"`
}
