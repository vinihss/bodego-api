package tab

import (
	"github.com/vinihss/bodego-api/internal/domain"
	"time"
)

type TabStatus string

const (
	TabStatusOpen   TabStatus = "open"
	TabStatusClosed TabStatus = "closed"
)

type Tab struct {
	domain.Entity
	ID          uint       `json:"id"`
	UserID      uint       `json:"user_id"`
	OpenDate    time.Time  `json:"open_date"`
	CloseDate   *time.Time `json:"close_date,omitempty"`
	Description string     `json:"description,omitempty"`
	Status      TabStatus  `json:"status"`
	Items       []Item     `json:"items,omitempty"`
}

type Item struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
