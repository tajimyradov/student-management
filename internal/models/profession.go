package models

type Profession struct {
	ID             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	Code           string `json:"code" db:"code"`
	DepartmentID   int    `json:"department_id" db:"department_id"`
	DepartmentName string `json:"department_name" db:"department_name"`
}

type ProfessionSearch struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	DepartmentID int    `json:"department_id" db:"department_id"`
	Limit        int    `json:"limit"`
	Page         int    `json:"page"`
}

type ProfessionAndPagination struct {
	Professions []Profession `json:"professions"`
	Pagination  Pagination   `json:"pagination"`
}
