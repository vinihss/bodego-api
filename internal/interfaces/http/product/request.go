package http_interfaces_product

type CreateProductRequest struct {
	Name        string  `json:"name"  binding:"required,min=2,max=100"`
	Price       float64 `json:"price" binding:"required,min=0"`
	Description string  `json:"description" binding:"omitempty,max=500"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name"  binding:"omitempty,min=2,max=100"`
	Price       float64 `json:"price" binding:"omitempty,min=0"`
	Description string  `json:"description" binding:"omitempty,max=500"`
}
