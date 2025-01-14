package models

type Teacher struct {
	ID             int    `json:"id" db:"id"`
	FirstName      string `json:"first_name" db:"first_name"`
	LastName       string `json:"last_name" db:"last_name"`
	MiddleName     string `json:"middle_name" db:"middle_name"`
	Code           string `json:"code" db:"code"`
	Gender         bool   `json:"gender" db:"gender"`
	Username       string `json:"username" db:"username"`
	Password       string `json:"password" db:"password"`
	GroupID        int    `json:"group_id" db:"group_id"`
	GroupName      string `json:"group_name" db:"group_name"`
	Image          string `json:"image" db:"image"`
	DepartmentId   int    `json:"department_id" db:"department_id"`
	DepartmentName string `json:"department_name" db:"department_name"`
}

type TeacherSearch struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Code         string `json:"code" db:"code"`
	Username     string `json:"username" db:"username"`
	FacultyID    int    `json:"faculty_id" db:"faculty_id"`
	DepartmentId int    `json:"department_id" db:"department_id"`
	Page         int    `json:"page" db:"page"`
	Limit        int    `json:"limit" db:"limit"`
}

type TeachersAndPagination struct {
	Teachers   []Teacher  `json:"teachers"`
	Pagination Pagination `json:"pagination"`
}
