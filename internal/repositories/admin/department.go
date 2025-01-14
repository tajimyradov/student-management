package admin

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"student-management/internal/config"
	"student-management/internal/models"
)

type DepartmentRepository struct {
	config    *config.AppConfig
	studentDB *sqlx.DB
}

func NewDepartmentService(studentDB *sqlx.DB, config *config.AppConfig) *DepartmentRepository {
	return &DepartmentRepository{studentDB: studentDB, config: config}
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

func (d *DepartmentRepository) AddFile(id int, fileURL string, name string) error {
	query := `insert into department_files (department_id, file_url,name) values ($1, $2,$3)`
	_, err := d.studentDB.Exec(query, id, fileURL, name)
	if err != nil {
		return err
	}
	return nil
}

func (d *DepartmentRepository) DeleteFile(id int) error {
	query := `delete from department_files where id = $1`
	_, err := d.studentDB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (d *DepartmentRepository) GetFileByID(id int) (string, error) {
	var res string
	query := `select file_url from department_files where id = $1`
	err := d.studentDB.Get(&res, query, id)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (d *DepartmentRepository) GetAllFiles(id int) ([]models.File, error) {
	var files []models.File
	query := fmt.Sprintf(`select id,'%s' || file_url as file_url,name from department_files where department_id = $1`, d.config.Domains.File)
	err := d.studentDB.Select(&files, query, id)
	if err != nil {
		return []models.File{}, err
	}

	return files, nil
}

func (d *DepartmentRepository) GetProfessions(id int) ([]models.Profession, error) {
	query := `select  p.id,p.name,p.code,p.department_id, d.name as department_name from professions as p join departments as d on d.id=p.department_id  where d.id = $1`
	var professions []models.Profession
	err := d.studentDB.Select(&professions, query, id)
	if err != nil {
		return []models.Profession{}, err
	}
	return professions, nil
}

func (d *DepartmentRepository) GetTeachers(id int) ([]models.Teacher, error) {
	query := `select t.id,t.first_name,t.middle_name,t.last_name,t.code,t.department_id,coalesce(t.image,'') as image,d.name as department_name from teachers as t join departments as d on d.id=t.department_id where d.id = $1`
	var teachers []models.Teacher
	err := d.studentDB.Select(&teachers, query, id)
	if err != nil {
		return []models.Teacher{}, err
	}
	return teachers, nil
}

func (d *DepartmentRepository) GetStudentsCount(id int) (int, error) {
	var count int
	query := `
			select count(*)
			from students as s
					 join groups g on s.group_id = g.id
					 join professions as p on g.profession_id = p.id
					 join departments as d on p.department_id = d.id
			where d.id = $1
			`
	err := d.studentDB.Get(&count, query, id)
	if err != nil {
		return 0, err
	}
	return count, nil

}

func (d *DepartmentRepository) GetGroupCount(id int) (int, error) {
	var count int
	query := `
		select count(*)
		from groups as g
				 join professions as p on g.profession_id = p.id
				 join departments as d on p.department_id = d.id
		where d.id = $1
		`
	err := d.studentDB.Get(&count, query, id)
	if err != nil {
		return 0, err
	}
	return count, nil
}
