package productrepository

import (
	"github.com/Gabriel-Macedogmc/clean-archicture-golang/core/domain"
	"github.com/Gabriel-Macedogmc/clean-archicture-golang/core/dto"
	"github.com/booscaaa/go-paginate/paginate"
)

func (repository repository) List(pagination *dto.PaginationRequestParams) (*domain.Pagination, error) {
	products := []domain.Product{}
	total := int32(0)

	pagin := paginate.Instance(pagination)
	query, queryCount := pagin.Query("SELECT * FROM products").
		Page(pagination.Page).
		Desc(pagination.Descending).
		Sort(pagination.Sort).
		RowsPerPage(pagination.ItemsPerPage).
		SearchBy(pagination.Search, "name", "description").
		Select()

	{
		rows, err := repository.DB.Table("products").Raw(*query).Rows()

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			product := domain.Product{}

			rows.Scan(
				&product.ID,
				&product.Name,
				&product.Price,
				&product.Description,
				&product.CreatedAt,
				&product.UpdatedAt,
			)

			products = append(products, product)
		}
	}
	err := repository.DB.Table("products").Raw(*queryCount).Scan(&total).Error

	if err != nil {
		return nil, err
	}

	return &domain.Pagination{
		Items: products,
		Total: total,
	}, nil

}
