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

func (s *StudentRepository) GetStudentsForTeacher(teacherID, lessonID, groupID, typeID int) ([]models.Student, error) {
	query := `select s.id,s.first_name, s.last_name from students as s join lesson_teacher_student_bindings as ltsb on ltsb.student_id = s.id where ltsb.teacher_id=$1 and ltsb.lesson_id=$2 and ltsb.group_id=$3 and ltsb.type_id=$4`
	var students []models.Student
	err := s.studentDB.Select(&students, query, teacherID, lessonID, groupID, typeID)
	return students, err
}

func (s *StudentRepository) CheckForAbsence(input models.Absence) error {
	tx, err := s.studentDB.Begin()
	if err != nil {
		return err
	}

	for _, student := range input.Students {
		query := `insert into absences(student_id,group_id, lesson_id,time_id, teacher_id, type_id, date,status,is_absent)`
		_, err = tx.Exec(query, student.StudentID, input.GroupID, input.LessonID, input.TimeID, input.TeacherID, input.TypeID, time.Now().Format(time.DateOnly), 1, student.IsAbsence)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
