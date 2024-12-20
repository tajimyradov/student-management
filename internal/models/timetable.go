package models

type Timetable struct {
	ID        int    `json:"id" db:"id"`
	Weekday   int    `json:"weekday" db:"weekday"`
	GroupID   int    `json:"group_id" db:"group_id"`
	GroupName string `json:"group_name" db:"group_name"`
	TimeID    int    `json:"time_id" db:"time_id"`
	StartTime string `json:"start_time" db:"start_time"`
	EndTime   string `json:"end_time" db:"end_time"`
	TypeID    int    `json:"type_id" db:"type_id"`
	TypeName  string `json:"type_name" db:"type_name"`

	Lessons []LessonsList `json:"lessons" db:"lessons"`
}

type LessonsList struct {
	LessonID     int    `json:"lesson_id" db:"lesson_id"`
	LessonName   string `json:"lesson_name" db:"lesson_name"`
	TeacherID    int    `json:"teacher_id" db:"teacher_id"`
	TeacherName  string `json:"teacher_name" db:"teacher_name"`
	AuditoryID   int    `json:"auditory_id" db:"auditory_id"`
	AuditoryName string `json:"auditory_name" db:"auditory_name"`
}
