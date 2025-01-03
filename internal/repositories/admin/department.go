package admin

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"student-management/internal/models"
)

type DepartmentRepository struct {
	studentDB *sqlx.DB
}

func NewDepartmentService(studentDB *sqlx.DB) *DepartmentRepository {
	return &DepartmentRepository{studentDB: studentDB}
}

func (d *DepartmentRepository) getPagination(query string, limit, page int) (models.Pagination, int, error) {
	var count, add int

	err := d.studentDB.QueryRow(query).Scan(&count)

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

func (d *DepartmentRepository) AddDepartment(department models.Department) (models.Department, error) {
	query := `insert into departments(name,code,faculty_id) values ($1,$2,$3) returning id`
	err := d.studentDB.QueryRow(query, department.Name, department.Code, department.FacultyID).Scan(&department.ID)
	if err != nil {
		return models.Department{}, err
	}
	return department, nil
}

func (d *DepartmentRepository) UpdateDepartment(department models.Department) error {
	query := `update departments set name=$1 , code=$2 , faculty_id=$3 where id=$4`
	_, err := d.studentDB.Exec(query, department.Name, department.Code, department.FacultyID, department.ID)
	if err != nil {
		return err
	}
	return nil
}

func (d *DepartmentRepository) GetDepartmentById(id int) (models.Department, error) {
	var department models.Department
	query := `select d.id,d.name,d.code,coalesce(d.faculty_id,0) as faculty_id,f.name as faculty_name from departments as d join faculties as f on f.id=d.faculty_id where d.id=$1`
	err := d.studentDB.Get(&department, query, id)
	if err != nil {
		return models.Department{}, err
	}
	return department, nil
}

func (d *DepartmentRepository) DeleteDepartment(id int) error {
	query := `delete from departments where id=$1`
	_, err := d.studentDB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (d *DepartmentRepository) GetDepartments(input models.DepartmentSearch) (models.DepartmentAndPagination, error) {
	setValues := make([]string, 0)
	//args := make([]interface{}, 0)
	argId := 1

	if input.Name != "" {
		setValues = append(setValues, fmt.Sprintf("d.name like'%%%s%%'", input.Name))
		//args = append(args, input.Name)
		argId++
	}

	if input.Code != "" {
		setValues = append(setValues, fmt.Sprintf("d.code = '%s'", input.Code))
		//args = append(args, input.Code)
		argId++
	}

	if input.ID != 0 {
		setValues = append(setValues, fmt.Sprintf("d.id = %d", input.ID))
		//args = append(args, input.ID)
		argId++
	}

	if input.FacultyID != 0 {
		setValues = append(setValues, fmt.Sprintf("d.faculty_id = %d", input.FacultyID))
		//args = append(args, input.FacultyID)
		argId++
	}

	queryArgs := strings.Join(setValues, " and ")

	var query string

	if argId > 1 || input.Name != "" {
		query = "select (select count(*) from teachers where department_id=d.id) as teachers_count,d.id,d.name,d.code,coalesce(d.faculty_id,0) as faculty_id,f.name as faculty_name from departments as d join faculties as f on f.id=d.faculty_id  where " + queryArgs
	} else {
		query = "select (select count(*) from teachers where department_id=d.id) as teachers_count,d.id,d.name,d.code,coalesce(d.faculty_id,0) as faculty_id,f.name as faculty_name from departments as d join faculties as f on f.id=d.faculty_id "
	}

	paginationQuery := fmt.Sprintf(`select count(*) from (%s) as s`, query)
	pagination, offset, err := d.getPagination(paginationQuery, input.Limit, input.Page)
	if err != nil {
		return models.DepartmentAndPagination{}, err
	}

	query += fmt.Sprintf(` limit %d offset %d`, input.Limit, offset)

	var departments []models.Department
	err = d.studentDB.Select(&departments, query)
	if err != nil {
		return models.DepartmentAndPagination{}, err
	}

	return models.DepartmentAndPagination{
		Departments: departments,
		Pagination:  pagination,
	}, nil
}
