package productservice

import (
	"encoding/json"
	"net/http"

	"github.com/Gabriel-Macedogmc/clean-archicture-golang/core/dto"
)

func (service service) Create(response http.ResponseWriter, request *http.Request) {
	productRequest, err := dto.FromJSONCreateProductRequest(request.Body)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	product, err := service.useCase.Create(productRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(product)
}
