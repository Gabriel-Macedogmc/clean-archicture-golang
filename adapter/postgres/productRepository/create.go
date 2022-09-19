package productrepository

import (
	"fmt"

	"github.com/Gabriel-Macedogmc/clean-archicture-golang/core/domain"
	"github.com/Gabriel-Macedogmc/clean-archicture-golang/core/dto"
)

func (repository repository) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	product := domain.Product{}
	err := repository.DB.Table("products").Create(productRequest).Scan(&product).Error

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return &product, nil
}
