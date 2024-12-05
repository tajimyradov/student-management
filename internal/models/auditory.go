package models

type Auditory struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type AuditorySearch struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Limit int    `db:"limit" json:"limit"`
	Page  int    `db:"page" json:"page"`
}

type AuditoriesAndPagination struct {
	Auditories []Auditory `json:"auditories"`
	Pagination Pagination `json:"pagination"`
}
