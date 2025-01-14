package models

type Faculty struct {
	ID              int    `json:"id" db:"id"`
	Position        int    `json:"position" db:"position"`
	Name            string `json:"name" db:"name"`
	Code            string `json:"code" db:"code"`
	Files           []File `json:"files" db:"files"`
	DepartmentCount int    `json:"department_count" db:"department_count"`
}

type FacultyInfo struct {
	Professions  []Profession `json:"professions" db:"professions"`
	Departments  []Department `json:"departments" db:"departments"`
	GroupCount   int          `json:"group_count" db:"group_count"`
	StudentCount int          `json:"student_count" db:"student_count"`
	Teachers     []Teacher    `json:"teachers" db:"teachers"`
}

type File struct {
	ID      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	FileURL string `json:"file_url" db:"file_url"`
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
