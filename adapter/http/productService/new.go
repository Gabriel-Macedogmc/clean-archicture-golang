package productservice

import "github.com/Gabriel-Macedogmc/clean-archicture-golang/core/domain"

type service struct {
	useCase domain.ProductUseCase
}

func New(useCase domain.ProductUseCase) domain.ProductService {
	return &service{useCase: useCase}
}
