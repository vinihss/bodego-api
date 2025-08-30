package http_interfaces_tab

type OpenTabRequest struct {
	UserID      uint   `json:"user_id" binding:"required"`
	Description string `json:"description,omitempty"`
}

type UpdateTabRequest struct {
	Description string `json:"description,omitempty"`
}

type CloseTabRequest struct {
	// No additional fields needed - ID comes from URL path
}

// Legacy support - kept for backward compatibility
type CreateTabRequest struct {
	Name  string `json:"name"  binding:"required,min=2,max=100"`
	Email string `json:"email" binding:"required,email,max=254"`
}
