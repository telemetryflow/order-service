// Package entity contains domain entities.
package entity

import (
	"github.com/google/uuid"
)

// Orderitem represents the orderitem domain entity
type Orderitem struct {
	Base
	OrderID   uuid.UUID `json:"order_id" gorm:"type:uuid;not null;index"`
	ProductID uuid.UUID `json:"product_id" gorm:"type:uuid;not null;index"`
	Quantity  int       `json:"quantity" gorm:"not null;default:1"`
	Price     float64   `json:"price" gorm:"type:decimal(15,2);not null;default:0"`
}

// TableName returns the table name for GORM
func (Orderitem) TableName() string {
	return "order_items"
}

// NewOrderitem creates a new Orderitem entity
func NewOrderitem(orderID uuid.UUID, productID uuid.UUID, quantity int, price float64) *Orderitem {
	return &Orderitem{
		Base:      NewBase(),
		OrderID:   orderID,
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
	}
}

// Update updates the orderitem fields
func (e *Orderitem) Update(orderID uuid.UUID, productID uuid.UUID, quantity int, price float64) {
	e.OrderID = orderID
	e.ProductID = productID
	e.Quantity = quantity
	e.Price = price
	e.MarkUpdated()
}

// Validate validates the entity
func (e *Orderitem) Validate() error {
	// Add validation logic here
	return nil
}
