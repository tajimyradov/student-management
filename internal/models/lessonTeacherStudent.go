package models

type LessonTeacherStudent struct {
	StudentID        int    `json:"student_id" db:"student_id"`
	StudentFirstName string `json:"student_first_name" db:"student_first_name"`
	StudentLastName  string `json:"student_last_name" db:"student_last_name"`
	TeacherID        int    `json:"teacher_id" db:"teacher_id"`
	TeacherFirstName string `json:"teacher_first_name" db:"teacher_first_name"`
	TeacherLastName  string `json:"teacher_last_name" db:"teacher_last_name"`
	GroupID          int    `json:"group_id" db:"group_id"`
	GroupName        string `json:"group_name" db:"group_name"`
	LessonID         int    `json:"lesson_id" db:"lesson_id"`
	LessonName       string `json:"lesson_name" db:"lesson_name"`
	TypeID           int    `json:"type_id" db:"type_id"`
	TypeName         string `json:"type_name" db:"type_name"`
}
