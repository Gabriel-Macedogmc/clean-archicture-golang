package productusecase

import "github.com/Gabriel-Macedogmc/clean-archicture-golang/core/domain"

type useCase struct {
	repository domain.ProductRepository
}

func New(repository domain.ProductRepository) domain.ProductUseCase {
	return &useCase{repository: repository}
}
