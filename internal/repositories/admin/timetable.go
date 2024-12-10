package admin

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
	query := `insert into lesson_teacher_student_bindings(lesson_id, teacher_id, student_id,group_id,type_id) values ($1,$2,$3,$4,$5)`
	_, err := t.studentDB.Exec(query, input.LessonID, input.TeacherID, input.StudentID, input.GroupID, input.TypeID)
	return err
}

func (t *TimetableRepository) DeleteStudentTeacherLessonBinding(input models.LessonTeacherStudent) error {
	query := `delete from lesson_teacher_student_bindings where lessong_id = $1 and teacher_id = $2 and student_id = $3 and group_id = $4 and type_id = $5`
	_, err := t.studentDB.Exec(query, input.LessonID, input.TeacherID, input.StudentID, input.GroupID, input.TypeID)
	return err
}

func (t *TimetableRepository) GetStudentTeacherLessonBinding(teacherID, lessonID int) (models.LessonTeacherStudent, error) {
	query := `select ltsb.lesson_id,
		   ltsb.teacher_id,
		   ltsb.student_id,
		   ltsb.group_id,
		   ltsb.type_id,
		   l.name       as lesson_name,
		   t.first_name as teacher_first_name,
		   t.last_name  as teacher_last_name,
		   s.first_name as student_first_name,
		   s.last_name  as student_last_name,
		   g.name       as group_name,
		   ty.name as type_name
		from lesson_teacher_student_bindings as ltsb
				 join lessons as l on l.id = ltsb.lesson_id
				 join teachers as t on t.id = ltsb.teacher_id
				 join students as s on s.id = ltsb.student_id
				 join groups as g on ltsb.group_id = g.id
				 join types as ty on ty.id = ltsb.type_id
				 where ltsb.teacher_id=$1 and ltsb.lesson_id=$2`
	var lts models.LessonTeacherStudent
	err := t.studentDB.Select(&lts, query, teacherID, lessonID)
	return lts, err
}

func (t *TimetableRepository) AddTimetable(input models.Timetable) error {
	query := `insert into timetables(weekday,group_id,lesson_id, time_id, auditory_id,alt_lesson_id, alt_auditory_id,type) values ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := t.studentDB.Exec(query, input.Weekday, input.GroupID, input.LessonID, input.TimeID, input.AuditoryID, input.AltLessonID, input.AltAuditoryID, input.TypeID)
	return err
}

func (t *TimetableRepository) DeleteTimetable(id int) error {
	query := `delete from timetables where id=$1`
	_, err := t.studentDB.Exec(query, id)
	return err
}

func (t *TimetableRepository) GetTimetableOfGroup(groupID int) ([]models.Timetable, error) {
	query := `select weekday,
				group_id,
				group_name,
				lesson_id,
				lesson_name,
				time_id,
				start_time,
				end_time,
				auditory_id,
				auditory_name,
				alt_lesson_id,
				alt_lesson_name,
				alt_auditory_id,
				alt_auditory_name,
				type_id,
				type_name,
			from timetable_view where group_id = $1`

	var timetables []models.Timetable
	err := t.studentDB.Select(&timetables, query, groupID)
	return timetables, err
}
