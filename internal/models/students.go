package models

type Student struct {
	ID             int    `json:"id,omitempty" db:"id"`
	FirstName      string `json:"first_name,omitempty" db:"first_name"`
	LastName       string `json:"last_name,omitempty" db:"last_name"`
	MiddleName     string `json:"middle_name" db:"middle_name"`
	Code           string `json:"code,omitempty" db:"code"`
	Gender         bool   `json:"gender" db:"gender"`
	BirthDate      string `json:"birth_date,omitempty" db:"birth_date"`
	Image          string `json:"image,omitempty" db:"image"`
	GroupID        int    `json:"group_id,omitempty" db:"group_id"`
	GroupName      string `json:"group_name,omitempty" db:"group_name"`
	DepartmentID   int    `json:"department_id,omitempty" db:"department_id"`
	DepartmentName string `json:"department_name,omitempty" db:"department_name"`
	FacultyID      int    `json:"faculty_id,omitempty" db:"faculty_id"`
	FacultyName    string `json:"faculty_name,omitempty" db:"faculty_name"`
	Username       string `json:"username,omitempty" db:"username"`
	Password       string `json:"password,omitempty" db:"password"`
	RegionID       int    `json:"region_id" db:"region_id"`
	RegionName     string `json:"region_name" db:"region_name"`
}

type StudentSearch struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Code      string `json:"code" db:"code"`
	BirthDate string `json:"birth_date" db:"birth_date"`
	Username  string `json:"username" db:"username"`
	GroupID   int    `json:"group_id" db:"group_id"`
	Page      int    `json:"page" db:"page"`
	Limit     int    `json:"limit" db:"limit"`
}

type StudentsAndPagination struct {
	Students   []Student  `json:"students"`
	Pagination Pagination `json:"pagination"`
}
