package admin

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"student-management/internal/config"
	"student-management/internal/services/admin"
)

type Admin struct {
	services *admin.Service
	logger   *zap.Logger
	config   *config.AppConfig
}

func NewHandler(services *admin.Service, logger *zap.Logger, config *config.AppConfig) *Admin {
	return &Admin{
		services: services,
		logger:   logger,
		config:   config,
	}
}

func (h *Admin) Init(admin *gin.RouterGroup) {

	faculty := admin.Group("/faculty")
	{
		faculty.POST("", h.addFaculty)
		faculty.PUT("/:fid", h.updateFaculty)
		faculty.DELETE("/:fid", h.deleteFaculty)
		faculty.GET("/:fid", h.getFacultyByID)
		faculty.GET("", h.getFaculties)
	}

	department := admin.Group("/department")
	{
		department.POST("", h.addDepartment)
		department.PUT("/:did", h.updateDepartment)
		department.DELETE("/:did", h.deleteDepartment)
		department.GET("/:did", h.getDepartmentByID)
		department.GET("", h.getDepartments)
	}

	profession := admin.Group("/profession")
	{
		profession.POST("", h.addProfession)
		profession.PUT("/:pid", h.updateProfession)
		profession.DELETE("/:pid", h.deleteProfession)
		profession.GET("/:pid", h.getProfessionByID)
		profession.GET("", h.getProfessions)
	}

	group := admin.Group("/group")
	{
		group.POST("", h.addGroup)
		group.PUT("/:gid", h.updateGroup)
		group.DELETE("/:gid", h.deleteGroup)
		group.GET("/:gid", h.getGroupByID)
		group.GET("", h.getGroups)
	}

	student := admin.Group("/student")
	{
		student.POST("", h.addStudent)
		student.PUT("/:sid", h.updateStudent)
		student.DELETE("/:sid", h.deleteStudent)
		student.GET("/:sid", h.getStudentByID)
		student.GET("", h.getStudents)
		student.POST("/:sid/upload-image", h.uploadStudentImage)
	}

	teacher := admin.Group("/teacher")
	{
		teacher.POST("", h.addTeacher)
		teacher.PUT("/:tid", h.updateTeacher)
		teacher.DELETE("/:tid", h.deleteTeacher)
		teacher.GET("/:tid", h.getTeacherByID)
		teacher.GET("", h.getTeachers)
		teacher.POST("/:tid/upload-image", h.uploadTeacherImage)
	}

	auditory := admin.Group("/auditory")
	{
		auditory.POST("", h.addAuditory)
		auditory.PUT("/:aid", h.updateAuditory)
		auditory.DELETE("/:aid", h.deleteAuditory)
		auditory.GET("/:aid", h.getAuditoryByID)
		auditory.GET("", h.getAuditories)
	}

	lesson := admin.Group("/lesson")
	{
		lesson.POST("", h.addLesson)
		lesson.PUT("/:lid", h.updateLesson)
		lesson.DELETE("/:lid", h.deleteLesson)
		lesson.GET("/:lid", h.getLessonByID)
		lesson.GET("", h.getLessons)
	}

	time := admin.Group("/time")
	{
		time.POST("", h.addTime)
		time.PUT("/:tid", h.updateTime)
		time.DELETE("/:tid", h.deleteTime)
		time.GET("/:tid", h.getTimeByID)
		time.GET("", h.getTimes)
	}

	studentLessonTeacherBinding := admin.Group("/bind")
	{
		studentLessonTeacherBinding.GET("", h.getLessonTeacherStudentBinding)
		studentLessonTeacherBinding.POST("", h.addLessonTeacherStudentBinding)
		studentLessonTeacherBinding.DELETE("", h.deleteLessonTeacherStudentBinding)

	}

	timetable := admin.Group("/timetable")
	{
		timetable.POST("", h.addTimetable)
		timetable.DELETE("/:timetable_id", h.deleteTimetable)
		timetable.GET("/:group_id", h.getTimetableOfGroup)
	}

}
