// Package command contains CQRS commands for Order.
package command

import (
	"github.com/google/uuid"
	"github.com/telemetryflow/order-service/internal/domain/entity"
)

// CreateOrderCommand represents the create order command
type CreateOrderCommand struct {
	CustomerID uuid.UUID `json:"customer_id" validate:"required"`
	Total      float64   `json:"total" validate:"required"`
	Status     string    `json:"status" validate:"required"`
}

// Validate validates the create command
func (c *CreateOrderCommand) Validate() error {
	// Add validation logic
	return nil
}

// ToEntity converts the command to an entity
func (c *CreateOrderCommand) ToEntity() *entity.Order {
	return entity.NewOrder(c.CustomerID, c.Total, c.Status)
}

// UpdateOrderCommand represents the update order command
type UpdateOrderCommand struct {
	ID         uuid.UUID `json:"id" validate:"required"`
	CustomerID uuid.UUID `json:"customer_id" validate:"required"`
	Total      float64   `json:"total" validate:"required"`
	Status     string    `json:"status" validate:"required"`
}

// Validate validates the update command
func (c *UpdateOrderCommand) Validate() error {
	if c.ID == uuid.Nil {
		return ErrInvalidID
	}
	return nil
}

// ToEntity converts the command to an entity
func (c *UpdateOrderCommand) ToEntity() *entity.Order {
	e := entity.NewOrder(c.CustomerID, c.Total, c.Status)
	e.ID = c.ID
	return e
}

// DeleteOrderCommand represents the delete order command
type DeleteOrderCommand struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

// Validate validates the delete command
func (c *DeleteOrderCommand) Validate() error {
	if c.ID == uuid.Nil {
		return ErrInvalidID
	}
	return nil
}
