package models

type LessonTeacherStudent struct {
	StudentID int `json:"student_id" db:"student_id"`
	TeacherID int `json:"teacher_id" db:"teacher_id"`
	GroupID   int `json:"group_id" db:"group_id"`
	LessonID  int `json:"lesson_id" db:"lesson_id"`
}
