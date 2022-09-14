package productrepository

import (
	"github.com/Gabriel-Macedogmc/clean-archicture-golang/core/domain"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) domain.ProductRepository {
	return repository{DB: db}
}
