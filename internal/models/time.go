package models

type Time struct {
	ID        int    `db:"id" json:"id"`
	StartTime string `db:"start_time" json:"start_time"` // Use time.Time if the column type is TIME
	EndTime   string `db:"end_time" json:"end_time"`     // Use time.Time if the column type is TIME
}

type TimeSearch struct {
	ID        int    `db:"id" json:"id"`
	StartTime string `db:"start_time" json:"start_time"`
	EndTime   string `db:"end_time" json:"end_time"`
	Limit     int    `db:"limit" json:"limit"`
	Page      int    `db:"page" json:"page"`
}

type TimesAndPagination struct {
	Times      []Time     `json:"times"`
	Pagination Pagination `json:"pagination"`
}
