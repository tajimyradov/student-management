package models

type Faculty struct {
	ID              int    `json:"id" db:"id"`
	Name            string `json:"name" db:"name"`
	Code            string `json:"code" db:"code"`
	DepartmentCount int    `json:"department_count" db:"department_count"`
}

type FacultySearch struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name"`
	Code  string `json:"code"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}

type FacultiesAndPagination struct {
	Faculties  []Faculty  `json:"faculties"`
	Pagination Pagination `json:"pagination"`
}
