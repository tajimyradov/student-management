package models

type Profession struct {
	ID             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	Code           string `json:"code" db:"code"`
	DepartmentID   int    `json:"department_id" db:"department_id"`
	Files          []File `json:"files" db:"files"`
	DepartmentName string `json:"department_name" db:"department_name"`
	GroupCount     int    `json:"group_count" db:"group_count"`
}

type ProfessionInfo struct {
	Groups       []Group `json:"groups" db:"groups"`
	Files        []File  `json:"files" db:"files"`
	StudentCount int     `json:"student_count" db:"student_count"`
}

type ProfessionSearch struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	FacultyID    int    `json:"faculty_id" db:"faculty_id"`
	DepartmentID int    `json:"department_id" db:"department_id"`
	Limit        int    `json:"limit"`
	Page         int    `json:"page"`
}

type ProfessionAndPagination struct {
	Professions []Profession `json:"professions"`
	Pagination  Pagination   `json:"pagination"`
}
