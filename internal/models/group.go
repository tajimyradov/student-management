package models

type Group struct {
	ID             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	Code           string `json:"code" db:"code"`
	Year           int    `json:"year" db:"year"`
	ProfessionID   int    `json:"profession_id" db:"profession_id"`
	ProfessionName string `json:"profession_name" db:"profession_name"`
}

type GroupSearch struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Code         string `json:"code" db:"code"`
	Year         int    `json:"year" db:"year"`
	ProfessionID int    `json:"profession_id" db:"profession_id"`
	Limit        int    `json:"limit" db:"limit"`
	Page         int    `json:"page" db:"page"`
}

type GroupsAndPagination struct {
	Groups     []Group    `json:"groups"`
	Pagination Pagination `json:"pagination"`
}
