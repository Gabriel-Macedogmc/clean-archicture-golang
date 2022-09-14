package productrepository

import (
	"github.com/Gabriel-Macedogmc/clean-archicture-golang/core/domain"
	"github.com/Gabriel-Macedogmc/clean-archicture-golang/core/dto"
	uuid "github.com/satori/go.uuid"
)

func (repository repository) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	product := domain.Product{}
	productRequest.ID = uuid.NewV4()

	err := repository.DB.Table("products").Create(&productRequest).Scan(&product).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}
