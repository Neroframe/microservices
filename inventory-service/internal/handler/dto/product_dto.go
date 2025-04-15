package dto

// go-playground/validator
type CreateProductRequest struct {
	Name     string  `json:"name" binding:"required"`
	Price    float64 `json:"price" binding:"gt=0"`
	Category string  `json:"category" binding:"required"`
	Stock    int     `json:"stock" binding:"min=0"`
}

type UpdateProductRequest struct {
	Name     *string  `json:"name,omitempty"`
	Price    *float64 `json:"price,omitempty"`
	Category *string  `json:"category,omitempty"`
	Stock    *int     `json:"stock,omitempty"`
}
