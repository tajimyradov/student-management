package repositories

import (
	"student-management/internal/config"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	FacultyRepository    *FacultyRepository
	DepartmentRepository *DepartmentRepository
	ProfessionRepository *ProfessionRepository
	GroupRepository      *GroupRepository
	TeacherRepository    *TeacherRepository
	StudentRepository    *StudentRepository
	AuditoryRepository   *AuditoryRepository
	LessonRepository     *LessonRepository
	TimeRepository       *TimeRepository
	TimetableRepository  *TimetableRepository
}

func NewRepository(studentDB *sqlx.DB, config *config.AppConfig) *Repository {
	return &Repository{
		FacultyRepository:    NewFacultyRepository(studentDB),
		DepartmentRepository: NewDepartmentService(studentDB),
		ProfessionRepository: NewProfessionRepository(studentDB),
		GroupRepository:      NewGroupRepository(studentDB),
		TeacherRepository:    NewTeacherRepository(studentDB),
		StudentRepository:    NewStudentRepository(studentDB),
		AuditoryRepository:   NewAuditoryRepository(studentDB),
		LessonRepository:     NewLessonRepository(studentDB),
		TimeRepository:       NewTimeRepository(studentDB),
		TimetableRepository:  NewTimetableRepository(studentDB),
	}
}
