package productrepository

import "github.com/Gabriel-Macedogmc/clean-archicture-golang/core/domain"

func (repository repository) FindByName(email string) *domain.Product {
	product := domain.Product{}

	repository.DB.Table("products").Where("name = ?", email).Find(&product)

	return &product
}
