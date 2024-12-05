package repositories

import (
	"github.com/jmoiron/sqlx"
	"student-management/internal/models"
)

type TimetableRepository struct {
	studentDB *sqlx.DB
}

func NewTimetableRepository(studentDB *sqlx.DB) *TimetableRepository {
	return &TimetableRepository{
		studentDB: studentDB,
	}
}

func (t *TimetableRepository) AddStudentTeacherLessonBinding(input models.LessonTeacherStudent) error {
	query := `insert into lesson_teacher_student_bindings(lessong_id, teacher_id, student_id,group_id) values ($1,$2,$3,$4)`
	_, err := t.studentDB.Exec(query, input.LessonID, input.TeacherID, input.StudentID, input.GroupID)
	return err
}

func (t *TimetableRepository) DeleteStudentTeacherLessonBinding(input models.LessonTeacherStudent) error {
	query := `delete from lesson_teacher_student_bindings where lessong_id = $1 and teacher_id = $2 and student_id = $3 and group_id = $4`
	_, err := t.studentDB.Exec(query, input.LessonID, input.TeacherID, input.StudentID, input.GroupID)
	return err
}

func (t *TimetableRepository) AddTimetable(input models.Timetable) error {
	query := `insert into timetables(weekday,group_id,lesson_id, time_id, auditory_id,alt_lesson_id, alt_auditory_id) values ($1,$2,$3,$4,$5,$6,$7)`
	_, err := t.studentDB.Exec(query, input.Weekday, input.GroupID, input.LessonID, input.TimeID, input.AuditoryID, input.AltLessonID, input.AltAuditoryID)
	return err
}

func (t *TimetableRepository) DeleteTimetable(id int) error {
	query := `delete from timetables where id=$1`
	_, err := t.studentDB.Exec(query, id)
	return err
}

func (t *TimetableRepository) GetTimetableOfGroup(groupID int) ([]models.Timetable, error) {
	query := `select id,weekday,group_id,lesson_id, time_id, auditory_id,alt_lesson_id, alt_auditory_id from timetables where group_id=$1 order by weekday asc, time_id asc`
	var timetables []models.Timetable
	err := t.studentDB.Select(&timetables, query, groupID)
	return timetables, err
}
