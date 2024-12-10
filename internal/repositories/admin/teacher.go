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
	query := `insert into teachers(first_name, last_name, code, gender, username, password, department_id) values ($1, $2, $3, $4, $5, $6, $7) returning id`
	err := t.studentDB.QueryRow(query,
		teacher.FirstName,
		teacher.LastName,
		teacher.Code,
		teacher.Gender,
		teacher.Username,
		teacher.Password,
		teacher.DepartmentId,
	).Scan(&teacher.ID)
	return teacher, err
}

func (t *TeacherRepository) UpdateTeacher(teacher models.Teacher) error {
	query := `update teachers set first_name=$1, last_name=$2, code=$3, gender=$4, username=$5, department_id=$6 where id=$7`
	_, err := t.studentDB.Exec(query,
		teacher.FirstName,
		teacher.LastName,
		teacher.Code,
		teacher.Gender,
		teacher.Username,
		teacher.DepartmentId,
		teacher.ID,
	)
	return err
}

func (t *TeacherRepository) DeleteTeacher(id int) error {
	query := `delete from teachers where id = $1`
	_, err := t.studentDB.Exec(query, id)
	return err
}

func (t *TeacherRepository) GetTeacherByID(ID int) (models.Teacher, error) {
	var teacher models.Teacher
	query := `select t.id,t.first_name,t.last_name,t.code,t.gender,t.username,t.password,t.department_id,t.image,d.name as department_name from teachers as t join departments as d on d.id=t.department_id where id = $1`
	err := t.studentDB.Get(&teacher, query, ID)
	return teacher, err
}

func (t *TeacherRepository) GetTeachers(input models.TeacherSearch) (models.TeachersAndPagination, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.FirstName != "" {
		setValues = append(setValues, fmt.Sprintf("t.first_name like'%%%s%%'", input.FirstName))
		//args = append(args, input.Name)
		//argId++
	}

	if input.LastName != "" {
		setValues = append(setValues, fmt.Sprintf("t.last_name like'%%%s%%'", input.LastName))
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

	if argId > 1 || input.FirstName != "" || input.LastName != "" || input.Username != "" {
		query = "select t.id,t.first_name,t.last_name,t.code,t.gender,t.username,t.password,t.department_id,t.image,d.name as department_name from teachers as t join departments as d on d.id=t.department_id where " + queryArgs
	} else {
		query = "select t.id,t.first_name,t.last_name,t.code,t.gender,t.username,t.password,t.department_id,t.image,d.name as department_name from teachers as t join departments as d on d.id=t.department_id"
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
