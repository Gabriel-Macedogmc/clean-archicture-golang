package productservice

import (
	"encoding/json"
	"net/http"

	"github.com/Gabriel-Macedogmc/clean-archicture-golang/core/dto"
)

func (service service) List(response http.ResponseWriter, request *http.Request) {
	paginationRequest, err := dto.FromValuePaginationRequestParams(request)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	products, err := service.useCase.List(paginationRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(products)
}
