package client

import (
	"github.com/jmoiron/sqlx"
	"student-management/internal/models"
)

type TimetableRepository struct {
	studentDB *sqlx.DB
}

func NewTimetableRepository(db *sqlx.DB) *TimetableRepository {
	return &TimetableRepository{
		studentDB: db,
	}
}

func (t *TimetableRepository) GetTimetableOfStudent(userID int, weekday int) ([]models.Timetable, error) {
	var timetables []models.Timetable
	query := `select t.weekday,
				t.group_id,
				t.group_name,
				t.lesson_id,
				t.lesson_name,
				t.time_id,
				t.start_time,
				t.end_time,
				t.auditory_id,
				t.auditory_name,
				t.alt_lesson_id,
				t.alt_lesson_name,
				t.alt_auditory_id,
				t.alt_auditory_name,
				t.type_id,
				t.type_name,
			from timetable_view as t where  t.group_id=(select group_id from students where id=$1)  and t.weekday=$2  order by t.time_id`
	err := t.studentDB.Select(&timetables, query, userID, weekday)
	return timetables, err
}

func (t *TimetableRepository) GetTimetableOfTeacher(userID int, weekday int) ([]models.Timetable, error) {
	var timetables []models.Timetable
	query := `select distinct t.weekday,
				t.group_id,
				t.group_name,
				t.lesson_id,
				t.lesson_name,
				t.time_id,
				t.start_time,
				t.end_time,
				t.auditory_id,
				t.auditory_name,
				t.alt_lesson_id,
				t.alt_lesson_name,
				t.alt_auditory_id,
				t.alt_auditory_name,
				t.type_id,
				t.type_name,
			from timetable_view as t join lesson_teacher_student_bindings as ltsb on ltsb.group_id=t.group_id and ltsb.lesson_id=t.lesson_id and ltsb.type_id=t.type_id where ltsb.teacher_id=$1 and t.weekday=$2`
	err := t.studentDB.Select(&timetables, query, userID, weekday)
	return timetables, err
}
