package admin

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"student-management/internal/models"
)

type GroupRepository struct {
	studentDB *sqlx.DB
}

func NewGroupRepository(studentDB *sqlx.DB) *GroupRepository {
	return &GroupRepository{
		studentDB: studentDB,
	}
}

func (g *GroupRepository) getPagination(query string, limit, page int) (models.Pagination, int, error) {
	var count, add int

	err := g.studentDB.QueryRow(query).Scan(&count)

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

func (g *GroupRepository) AddGroup(input models.Group) (models.Group, error) {
	query := `insert into groups(name,code,year,profession_id,teacher_id) values ($1,$2,$3,$4,$5) returning id`
	err := g.studentDB.QueryRow(query, input.Name, input.Code, input.Year, input.ProfessionID, input.TeacherID).Scan(&input.ID)
	if err != nil {
		return models.Group{}, err
	}
	return input, nil
}

func (g *GroupRepository) UpdateGroup(input models.Group) error {
	query := `update groups set name=$1, code=$2, year=$3, profession_id=$4,teacher_id=$5 where id=$6`
	_, err := g.studentDB.Exec(query, input.Name, input.Code, input.Year, input.ProfessionID, input.TeacherID, input.ID)
	if err != nil {
		return err
	}
	return nil
}

func (g *GroupRepository) DeleteGroup(id int) error {
	query := `delete from groups where id=$1`
	_, err := g.studentDB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (g *GroupRepository) GetGroupByID(id int) (models.Group, error) {

	var group models.Group
	query := `select g.id,g.name,g.code,g.year,g.profession_id,coalesce(g.teacher_id,0) as teacher_id,coalesce(t.first_name,'') as teacher_first_name, coalesce(t.last_name,'') as teacher_last_name, p.name as profession_name from groups as g join professions as p on p.id =g.profession_id left join teachers as t on t.id=g.teacher_id where g.id=$1`
	err := g.studentDB.Get(&group, query, id)
	if err != nil {
		return models.Group{}, err
	}
	return group, nil
}

func (g *GroupRepository) GetGroups(input models.GroupSearch) (models.GroupsAndPagination, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != "" {
		setValues = append(setValues, fmt.Sprintf("g.name like'%%%s%%'", input.Name))
		//args = append(args, input.Name)
		//argId++
	}

	if input.Code != "" {
		setValues = append(setValues, fmt.Sprintf("g.code = %s", input.Code))
		//args = append(args, input.Code)
		argId++
	}

	if input.ID != 0 {
		setValues = append(setValues, fmt.Sprintf("g.id = %d", input.ID))
		//args = append(args, input.ID)
		argId++
	}

	if input.ProfessionID != 0 {
		setValues = append(setValues, fmt.Sprintf("g.profession_id = %d", input.ProfessionID))
		//args = append(args, input.ID)
		argId++
	}

	if input.Year != 0 {
		setValues = append(setValues, fmt.Sprintf("g.year = %d", input.Year))
		//args = append(args, input.Year)
		argId++
	}

	queryArgs := strings.Join(setValues, " and ")

	var query string

	if argId > 1 || input.Name != "" {
		query = "select (select count(*) from students where group_id=g.id) as student_count,g.id,g.name,g.code,g.year,g.profession_id,coalesce(g.teacher_id,0) as teacher_id,coalesce(t.first_name,'') as teacher_first_name, coalesce(t.last_name,'') as teacher_last_name, p.name as profession_name from groups as g join professions as p on p.id =g.profession_id left join teachers as t on t.id=g.teacher_id where " + queryArgs
	} else {
		query = "select (select count(*) from students where group_id=g.id) as student_count,g.id,g.name,g.code,g.year,g.profession_id,coalesce(g.teacher_id,0) as teacher_id,coalesce(t.first_name,'') as teacher_first_name, coalesce(t.last_name,'') as teacher_last_name, p.name as profession_name from groups as g join professions as p on p.id =g.profession_id left join teachers as t on t.id=g.teacher_id "
	}

	paginationQuery := fmt.Sprintf(`select count(*) from (%s) as s`, query)
	pagination, offset, err := g.getPagination(paginationQuery, input.Limit, input.Page)
	if err != nil {
		return models.GroupsAndPagination{}, err
	}

	query += fmt.Sprintf(` limit %d offset %d`, input.Limit, offset)

	var groups []models.Group
	err = g.studentDB.Select(&groups, query, args...)
	if err != nil {
		return models.GroupsAndPagination{}, err
	}

	return models.GroupsAndPagination{
		Groups:     groups,
		Pagination: pagination,
	}, nil
}
