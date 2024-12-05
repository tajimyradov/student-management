package models

type Student struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Code      string `json:"code" db:"code"`
	Gender    bool   `json:"gender" db:"gender"`
	BirthDate string `json:"birth_date" db:"birth_date"`
	GroupID   int    `json:"group_id" db:"group_id"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
}

type StudentSearch struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
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
