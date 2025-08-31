package http_interfaces_tab

import "time"

type TabResponse struct {
	ID          uint       `json:"id"`
	UserID      uint       `json:"user_id"`
	OpenDate    time.Time  `json:"open_date"`
	CloseDate   *time.Time `json:"close_date,omitempty"`
	Description string     `json:"description,omitempty"`
	Status      string     `json:"status"`
}
