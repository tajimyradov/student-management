package admin

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"student-management/internal/models"
)

type TeacherRepository struct {
	studentDB *sqlx.DB
}

func NewTeacherRepository(studentDB *sqlx.DB) *TeacherRepository {
	return &TeacherRepository{
		studentDB: studentDB,
	}
}

func (t *TeacherRepository) getPagination(query string, limit, page int) (models.Pagination, int, error) {
	var count, add int

	err := t.studentDB.QueryRow(query).Scan(&count)

	if err != nil {
		return models.Pagination{}, 0, err
	}

	if count%limit > 0 {
		add = 1
	}

	return models.Pagination{
		TotalPages:  count/limit + add,
		TotalCount:  count,
		Limit:       limit,
		CurrentPage: page,
	}, (page - 1) * limit, nil

}

func (t *TeacherRepository) AddTeacher(teacher models.Teacher) (models.Teacher, error) {
	query := `insert into teachers(first_name, last_name, code, gender, username, password, department_id,middle_name) values ($1, $2, $3, $4, $5, $6, $7,$8) returning id`
	err := t.studentDB.QueryRow(query,
		teacher.FirstName,
		teacher.LastName,
		teacher.Code,
		teacher.Gender,
		teacher.Username,
		teacher.Password,
		teacher.DepartmentId,
		teacher.MiddleName,
	).Scan(&teacher.ID)
	return teacher, err
}

func (t *TeacherRepository) UpdateTeacher(teacher models.Teacher) error {
	query := `update teachers set first_name=$1, last_name=$2, code=$3, gender=$4, username=$5, department_id=$6,middle_name=$7 where id=$8`
	_, err := t.studentDB.Exec(query,
		teacher.FirstName,
		teacher.LastName,
		teacher.Code,
		teacher.Gender,
		teacher.Username,
		teacher.DepartmentId,
		teacher.MiddleName,
		teacher.ID,
	)
	return err
}

func (t *TeacherRepository) DeleteTeacher(id int) error {
	_, err := t.studentDB.Exec(`delete from deans where teacher_id=$1`, id)
	if err != nil {
		return err
	}

	_, err = t.studentDB.Exec(`delete from department_leads where teacher_id=$1`, id)
	if err != nil {
		return err
	}

	query := `delete from teachers where id = $1`
	_, err = t.studentDB.Exec(query, id)
	return err
}

func (t *TeacherRepository) GetTeacherByID(ID int) (models.Teacher, error) {
	var teacher models.Teacher
	query := `select t.id,t.first_name,t.middle_name,t.last_name,t.code,t.gender,coalesce(t.username,'') as username,coalesce(t.password,'') as password,t.department_id,coalesce(t.image,'') as image,d.name as department_name from teachers as t join departments as d on d.id=t.department_id where t.id = $1`
	err := t.studentDB.Get(&teacher, query, ID)
	return teacher, err
}

func (t *TeacherRepository) GetTeachers(input models.TeacherSearch) (models.TeachersAndPagination, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != "" {
		input.Name = strings.ToLower(input.Name)
		setValues = append(setValues, fmt.Sprintf("lower(t.first_name) like'%%%s%%' or lower(t.last_name) like'%%%s%%' or lower(t.middle_name) like'%%%s%%'", input.Name, input.Name, input.Name))
		//args = append(args, input.Name)
		//argId++
	}

	if input.Code != "" {
		setValues = append(setValues, fmt.Sprintf("t.code = $%d", argId))
		args = append(args, input.Code)
		argId++
	}

	if input.ID != 0 {
		setValues = append(setValues, fmt.Sprintf("t.id = $%d", argId))
		args = append(args, input.ID)
		argId++
	}

	if input.Username != "" {
		setValues = append(setValues, fmt.Sprintf("t.username like '%%%s%%'", input.Username))
	}

	if input.DepartmentId != 0 {
		setValues = append(setValues, fmt.Sprintf("t.department_id = $%d", argId))
		args = append(args, input.DepartmentId)
		argId++
	}

	queryArgs := strings.Join(setValues, " and ")

	var query string

	if argId > 1 || input.Name != "" || input.Username != "" {
		query = "select t.id,t.first_name,t.middle_name,t.last_name,t.code,t.gender,coalesce(t.username,'') as username,coalesce(t.password,'') as password,t.department_id,coalesce(t.image,'') as image,d.name as department_name, coalesce(g.id,0) as group_id, coalesce(g.name,'') as group_name from teachers as t join departments as d on d.id=t.department_id left join groups as g on g.teacher_id=t.id where " + queryArgs
	} else {
		query = "select t.id,t.first_name,t.middle_name,t.last_name,t.code,t.gender,coalesce(t.username,'') as username,coalesce(t.password,'') as password,t.department_id,coalesce(t.image,'') as image,d.name as department_name, coalesce(g.id,0) as group_id, coalesce(g.name,'') as group_name from teachers as t join departments as d on d.id=t.department_id left join groups as g on g.teacher_id=t.id "
	}

	paginationQuery := fmt.Sprintf(`select count(*) from (%s) as s`, query)
	pagination, offset, err := t.getPagination(paginationQuery, input.Limit, input.Page)
	if err != nil {
		return models.TeachersAndPagination{}, err
	}

	query += fmt.Sprintf(` limit %d offset %d`, input.Limit, offset)

	var teachers []models.Teacher
	err = t.studentDB.Select(&teachers, query, args...)
	if err != nil {
		return models.TeachersAndPagination{}, err
	}

	return models.TeachersAndPagination{
		Teachers:   teachers,
		Pagination: pagination,
	}, nil
}

func (t *TeacherRepository) UpdateTeachersImage(image string, id int) error {
	query := `update teachers set image=$1 where id=$2`
	_, err := t.studentDB.Exec(query, image, id)
	return err
}
