package tab

import (
	"time"
	"github.com/vinihss/bodego-api/internal/domain"
)

type TabStatus string

const (
	TabStatusOpen   TabStatus = "open"
	TabStatusClosed TabStatus = "closed"
)

type Tab struct {
	domain.Entity
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	OpenDate    time.Time `json:"open_date"`
	CloseDate   *time.Time `json:"close_date,omitempty"`
	Description string    `json:"description,omitempty"`
	Status      TabStatus `json:"status"`
}
