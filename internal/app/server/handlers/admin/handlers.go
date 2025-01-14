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
		faculty.POST("/:fid/file", h.uploadFileOfFaculty)
		faculty.DELETE("/file/:fid", h.deleteFileOfFaculty)
		faculty.GET("/:fid/info", h.getFacultyInfo)
	}

	department := admin.Group("/department")
	{
		department.POST("", h.addDepartment)
		department.PUT("/:did", h.updateDepartment)
		department.DELETE("/:did", h.deleteDepartment)
		department.GET("/:did", h.getDepartmentByID)
		department.GET("", h.getDepartments)
		department.POST("/:did/file", h.uploadFileOfDepartment)
		department.DELETE("/file/:did", h.deleteFileOfDepartment)
		department.GET("/:did/info", h.getDepartmentInfo)
	}

	profession := admin.Group("/profession")
	{
		profession.POST("", h.addProfession)
		profession.PUT("/:pid", h.updateProfession)
		profession.DELETE("/:pid", h.deleteProfession)
		profession.GET("/:pid", h.getProfessionByID)
		profession.GET("", h.getProfessions)
		profession.POST("/:pid/file", h.uploadFileOfProfession)
		profession.DELETE("/file/:pid", h.deleteFileOfProfession)
		profession.GET("/:pid/info", h.getProfessionInfo)
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

	timetable := admin.Group("/timetable")
	{
		timetable.POST("", h.addTimetable)
		timetable.DELETE("/:timetable_id", h.deleteTimetable)
		timetable.GET("/:group_id", h.getTimetableOfGroup)
	}

	absence := admin.Group("/absence")
	{
		absence.POST("", h.getAbsences)
		absence.POST("/:absence_id", h.updateAbsences)
		absence.GET("/:absence_id", h.getAbsenceByID)
		absence.POST("/sync", h.SyncAbsence)
	}

	employeeRate := admin.Group("/employee")
	{
		employeeRate.GET("/rate", h.getEmployeeRate)
		employeeRate.GET("/rate/:emp_id", h.getEmployeeRateByID)
		employeeRate.POST("/rate", h.addEmployeeRate)
		employeeRate.DELETE("/rate/:emp_id", h.deleteEmployeeRate)
		employeeRate.PUT("/rate/:emp_id", h.updateEmployeeRate)
	}

	stats := admin.Group("/statistics")
	{
		stats.GET("/:faculty_id/gender", h.getStatisticsByGender)
		stats.GET("/:faculty_id/profession", h.getStatisticsOfProfession)
		stats.GET("/:faculty_id/age", h.getStatisticsByAge)
		stats.GET("/:faculty_id/region", h.getStatisticsByRegions)

	}

	admin.GET("/positions", h.getPositions)

	admin.GET("/regions", h.getRegions)
}
