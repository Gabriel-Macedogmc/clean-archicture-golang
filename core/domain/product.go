package domain

import (
	"net/http"
	"time"

	"github.com/Gabriel-Macedogmc/clean-archicture-golang/core/dto"
	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name        string    `gorm:"type:varchar" json:"name"`
	Price       float32   `gorm:"type:float" json:"price"`
	Description string    `gorm:"type:varchar" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime"  json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ProductService interface {
	Create(response http.ResponseWriter, request *http.Request)
	List(response http.ResponseWriter, request *http.Request)
}

type ProductUseCase interface {
	Create(productRequest *dto.CreateProductRequest) (*Product, error)
	List(productRequest *dto.PaginationRequestParams) (*Pagination, error)
}

type ProductRepository interface {
	Create(productRequest *dto.CreateProductRequest) (*Product, error)
	List(productRequest *dto.PaginationRequestParams) (*Pagination, error)
	FindByName(email string) *Product
}
