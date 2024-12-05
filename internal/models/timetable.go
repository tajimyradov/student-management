package models

type Timetable struct {
	ID         int `json:"id" db:"id"`
	Weekday    int `json:"weekday" db:"weekday"`
	GroupID    int `json:"group_id" db:"group_id"`
	LessonID   int `json:"lesson_id" db:"lesson_id"`
	TimeID     int `json:"time_id" db:"time_id"`
	AuditoryID int `json:"auditory_id" db:"auditory_id"`

	AltLessonID   int `json:"alt_lesson_id" db:"alt_lesson_id"`
	AltAuditoryID int `json:"alt_auditory_id" db:"alt_auditory_id"`
}
