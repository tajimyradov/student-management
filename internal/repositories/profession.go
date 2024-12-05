package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"student-management/internal/models"
)

type ProfessionRepository struct {
	studentDB *sqlx.DB
}

func NewProfessionRepository(studentDB *sqlx.DB) *ProfessionRepository {
	return &ProfessionRepository{
		studentDB: studentDB,
	}
}

func (p *ProfessionRepository) getPagination(query string, limit, page int) (models.Pagination, int, error) {
	var count, add int

	err := p.studentDB.QueryRow(query).Scan(&count)

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

func (p *ProfessionRepository) AddProfession(input models.Profession) (models.Profession, error) {
	query := `insert into professions(name,code,department_id) values ($1,$2,$3) returning id`
	var profession models.Profession
	err := p.studentDB.QueryRow(query, input.Name, input.Code, input.DepartmentID).Scan(&profession.ID)
	if err != nil {
		return models.Profession{}, err
	}
	return profession, nil
}

func (p *ProfessionRepository) UpdateProfession(input models.Profession) error {
	query := `update professions set name=$1, code=$2, department_id=$3 where id=$4`
	_, err := p.studentDB.Exec(query, input.Name, input.Code, input.DepartmentID, input.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProfessionRepository) DeleteProfession(id int) error {
	query := `delete from professions where id=$1`
	_, err := p.studentDB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProfessionRepository) GetProfessionByID(id int) (models.Profession, error) {
	query := `select id,name,code,department_id from professions where id=$1`
	var profession models.Profession
	err := p.studentDB.Get(&profession, query, id)
	if err != nil {
		return models.Profession{}, err
	}
	return profession, nil
}

func (p *ProfessionRepository) GetProfessions(input models.ProfessionSearch) (models.ProfessionAndPagination, error) {

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

	if input.DepartmentID != 0 {
		setValues = append(setValues, fmt.Sprintf("department_id = $%d", argId))
		args = append(args, input.ID)
		argId++
	}

	queryArgs := strings.Join(setValues, " and ")

	var query string

	if argId > 1 || input.Name != "" {
		query = "select id,name,code,faculty_id from departments where " + queryArgs
	} else {
		query = "select id,name,code,faculty_id from departments"
	}

	paginationQuery := fmt.Sprintf(`select count(*) from (%s) as s`, query)
	pagination, offset, err := p.getPagination(paginationQuery, input.Limit, input.Page)
	if err != nil {
		return models.ProfessionAndPagination{}, err
	}

	query += fmt.Sprintf(` limit %d offset %d`, input.Limit, offset)

	var professions []models.Profession
	err = p.studentDB.Select(&professions, query, args...)
	if err != nil {
		return models.ProfessionAndPagination{}, err
	}

	return models.ProfessionAndPagination{
		Professions: professions,
		Pagination:  pagination,
	}, nil

}
