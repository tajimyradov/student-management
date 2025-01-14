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
	query := `insert into students(first_name, last_name, code, gender, username, password, group_id,birth_date,region_id,middle_name) values ($1, $2, $3, $4, $5, $6, $7,$8,$9,$10) returning id`
	err := s.studentDB.QueryRow(query,
		student.FirstName,
		student.LastName,
		student.Code,
		student.Gender,
		student.Username,
		student.Password,
		student.GroupID,
		student.BirthDate,
		student.RegionID,
		student.MiddleName,
	).Scan(&student.ID)
	return student, err
}

func (s *StudentRepository) UpdateStudent(student models.Student) error {
	query := `update students set first_name=$1, last_name=$2, code=$3, gender=$4, username=$5, group_id=$6,birth_date=$7,region_id=$8,middle_name=$9 where id=$10`
	_, err := s.studentDB.Exec(query,
		student.FirstName,
		student.LastName,
		student.Code,
		student.Gender,
		student.Username,
		student.GroupID,
		student.BirthDate,
		student.RegionID,
		student.MiddleName,
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
				   s.middle_name,
				   s.code,
				   s.gender,
				   coalesce(s.username, '') as username,
				   coalesce(s.password, '') as password,
				   s.group_id,
				   s.birth_date,
				   coalesce(s.image, '')    as image,
				   f.id as faculty_id,
				   f.name as faculty_name,
				   d.id as department_id,
				   d.name as department_name,
				   g.name                   as group_name,
				   coalesce(r.id,0) as region_id,
				   coalesce(r.name,'') as region_name
			from students as s
					 join groups as g on g.id = s.group_id
					join professions as p on g.profession_id = p.id
			join departments as d on p.department_id = d.id
			join faculties as f on d.faculty_id = f.id 
			left join regions as r on s.region_id = r.id
			where s.id=$1`
	err := s.studentDB.Get(&student, query, id)
	return student, err
}

func (s *StudentRepository) UpdateStudentsImage(image string, id int) error {
	query := `update students set image=$1 where id=$2`
	_, err := s.studentDB.Exec(query, image, id)
	return err
}

func (s *StudentRepository) GetStudents(input models.StudentSearch) (models.StudentsAndPagination, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Name != "" {
		input.Name = strings.ToLower(input.Name)
		setValues = append(setValues, fmt.Sprintf("lower(s.first_name) like'%%%s%%' or lower(s.last_name) like'%%%s%%'  or lower(s.middle_name) like'%%%s%%' ", input.Name, input.Name, input.Name))
		//args = append(args, input.Name)
		//argId++
	}

	if input.Code != "" {
		setValues = append(setValues, fmt.Sprintf("s.code = %s", input.Code))
	}

	if input.ID != 0 {
		setValues = append(setValues, fmt.Sprintf("s.id = %d", input.ID))
		//args = append(args, input.ID)
		//argId++
	}

	if input.Username != "" {
		setValues = append(setValues, fmt.Sprintf("s.username like '%%%s%%'", input.Username))
	}

	if input.BirthDate != "" {
		setValues = append(setValues, fmt.Sprintf("s.birth_date = %s", input.BirthDate))
		//args = append(args, input.BirthDate)
		//argId++
	}

	if input.FacultyID != 0 {
		setValues = append(setValues, fmt.Sprintf("f.id = %d", input.FacultyID))
		argId++
	}

	if input.GroupID != 0 {
		setValues = append(setValues, fmt.Sprintf("g.id = %d", input.GroupID))
		argId++
	}

	if input.ProfessionID != 0 {
		setValues = append(setValues, fmt.Sprintf("p.id = %d", input.ProfessionID))
		argId++
	}

	queryArgs := strings.Join(setValues, " and ")

	var query string

	if argId > 1 || input.Name != "" || input.Username != "" {
		query = `
			select s.id,
				   s.first_name,
				   s.last_name,
				   s.middle_name,
				   s.code,
				   s.gender,
				   coalesce(s.username, '') as username,
				   coalesce(s.password, '') as password,
				   s.group_id,
				   s.birth_date,
				   coalesce(s.image, '')    as image,
				   f.id as faculty_id,
				   f.name as faculty_name,
				   d.id as department_id,
				   d.name as department_name,
				   g.name                   as group_name,
				   coalesce(r.id,0) as region_id,
				   coalesce(r.name,'') as region_name
			from students as s
					 join groups as g on g.id = s.group_id
					join professions as p on g.profession_id = p.id
			join departments as d on p.department_id = d.id
			join faculties as f on d.faculty_id = f.id
			left join regions as r on s.region_id = r.id
			where ` + queryArgs
	} else {
		query = `
			select s.id,
				   s.first_name,
				   s.last_name,
				   s.middle_name,
				   s.code,
				   s.gender,
				   coalesce(s.username, '') as username,
				   coalesce(s.password, '') as password,
				   s.group_id,
				   s.birth_date,
				   coalesce(s.image, '')    as image,
				   f.id as faculty_id,
				   f.name as faculty_name,
				   d.id as department_id,
				   d.name as department_name,
				   g.name                   as group_name,
				   coalesce(r.id,0) as region_id,
				   coalesce(r.name,'') as region_name
			from students as s
					 join groups as g on g.id = s.group_id
					join professions as p on g.profession_id = p.id
			join departments as d on p.department_id = d.id
			join faculties as f on d.faculty_id = f.id
			left join regions as r on s.region_id = r.id
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
