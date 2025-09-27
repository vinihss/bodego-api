package http_interfaces_drink

type CreateDrinkRequest struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

type UpdateDrinkRequest struct {
	ID    uint    `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}
