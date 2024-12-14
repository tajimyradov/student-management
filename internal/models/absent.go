package models

type Absence struct {
	GroupID   int `json:"group_id" db:"group_id"`
	LessonID  int `json:"lesson_id" db:"lesson_id"`
	TimeID    int `json:"time_id" db:"time_id"`
	TeacherID int `json:"teacher_id" db:"teacher_id"`
	TypeID    int `json:"type_id" db:"type_id"`
	StudentID int `json:"student_id" db:"student_id"`
}
