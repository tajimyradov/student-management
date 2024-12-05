package models

type Pagination struct {
	TotalPages  int `json:"totalPages"`
	TotalCount  int `json:"totalCount"`
	Limit       int `json:"limit"`
	CurrentPage int `json:"currentPage"`
}
