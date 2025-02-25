package admin

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"student-management/internal/config"
	"student-management/internal/models"
)

type ProfessionRepository struct {
	studentDB *sqlx.DB
	config    *config.AppConfig
}

func NewProfessionRepository(studentDB *sqlx.DB, config *config.AppConfig) *ProfessionRepository {
	return &ProfessionRepository{
		studentDB: studentDB,
		config:    config,
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
	query := `insert into professions(name,code,department_id, position) values ($1,$2,$3,$4) returning id`
	var profession models.Profession
	err := p.studentDB.QueryRow(query, input.Name, input.Code, input.DepartmentID, input.Position).Scan(&profession.ID)
	if err != nil {
		return models.Profession{}, err
	}
	return profession, nil
}

func (p *ProfessionRepository) UpdateProfession(input models.Profession) error {
	query := `update professions set name=$1, code=$2, department_id=$3, position=$4 where id=$5`
	_, err := p.studentDB.Exec(query, input.Name, input.Code, input.DepartmentID, input.Position, input.ID)
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
	query := `select p.id,p.name,p.code,p.department_id, p.position, d.name as department_name from professions as p join departments as d on d.id=p.department_id where p.id=$1`
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
		setValues = append(setValues, fmt.Sprintf("p.name like'%%%s%%'", input.Name))
		//args = append(args, input.Name)
		//argId++
	}

	if input.Code != "" {
		setValues = append(setValues, fmt.Sprintf("p.code = %s", input.Code))
		//args = append(args, input.Code)
		argId++
	}

	if input.ID != 0 {
		setValues = append(setValues, fmt.Sprintf("p.id = %d", input.ID))
		//args = append(args, input.ID)
		argId++
	}

	if input.DepartmentID != 0 {
		setValues = append(setValues, fmt.Sprintf("p.department_id = %d", input.DepartmentID))
		//args = append(args, input.ID)
		argId++
	}

	if input.FacultyID != 0 {
		setValues = append(setValues, fmt.Sprintf("f.faculty_id = %d", input.FacultyID))
		argId++
	}

	queryArgs := strings.Join(setValues, " and ")

	var query string

	if argId > 1 || input.Name != "" {
		query = "select (select count(*) from groups where profession_id=p.id) as group_count, p.id,p.name,p.code,p.department_id, p.position, d.name as department_name from professions as p join departments as d on d.id=p.department_id join faculties as f on f.id=d.faculty_id where " + queryArgs
	} else {
		query = "select (select count(*) from groups where profession_id=p.id) as group_count, p.id,p.name,p.code,p.department_id, p.position, d.name as department_name from professions as p join departments as d on d.id=p.department_id join faculties as f on f.id=d.faculty_id"
	}

	paginationQuery := fmt.Sprintf(`select count(*) from (%s) as s`, query)
	pagination, offset, err := p.getPagination(paginationQuery, input.Limit, input.Page)
	if err != nil {
		return models.ProfessionAndPagination{}, err
	}

	query += fmt.Sprintf(` order by p.position limit %d offset %d`, input.Limit, offset)

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

func (p *ProfessionRepository) AddFile(id int, fileURL string, name string) error {
	query := `insert into profession_files (profession_id, file_url,name) values ($1, $2,$3)`
	_, err := p.studentDB.Exec(query, id, fileURL, name)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProfessionRepository) DeleteFile(id int) error {
	query := `delete from profession_files where id = $1`
	_, err := p.studentDB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProfessionRepository) GetFileByID(id int) (string, error) {
	var res string
	query := `select file_url from profession_files where id = $1`
	err := p.studentDB.Get(&res, query, id)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (p *ProfessionRepository) GetAllFiles(id int) ([]models.File, error) {
	var files []models.File
	query := fmt.Sprintf(`select id,'%s' || file_url as file_url,name from profession_files where profession_id = $1`, p.config.Domains.File)
	err := p.studentDB.Select(&files, query, id)
	if err != nil {
		return []models.File{}, err
	}

	return files, nil
}

func (p *ProfessionRepository) GetStudents(id int) (int, error) {
	var studentsCount int
	query := `
			select count(*)
			from students as s
					 join groups as g on g.id = s.group_id
					join professions as p on g.profession_id = p.id
			where p.id = $1
		`
	err := p.studentDB.Get(&studentsCount, query, id)
	if err != nil {
		return 0, err
	}
	return studentsCount, nil
}

func (p *ProfessionRepository) GetGroups(id int) ([]models.Group, error) {
	query := `
		select 
			   g.id,
			   g.name,
			   g.code,
			   g.year,
			   g.profession_id
		from groups as g
				 join professions as p on p.id = g.profession_id
		where p.id = $1
		`
	var groups []models.Group
	err := p.studentDB.Select(&groups, query, id)
	if err != nil {
		return []models.Group{}, err
	}
	return groups, nil
}
