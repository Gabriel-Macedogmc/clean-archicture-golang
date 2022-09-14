package dto

import (
	"net/http"
	"strconv"
	"strings"
)

type PaginationRequestParams struct {
	Search       string   `json:"search"`
	Descending   []string `json:"descending"`
	Page         int      `json:"page"`
	ItemsPerPage int      `json:"itemsPerPage"`
	Sort         []string `json:"sort"`
}

func FromValuePaginationRequestParams(request *http.Request) (*PaginationRequestParams, error) {
	page, _ := strconv.Atoi(request.FormValue("page"))
	itemsPerPage, _ := strconv.Atoi(request.FormValue("itemsPerPage"))

	paginationRequestParams := PaginationRequestParams{
		Search:       request.FormValue("search"),
		Page:         page,
		ItemsPerPage: itemsPerPage,
		Descending:   strings.Split(request.FormValue("descending"), ","),
		Sort:         strings.Split(request.FormValue("sort"), ","),
	}

	return &paginationRequestParams, nil
}
