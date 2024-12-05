package models

type Department struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Code      string `json:"code" db:"code"`
	FacultyID int    `json:"faculty_id" db:"faculty_id"`
}

type DepartmentSearch struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	FacultyID int    `json:"faculty_id" db:"faculty_id"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
}

type DepartmentAndPagination struct {
	Departments []Department `json:"departments"`
	Pagination  Pagination   `json:"pagination"`
}
