package admin

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"student-management/internal/models"
)

type StudentRepository struct {
	studentDB *sqlx.DB
}

func NewStudentRepository(studentDB *sqlx.DB) *StudentRepository {
	return &StudentRepository{
		studentDB: studentDB,
	}
}

func (s *StudentRepository) getPagination(query string, limit, page int) (models.Pagination, int, error) {
	var count, add int

	err := s.studentDB.QueryRow(query).Scan(&count)

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

func (s *StudentRepository) AddStudent(student models.Student) (models.Student, error) {
	query := `insert into students(first_name, last_name, code, gender, username, password, group_id,birth_date) values ($1, $2, $3, $4, $5, $6, $7,$8) returning id`
	err := s.studentDB.QueryRow(query,
		student.FirstName,
		student.LastName,
		student.Code,
		student.Gender,
		student.Username,
		student.Password,
		student.GroupID,
		student.BirthDate,
	).Scan(&student.ID)
	return student, err
}

func (s *StudentRepository) UpdateStudent(student models.Student) error {
	query := `update students set first_name=$1, last_name=$2, code=$3, gender=$4, username=$5, group_id=$6,birth_date=$7 where id=$8`
	_, err := s.studentDB.Exec(query,
		student.FirstName,
		student.LastName,
		student.Code,
		student.Gender,
		student.Username,
		student.GroupID,
		student.BirthDate,
		student.ID,
	)
	return err
}

func (s *StudentRepository) DeleteStudent(id int) error {
	query := `delete from students where id=$1`
	_, err := s.studentDB.Exec(query, id)
	return err
}

func (s *StudentRepository) GetStudentByID(id int) (models.Student, error) {
	var student models.Student
	query := `select s.id,
				   s.first_name,
				   s.last_name,
				   s.code,
				   s.gender,
				   coalesce(s.username, '') as username,
				   coalesce(s.password, '') as password,
				   s.group_id,
				   s.birth_date,
				   coalesce(s.image, '')    as image,
				   f.id,
				   f.name,
				   d.id,
				   d.name,
				   g.name                   as group_name
			from students as s
					 join groups as g on g.id = s.group_id
					join professions as p on g.profession_id = p.id
			join departments as d on p.department_id = d.id
			join faculties as f on d.faculty_id = f.id where s.id=$1`
	err := s.studentDB.Get(&student, query, id)
	return student, err
}

func (s *StudentRepository) UpdateTeachersImage(image string, id int) error {
	query := `update students set image=$1 where id=$2`
	_, err := s.studentDB.Exec(query, image, id)
	return err
}

func (s *StudentRepository) GetStudents(input models.StudentSearch) (models.StudentsAndPagination, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.FirstName != "" {
		setValues = append(setValues, fmt.Sprintf("s.first_name like'%%%s%%'", input.FirstName))
		//args = append(args, input.Name)
		//argId++
	}

	if input.LastName != "" {
		setValues = append(setValues, fmt.Sprintf("s.last_name like'%%%s%%'", input.LastName))
		//args = append(args, input.Name)
		//argId++
	}

	if input.Code != "" {
		setValues = append(setValues, fmt.Sprintf("s.code = $%d", argId))
		args = append(args, input.Code)
		argId++
	}

	if input.ID != 0 {
		setValues = append(setValues, fmt.Sprintf("s.id = $%d", argId))
		args = append(args, input.ID)
		argId++
	}

	if input.Username != "" {
		setValues = append(setValues, fmt.Sprintf("s.username like '%%%s%%'", input.Username))
	}

	if input.BirthDate != "" {
		setValues = append(setValues, fmt.Sprintf("s.birth_date = $%d", argId))
		args = append(args, input.BirthDate)
		argId++
	}

	queryArgs := strings.Join(setValues, " and ")

	var query string

	if argId > 1 || input.FirstName != "" || input.LastName != "" || input.Username != "" {
		query = `
			select s.id,
				   s.first_name,
				   s.last_name,
				   s.code,
				   s.gender,
				   coalesce(s.username, '') as username,
				   coalesce(s.password, '') as password,
				   s.group_id,
				   s.birth_date,
				   coalesce(s.image, '')    as image,
				   f.id,
				   f.name,
				   d.id,
				   d.name,
				   g.name                   as group_name
			from students as s
					 join groups as g on g.id = s.group_id
					join professions as p on g.profession_id = p.id
			join departments as d on p.department_id = d.id
			join faculties as f on d.faculty_id = f.id 
			where ` + queryArgs
	} else {
		query = `
			select s.id,
				   s.first_name,
				   s.last_name,
				   s.code,
				   s.gender,
				   coalesce(s.username, '') as username,
				   coalesce(s.password, '') as password,
				   s.group_id,
				   s.birth_date,
				   coalesce(s.image, '')    as image,
				   f.id,
				   f.name,
				   d.id,
				   d.name,
				   g.name                   as group_name
			from students as s
					 join groups as g on g.id = s.group_id
					join professions as p on g.profession_id = p.id
			join departments as d on p.department_id = d.id
			join faculties as f on d.faculty_id = f.id
		`
	}

	paginationQuery := fmt.Sprintf(`select count(*) from (%s) as s`, query)
	pagination, offset, err := s.getPagination(paginationQuery, input.Limit, input.Page)
	if err != nil {
		return models.StudentsAndPagination{}, err
	}

	query += fmt.Sprintf(` limit %d offset %d`, input.Limit, offset)

	var students []models.Student
	err = s.studentDB.Select(&students, query, args...)
	if err != nil {
		return models.StudentsAndPagination{}, err
	}

	return models.StudentsAndPagination{
		Students:   students,
		Pagination: pagination,
	}, nil
}
