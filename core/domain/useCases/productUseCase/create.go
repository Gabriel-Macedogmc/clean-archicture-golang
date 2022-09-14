package productusecase

import (
	"errors"

	"github.com/Gabriel-Macedogmc/clean-archicture-golang/core/domain"
	"github.com/Gabriel-Macedogmc/clean-archicture-golang/core/dto"
)

func (useCase useCase) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	productExist := useCase.repository.FindByName(productRequest.Name)

	if productExist != nil {
		return nil, errors.New("product already")
	}

	product, err := useCase.repository.Create(productRequest)

	if err != nil {
		return nil, err
	}

	return product, nil
}
