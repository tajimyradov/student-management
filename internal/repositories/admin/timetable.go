package admin

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
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

func (t *TimetableRepository) GetAbsences(input models.AbsenceSearch) ([]models.Absence, error) {

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	query := `
		select a.profession_id, a.profession_name,a.faculty_id,a.faculty_name, a.department_id, a.department_name,a.id,a.group_id, a.group_name,a.lesson_id, a.lesson_name, a.time_id,a.start_time, a.end_time,a.teacher_id, a.teacher_first_name, a.teacher_last_name,a.student_id, a.student_first_name, a.student_last_name,a.type_id, a.type_name, a.date,a.note, a.status
		from absences_view as a where %s
		`

	if input.FacultyID != 0 {
		setValues = append(setValues, fmt.Sprintf("a.faculty_id = $%d", argId))
		args = append(args, input.FacultyID)
		argId++
	}

	if input.DepartmentID != 0 {
		setValues = append(setValues, fmt.Sprintf("a.department_id = $%d", argId))
		args = append(args, input.DepartmentID)
		argId++
	}

	if input.GroupID != 0 {
		setValues = append(setValues, fmt.Sprintf("a.group_id = $%d", argId))
		args = append(args, input.GroupID)
		argId++
	}

	if input.TypeID != 0 {
		setValues = append(setValues, fmt.Sprintf("a.type_id = $%d", argId))
		args = append(args, input.TypeID)
		argId++
	}

	if input.TeacherID != 0 {
		setValues = append(setValues, fmt.Sprintf("a.teacher_id = $%d", argId))
		args = append(args, input.TeacherID)
		argId++
	}

	if input.StudentID != 0 {
		setValues = append(setValues, fmt.Sprintf("a.student_id = $%d", argId))
		args = append(args, input.StudentID)
		argId++
	}

	if input.StudentFirstName != "" {
		setValues = append(setValues, fmt.Sprintf("a.student_first_name like'%%%s%%'", input.StudentFirstName))
	}

	if input.StudentLastName != "" {
		setValues = append(setValues, fmt.Sprintf("a.student_lasy_name like'%%%s%%'", input.StudentLastName))
	}

	setValues = append(setValues, fmt.Sprintf(`(date  between $%d and $%d) `, argId, argId+1))
	args = append(args, input.From, input.To)
	queryArgs := strings.Join(setValues, " and ")
	query = fmt.Sprintf(query, queryArgs)

	var res []models.Absence
	err := t.studentDB.Select(&res, query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (t *TimetableRepository) UpdateAbsence(status, id int) error {
	query := `update absences set status=$1 where id=$2`
	_, err := t.studentDB.Exec(query, status, id)
	return err
}

func (t *TimetableRepository) Sync() error {
	query := `refresh materialized view  absences_view`
	_, err := t.studentDB.Exec(query)
	return err
}

func (t *TimetableRepository) GetAbsenceByID(id int) (models.Absence, error) {
	query := `
select a.profession_id, a.profession_name,a.student_year,a.faculty_dean_first_name,a.faculty_dean_last_name,a.department_lead_first_name,a.department_lead_last_name, a.faculty_id,a.faculty_name, a.department_id, a.department_name,a.id,a.group_id, a.group_name,a.lesson_id, a.lesson_name, a.time_id,a.start_time, a.end_time,a.teacher_id, a.teacher_first_name, a.teacher_last_name,a.student_id, a.student_first_name, a.student_last_name,a.type_id, a.type_name, a.date,a.note, a.status
		from absences_view as a where a.id=$1
	`
	var res models.Absence
	err := t.studentDB.Get(&res, query, id)
	if err != nil {
		return models.Absence{}, err
	}
	return res, nil
}
