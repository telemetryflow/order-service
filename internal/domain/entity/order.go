// Package entity contains domain entities.
package entity

import (
	"github.com/google/uuid"
)

// Order represents the order domain entity
type Order struct {
	Base
	CustomerID uuid.UUID   `json:"customer_id" gorm:"type:uuid;not null;index"`
	Total      float64     `json:"total" gorm:"type:decimal(15,2);not null;default:0"`
	Status     string      `json:"status" gorm:"type:varchar(50);not null;default:'pending';index"`
	Items      []Orderitem `json:"items,omitempty" gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// TableName returns the table name for GORM
func (Order) TableName() string {
	return "orders"
}

// NewOrder creates a new Order entity
func NewOrder(customerID uuid.UUID, total float64, status string) *Order {
	return &Order{
		Base:       NewBase(),
		CustomerID: customerID,
		Total:      total,
		Status:     status,
	}
}

// Update updates the order fields
func (e *Order) Update(customerID uuid.UUID, total float64, status string) {
	e.CustomerID = customerID
	e.Total = total
	e.Status = status
	e.MarkUpdated()
}

// Validate validates the entity
func (e *Order) Validate() error {
	// Add validation logic here
	return nil
}
