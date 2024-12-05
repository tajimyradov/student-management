package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"student-management/internal/models"
)

type FacultyRepository struct {
	studentDB *sqlx.DB
}

func NewFacultyRepository(studentDB *sqlx.DB) *FacultyRepository {
	return &FacultyRepository{studentDB: studentDB}
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
	query := `insert into faculties (name,code) values ($1, $2) returning id`
	err := f.studentDB.QueryRow(query, input.Name, input.Code).Scan(&input.ID)
	if err != nil {
		return models.Faculty{}, err
	}
	return input, nil
}

func (f *FacultyRepository) UpdateFaculty(input models.Faculty) error {
	query := `update faculties set name = $1, code = $2 where id = $3`
	_, err := f.studentDB.Exec(query, input.Name, input.Code, input.ID)
	if err != nil {
		return err
	}
	return nil
}

func (f *FacultyRepository) GetFacultyByID(id int) (models.Faculty, error) {
	query := `select id,name,code from faculties where id = $1`
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
		setValues = append(setValues, fmt.Sprintf("name like'%%%s%%'", input.Name))
		//args = append(args, input.Name)
		//argId++
	}

	if input.Code != "" {
		setValues = append(setValues, fmt.Sprintf("code = $%d", argId))
		args = append(args, input.Code)
		argId++
	}

	if input.ID != 0 {
		setValues = append(setValues, fmt.Sprintf("id = $%d", argId))
		args = append(args, input.ID)
		argId++
	}

	queryArgs := strings.Join(setValues, " and ")

	var query string

	if argId > 1 || input.Name != "" {
		query = "select id,name,code from faculties where " + queryArgs
	} else {
		query = "select id,name,code from faculties"
	}

	paginationQuery := fmt.Sprintf(`select count(*) from (%s) as s`, query)
	pagination, offset, err := f.getPagination(paginationQuery, input.Limit, input.Page)
	if err != nil {
		return models.FacultiesAndPagination{}, err
	}

	query += fmt.Sprintf(` limit %d offset %d`, input.Limit, offset)

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
