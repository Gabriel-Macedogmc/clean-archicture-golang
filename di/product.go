package di

import (
	productservice "github.com/Gabriel-Macedogmc/clean-archicture-golang/adapter/http/productService"
	productrepository "github.com/Gabriel-Macedogmc/clean-archicture-golang/adapter/postgres/productRepository"
	"github.com/Gabriel-Macedogmc/clean-archicture-golang/core/domain"
	productusecase "github.com/Gabriel-Macedogmc/clean-archicture-golang/core/domain/useCases/productUseCase"
	"gorm.io/gorm"
)

func ConfigProductDI(conn *gorm.DB) domain.ProductService {
	productRepository := productrepository.New(conn)
	productUseCase := productusecase.New(productRepository)
	productService := productservice.New(productUseCase)

	return productService
}
