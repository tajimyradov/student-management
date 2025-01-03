package models

type Group struct {
	ID               int    `json:"id" db:"id"`
	Name             string `json:"name" db:"name"`
	Code             string `json:"code" db:"code"`
	Year             int    `json:"year" db:"year"`
	ProfessionID     int    `json:"profession_id" db:"profession_id"`
	ProfessionName   string `json:"profession_name" db:"profession_name"`
	StudentCount     int    `json:"student_count" db:"student_count"`
	TeacherID        int    `json:"teacher_id" db:"teacher_id"`
	TeacherFirstName string `json:"teacher_first_name" db:"teacher_first_name"`
	TeacherLastName  string `json:"teacher_last_name" db:"teacher_last_name"`
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
