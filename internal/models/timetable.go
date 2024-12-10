package models

type Timetable struct {
	ID              int    `json:"id" db:"id"`
	Weekday         int    `json:"weekday" db:"weekday"`
	GroupID         int    `json:"group_id" db:"group_id"`
	GroupName       string `json:"group_name" db:"group_name"`
	LessonID        int    `json:"lesson_id" db:"lesson_id"`
	LessonName      string `json:"lesson_name" db:"lesson_name"`
	TimeID          int    `json:"time_id" db:"time_id"`
	StartTime       string `json:"start_time" db:"start_time"`
	EndTime         string `json:"end_time" db:"end_time"`
	AuditoryID      int    `json:"auditory_id" db:"auditory_id"`
	AuditoryName    string `json:"auditory_name" db:"auditory_name"`
	AltLessonID     int    `json:"alt_lesson_id" db:"alt_lesson_id"`
	AltLessonName   string `json:"alt_lesson_name" db:"alt_lesson_name"`
	AltAuditoryID   int    `json:"alt_auditory_id" db:"alt_auditory_id"`
	AltAuditoryName string `json:"alt_auditory_name" db:"alt_auditory_name"`
	TypeID          int    `json:"type_id" db:"type_id"`
	TypeName        string `json:"type_name" db:"type_name"`
}
