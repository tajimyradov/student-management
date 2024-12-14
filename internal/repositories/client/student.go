package client

import (
	"github.com/jmoiron/sqlx"
	"student-management/internal/models"
	"time"
)

type StudentRepository struct {
	studentDB *sqlx.DB
}

func NewStudentRepository(db *sqlx.DB) *StudentRepository {
	return &StudentRepository{
		studentDB: db,
	}
}

func (s *StudentRepository) GetStudentsForTeacher(groupID int) ([]models.Student, error) {
	query := `select s.id,s.first_name, s.last_name from students as s where s.group_id=$1`
	var students []models.Student
	err := s.studentDB.Select(&students, query, groupID)
	return students, err
}

func (s *StudentRepository) CheckForAbsence(input models.Absence) error {
	query := `insert into absences(student_id,group_id, lesson_id,time_id, teacher_id, type_id, date,status)`
	_, err := s.studentDB.Exec(query, input.StudentID, input.GroupID, input.LessonID, input.TimeID, input.TeacherID, input.TypeID, time.Now().Format(time.DateOnly), 1)
	if err != nil {
		return err
	}
	return nil
}

func (s *StudentRepository) GetFaculties() ([]models.Faculty, error) {
	query := `select id,name from faculties`
	var faculties []models.Faculty
	err := s.studentDB.Select(&faculties, query)
	return faculties, err
}

func (s *StudentRepository) GetDepartments(facultyID int) ([]models.Department, error) {
	query := `select id,name from departments where faculty_id=$1`
	var departments []models.Department
	err := s.studentDB.Select(&departments, query, facultyID)
	return departments, err
}

func (s *StudentRepository) GetGroups(departmentID int) ([]models.Group, error) {
	query := `select id,name from groups where profession_id in (select id from professions where department_id=$1)`
	var groups []models.Group
	err := s.studentDB.Select(&groups, query, departmentID)
	return groups, err
}

func (s *StudentRepository) GetLessons() ([]models.Lesson, error) {
	query := `select id,name from lessons order by name`
	var lessons []models.Lesson
	err := s.studentDB.Select(&lessons, query)
	return lessons, err
}

func (s *StudentRepository) GetTypes() ([]models.Type, error) {
	query := `select id,name from types`
	var types []models.Type
	err := s.studentDB.Select(&types, query)
	return types, err
}

func (s *StudentRepository) GetTimes() ([]models.Time, error) {
	query := `select id,name from times`
	var times []models.Time
	err := s.studentDB.Select(&times, query)
	return times, err
}
