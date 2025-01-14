package admin

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"student-management/internal/config"
	"student-management/internal/models"
)

type FacultyRepository struct {
	config    *config.AppConfig
	studentDB *sqlx.DB
}

func NewFacultyRepository(studentDB *sqlx.DB, config *config.AppConfig) *FacultyRepository {
	return &FacultyRepository{studentDB: studentDB, config: config}
}

func (f *FacultyRepository) getPagination(query string, limit, page int) (models.Pagination, int, error) {
	var count, add int

	err := f.studentDB.QueryRow(query).Scan(&count)

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

func (f *FacultyRepository) AddFaculty(input models.Faculty) (models.Faculty, error) {
	query := `insert into faculties (name,code,position) values ($1, $2,$3) returning id`
	err := f.studentDB.QueryRow(query, input.Name, input.Code, input.Position).Scan(&input.ID)
	if err != nil {
		return models.Faculty{}, err
	}
	return input, nil
}

func (f *FacultyRepository) UpdateFaculty(input models.Faculty) error {
	query := `update faculties set name = $1, code = $2,position=$3 where id = $4`
	_, err := f.studentDB.Exec(query, input.Name, input.Code, input.Position, input.ID)
	if err != nil {
		return err
	}
	return nil
}

func (f *FacultyRepository) GetFacultyByID(id int) (models.Faculty, error) {
	query := `select id,name,code,position from faculties where id = $1`
	var faculty models.Faculty
	err := f.studentDB.Get(&faculty, query, id)
	if err != nil {
		return models.Faculty{}, err
	}
	return faculty, nil
}

func (f *FacultyRepository) DeleteFaculty(id int) error {
	query := `delete from faculties where id = $1`
	_, err := f.studentDB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (f *FacultyRepository) GetFaculties(input models.FacultySearch) (models.FacultiesAndPagination, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != "" {
		setValues = append(setValues, fmt.Sprintf("f.name like'%%%s%%'", input.Name))
		//args = append(args, input.Name)
		//argId++
	}

	if input.Code != "" {
		setValues = append(setValues, fmt.Sprintf("f.code = $%d", argId))
		args = append(args, input.Code)
		argId++
	}

	if input.ID != 0 {
		setValues = append(setValues, fmt.Sprintf("f.id = $%d", argId))
		args = append(args, input.ID)
		argId++
	}

	queryArgs := strings.Join(setValues, " and ")

	var query string

	if argId > 1 || input.Name != "" {
		query = "select f.id,f.name,f.position,f.code,(select count(*) from departments where faculty_id=f.id) as department_count from faculties as f where " + queryArgs
	} else {
		query = "select f.id,f.name,f.position,f.code,(select count(*) from departments where faculty_id=f.id) as department_count from faculties as f "
	}

	paginationQuery := fmt.Sprintf(`select count(*) from (%s) as s`, query)
	pagination, offset, err := f.getPagination(paginationQuery, input.Limit, input.Page)
	if err != nil {
		return models.FacultiesAndPagination{}, err
	}

	query += fmt.Sprintf(`order by f.position  limit %d offset %d`, input.Limit, offset)

	var faculties []models.Faculty
	err = f.studentDB.Select(&faculties, query, args...)
	if err != nil {
		return models.FacultiesAndPagination{}, err
	}

	return models.FacultiesAndPagination{
		Faculties:  faculties,
		Pagination: pagination,
	}, nil

}

func (f *FacultyRepository) AddFile(id int, fileURL string, name string) error {
	query := `insert into faculty_files (faculty_id, file_url,name) values ($1, $2,$3)`
	_, err := f.studentDB.Exec(query, id, fileURL, name)
	if err != nil {
		return err
	}
	return nil
}

func (f *FacultyRepository) GetFileByID(id int) (string, error) {
	var res string
	query := `select file_url from faculty_files where id = $1`
	err := f.studentDB.Get(&res, query, id)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (f *FacultyRepository) DeleteFile(fileID int) error {
	query := `delete from faculty_files where id = $1`
	_, err := f.studentDB.Exec(query, fileID)
	if err != nil {
		return err
	}
	return nil
}

func (f *FacultyRepository) GetAllFiles(id int) ([]models.File, error) {
	var files []models.File
	query := fmt.Sprintf(`select id, '%s' || file_url as file_url,name from faculty_files where faculty_id = $1`, f.config.Domains.File)
	err := f.studentDB.Select(&files, query, id)
	if err != nil {
		return []models.File{}, err
	}

	return files, nil
}

func (f *FacultyRepository) GetProfessions(id int) ([]models.Profession, error) {
	query := `select  p.id,p.name,p.code,p.department_id, d.name as department_name from professions as p join departments as d on d.id=p.department_id join faculties as f on f.id=d.faculty_id where f.id = $1`
	var professions []models.Profession
	err := f.studentDB.Select(&professions, query, id)
	if err != nil {
		return []models.Profession{}, err
	}
	return professions, nil
}

func (f *FacultyRepository) GetDepartments(id int) ([]models.Department, error) {
	query := `select d.id,d.name,d.code,coalesce(d.faculty_id,0) as faculty_id,f.name as faculty_name from departments as d join faculties as f on f.id=d.faculty_id where f.id = $1`
	var departments []models.Department
	err := f.studentDB.Select(&departments, query, id)
	if err != nil {
		return []models.Department{}, err
	}
	return departments, nil
}

func (f *FacultyRepository) GetTeachers(id int) ([]models.Teacher, error) {
	query := `select t.id,t.first_name,t.middle_name,t.last_name,t.code,t.department_id,coalesce(t.image,'') as image,d.name as department_name from teachers as t join departments as d on d.id=t.department_id join faculties as f on d.faculty_id = f.id where f.id = $1`
	var teachers []models.Teacher
	err := f.studentDB.Select(&teachers, query, id)
	if err != nil {
		return []models.Teacher{}, err
	}
	return teachers, nil
}

func (f *FacultyRepository) GetStudentsCount(id int) (int, error) {
	var count int
	query := `
			select count(*)
			from students as s
					 join groups g on s.group_id = g.id
					 join professions as p on g.profession_id = p.id
					 join departments as d on p.department_id = d.id
					 join faculties as f on d.faculty_id = f.id
			where f.id = $1
			`
	err := f.studentDB.Get(&count, query, id)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (f *FacultyRepository) GetGroupCount(id int) (int, error) {
	var count int
	query := `
		select count(*)
		from groups as g
				 join professions as p on g.profession_id = p.id
				 join departments as d on p.department_id = d.id
				 join faculties as f on d.faculty_id = f.id
		where f.id = $1
		`
	err := f.studentDB.Get(&count, query, id)
	if err != nil {
		return 0, err
	}
	return count, nil
}
