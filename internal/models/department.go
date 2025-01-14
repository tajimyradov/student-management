package models

type Department struct {
	ID            int    `json:"id" db:"id"`
	Name          string `json:"name" db:"name"`
	Code          string `json:"code" db:"code"`
	FacultyID     int    `json:"faculty_id" db:"faculty_id"`
	FacultyName   string `json:"faculty_name" db:"faculty_name"`
	Files         []File `json:"files" db:"files"`
	TeachersCount int    `json:"teachers_count" db:"teachers_count"`
}

type DepartmentInfo struct {
	Professions  []Profession `json:"professions" db:"professions"`
	GroupCount   int          `json:"group_count" db:"group_count"`
	StudentCount int          `json:"student_count" db:"student_count"`
	Teachers     []Teacher    `json:"teachers" db:"teachers"`
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
