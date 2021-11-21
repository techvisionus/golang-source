package contants

// ProductRequest input
type ProductRequest struct {
	Code string `json:"code" binding:"required,min=1"`
	Name string `json:"name" binding:"required,min=1"`
	Price int64 `json:"price" binding:"required,gt=0"`
}

// GetProductRequest input
type GetProductRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}