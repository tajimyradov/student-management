package models

type Absence struct {
	ID                      int    `json:"id" db:"id"`
	GroupID                 int    `json:"group_id" db:"group_id"`
	GroupName               string `json:"group_name" db:"group_name"`
	LessonID                int    `json:"lesson_id" db:"lesson_id"`
	LessonName              string `json:"lesson_name" db:"lesson_name"`
	TimeID                  int    `json:"time_id" db:"time_id"`
	StartTime               string `json:"start_time" db:"start_time"`
	EndTime                 string `json:"end_time" db:"end_time"`
	TeacherID               int    `json:"teacher_id" db:"teacher_id"`
	FacultyId               int    `json:"faculty_id" db:"faculty_id"`
	FacultyName             string `json:"faculty_name" db:"faculty_name"`
	DepartmentID            int    `json:"department_id" db:"department_id"`
	DepartmentName          string `json:"department_name" db:"department_name"`
	TeacherFirstName        string `json:"teacher_first_name" db:"teacher_first_name"`
	TeacherLastName         string `json:"teacher_last_name" db:"teacher_last_name"`
	StudentYear             int    `json:"student_year" db:"student_year"`
	TypeID                  int    `json:"type_id" db:"type_id"`
	TypeName                string `json:"type_name" db:"type_name"`
	StudentID               int    `json:"student_id" db:"student_id"`
	StudentFirstName        string `json:"student_first_name" db:"student_first_name"`
	StudentLastName         string `json:"student_last_name" db:"student_last_name"`
	FacultyDeanFirstName    string `json:"faculty_dean_first_name" db:"faculty_dean_first_name"`
	FacultyDeanLastName     string `json:"faculty_dean_last_name" db:"faculty_dean_last_name"`
	DepartmentLeadFirstName string `json:"department_lead_first_name" db:"department_lead_first_name"`
	DepartmentLeadLastName  string `json:"department_lead_last_name" db:"department_lead_last_name"`
	Note                    string `json:"note" db:"note"`
	Status                  int    `json:"status" db:"status"`
	ProfessionID            int    `json:"profession_id" db:"profession_id"`
	ProfessionName          string `json:"profession_name" db:"profession_name"`
	Date                    string `json:"date" db:"date"`
}

type AbsenceSearch struct {
	FacultyID        int    `json:"faculty_id" db:"faculty_id"`
	DepartmentID     int    `json:"department_id" db:"department_id"`
	GroupID          int    `json:"group_id" db:"group_id"`
	LessonID         int    `json:"lesson_id" db:"lesson_id"`
	TypeID           int    `json:"type_id" db:"type_id"`
	TeacherID        int    `json:"teacher_id" db:"teacher_id"`
	StudentID        int    `json:"student_id" db:"student_id"`
	StudentFirstName string `json:"student_first_name" db:"student_first_name"`
	StudentLastName  string `json:"student_last_name" db:"student_last_name"`
	From             string `json:"from" db:"from"`
	To               string `json:"to" db:"to"`
}

type AbsenceAndPagination struct {
	Pagination Pagination `json:"pagination"`
	Absences   []Absence  `json:"absences"`
}
