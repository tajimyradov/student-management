package admin

import (
	"go.uber.org/zap"
	"student-management/internal/config"
	"student-management/internal/repositories/admin"
)

type Service struct {
	FacultyService    *FacultyService
	DepartmentService *DepartmentService
	ProfessionService *ProfessionService
	GroupService      *GroupService
	StudentService    *StudentService
	TeacherService    *TeacherService
	AuditoryService   *AuditoryService
	LessonService     *LessonService
	TimeService       *TimeService
	TimetableService  *TimetableService
}

type ServiceDeps struct {
	Repos *admin.Repository

	Logger *zap.Logger
	Config *config.AppConfig
}

func NewService(deps *ServiceDeps) *Service {
	return &Service{
		FacultyService:    NewFacultyService(deps.Repos.FacultyRepository, deps.Logger),
		DepartmentService: NewDepartmentService(deps.Repos.DepartmentRepository, deps.Logger),
		ProfessionService: NewProfessionService(deps.Repos.ProfessionRepository, deps.Logger),
		GroupService:      NewGroupService(deps.Repos.GroupRepository, deps.Logger),
		StudentService:    NewStudentService(deps.Repos.StudentRepository, deps.Logger, deps.Config),
		TeacherService:    NewTeacherService(deps.Repos.TeacherRepository, deps.Logger, deps.Config),
		AuditoryService:   NewAuditoryService(deps.Repos.AuditoryRepository, deps.Logger),
		LessonService:     NewLessonService(deps.Repos.LessonRepository, deps.Logger),
		TimeService:       NewTimeService(deps.Repos.TimeRepository, deps.Logger),
		TimetableService:  NewTimetableService(deps.Repos.TimetableRepository, deps.Logger),
	}
}
