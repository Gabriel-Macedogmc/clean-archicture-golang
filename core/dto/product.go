package dto

import (
	"encoding/json"
	"io"
	"time"

	uuid "github.com/satori/go.uuid"
)

type CreateProductRequest struct {
	ID          uuid.UUID
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func FromJSONCreateProductRequest(body io.Reader) (*CreateProductRequest, error) {
	createProductRequest := CreateProductRequest{}

	if err := json.NewDecoder(body).Decode(&createProductRequest); err != nil {
		return nil, err
	}

	return &createProductRequest, nil
}
