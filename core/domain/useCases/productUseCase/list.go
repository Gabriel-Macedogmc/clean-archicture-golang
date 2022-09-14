package productusecase

import (
	"github.com/Gabriel-Macedogmc/clean-archicture-golang/core/domain"
	"github.com/Gabriel-Macedogmc/clean-archicture-golang/core/dto"
)

func (useCase useCase) List(paginationRequest *dto.PaginationRequestParams) (*domain.Pagination, error) {
	products, err := useCase.repository.List(paginationRequest)

	if err != nil {
		return nil, err
	}

	return products, nil
}
