package productrepository_test

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	productrepository "github.com/Gabriel-Macedogmc/clean-archicture-golang/adapter/postgres/productRepository"
	"github.com/Gabriel-Macedogmc/clean-archicture-golang/core/domain"
	"github.com/Gabriel-Macedogmc/clean-archicture-golang/core/dto"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type v2Suite struct {
	db   *gorm.DB
	mock sqlmock.Sqlmock
}

func setupCreate(t *testing.T) ([]string, dto.CreateProductRequest, domain.Product, sqlmock.Sqlmock, *gorm.DB) {
	s := &v2Suite{}
	cols := []string{"id", "name", "price", "description", "created_at", "updated_at"}
	fakeProductRequest := dto.CreateProductRequest{
		Name:        "test",
		Price:       10,
		Description: "test description",
	}
	fakeProductDBResponse := domain.Product{}
	//faker.FakeData(&fakeProductDBResponse)

	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()

	if err != nil {
		log.Fatalf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	s.db, err = gorm.Open(dialector, &gorm.Config{})

	return cols, fakeProductRequest, fakeProductDBResponse, s.mock, s.db
}

func TestCreate(t *testing.T) {
	cols, fakeProductRequest, fakeProductDBResponse, mock, db := setupCreate(t)
	qStr := `INSERT  INTO "products" ("name","price","description") 
	VALUES ($1,$2,$3) RETURNING "products"."id","products"."name","products"."price","products"."description","products"."created_at","products"."updated_at"`

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(qStr)).WithArgs(
		fakeProductRequest.Name,
		fakeProductRequest.Price,
		fakeProductRequest.Description,
	).WillReturnRows(mock.NewRows(cols).AddRow(
		fakeProductDBResponse.ID,
		fakeProductDBResponse.Name,
		fakeProductDBResponse.Price,
		fakeProductDBResponse.Description,
		fakeProductDBResponse.CreatedAt,
		fakeProductDBResponse.UpdatedAt,
	))
	mock.ExpectCommit()

	t.Logf("REQUEST REPO: %v", fakeProductDBResponse)

	sut := productrepository.New(db)
	product, err := sut.Create(&fakeProductRequest)

	t.Logf("PRODUCT REPOSITORY: %v", product)
	t.Logf("ERRO REPOSITORY: %v", err.Error())

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("havia expectativas não cumpridas: %v", err)
	}

	// require.Nil(t, err)
	//require.NotEmpty(t, product.ID)
	// require.Empty(t, product)
	// require.Equal(t, product.Name, fakeProductDBResponse.Name)
	// require.Equal(t, product.Price, fakeProductDBResponse.Price)
	// require.Equal(t, product.Description, fakeProductDBResponse.Description)

}

func TestCreate_DBError(t *testing.T) {
	_, fakerProductRequest, _, mock, db := setupCreate(t)

	mock.ExpectQuery("INSERT INTO products (.+)").WithArgs(
		fakerProductRequest.Name,
		fakerProductRequest.Price,
		fakerProductRequest.Description,
	).WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

	sut := productrepository.New(db)
	product, err := sut.Create(&fakerProductRequest)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("havia expectativas não cumpridas: %v", err)
	}

	require.NotNil(t, err)
	require.Nil(t, product)
}
