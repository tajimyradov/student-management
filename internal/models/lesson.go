package models

type Lesson struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Code string `db:"code" json:"code"`
}

type LessonSearch struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Code  string `db:"code" json:"code"`
	Limit int    `db:"limit" json:"limit"`
	Page  int    `db:"page" json:"page"`
}

type LessonsAndPagination struct {
	Lessons    []Lesson   `json:"lessons"`
	Pagination Pagination `json:"pagination"`
}
