package admin

import (
	"github.com/jmoiron/sqlx"
	"student-management/internal/models"
)

type StatisticsRepository struct {
	studentDB *sqlx.DB
}

func NewStatisticsRepository(db *sqlx.DB) *StatisticsRepository {
	return &StatisticsRepository{
		studentDB: db,
	}
}

func (s *StatisticsRepository) GetStatisticsByGender(facultyID int) ([]models.FacultyGender, error) {
	query := `
			select
				g.year,
				   s.gender,
				   count(*)
			
			from students as s
					 join groups as g on s.group_id = g.id
					 join professions as p on g.profession_id = p.id
					 join departments as d on p.department_id = d.id
					 join faculties as f on d.faculty_id = f.id
			where f.id = $1
			group by g.year, s.gender
			order by g.year,s.gender
`

	rows, err := s.studentDB.Query(query, facultyID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	res := make(map[int]models.FacultyGender)

	for rows.Next() {
		var year, count, c, cc int
		var gender bool
		err = rows.Scan(&year, &gender, &count)
		if err != nil {
			return nil, err
		}

		c = count
		if !gender {
			c = 0
			cc = count
		}

		res[year] = models.FacultyGender{
			Year:   year,
			Male:   c + res[year].Male,
			Female: cc + res[year].Female,
		}

	}

	var result []models.FacultyGender

	for _, gender := range res {
		result = append(result, gender)
	}

	return result, nil
}

func (s *StatisticsRepository) GetStatisticsProfession(facultyID int) ([]models.FacultyProfession, error) {
	query := `
	select
		g.year,
		   p.id as profession_id,
		   p.name as profession_name,
		   count(s.id)
	
	from students as s
			 join groups as g on s.group_id = g.id
			 join professions as p on g.profession_id = p.id
			 join departments as d on p.department_id = d.id
			 join faculties as f on d.faculty_id = f.id
	where f.id = $1
	group by g.year, p.id
	order by g.year,p.id
`
	var res []models.FacultyProfession
	err := s.studentDB.Select(&res, query, facultyID)
	return res, err
}

func (s *StatisticsRepository) GetStatisticsByAge(facultyID int) ([]models.FacultyAge, error) {
	query := `
select EXTRACT(YEAR FROM AGE((extract(year from now()) || '-12-31')::date, s.birth_date)) as age,
       count(*),
       s.gender
from students as s
         join groups as g on s.group_id = g.id
         join professions as p on g.profession_id = p.id
         join departments as d on p.department_id = d.id
         join faculties as f on d.faculty_id = f.id
where f.id = $1
group by age, s.gender
order by age,s.gender
`

	rows, err := s.studentDB.Query(query, facultyID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	res := make(map[int]models.FacultyAge)

	for rows.Next() {
		var age, count, c, cc int
		var gender bool
		err = rows.Scan(&age, &count, &gender)
		if err != nil {
			return nil, err
		}

		c = count
		if !gender {
			c = 0
			cc = count
		}

		res[age] = models.FacultyAge{
			Age:    age,
			Male:   c + res[age].Male,
			Female: cc + res[age].Female,
		}

	}

	var result []models.FacultyAge

	for _, gender := range res {
		result = append(result, gender)
	}

	return result, nil
}

func (s *StatisticsRepository) GetStatisticsByRegions(facultyID int) ([]models.FacultyRegion, error) {
	query := `
select
    g.year,
       r.id as region_id,
       r.name as region_name,
       count(*)

from students as s
         join groups as g on s.group_id = g.id
         join professions as p on g.profession_id = p.id
         join departments as d on p.department_id = d.id
         join faculties as f on d.faculty_id = f.id
         join regions as r on s.region_id = r.id
where f.id = $1
group by g.year, r.id
order by g.year,r.id
`

	var res []models.FacultyRegion
	err := s.studentDB.Select(&res, query, facultyID)
	return res, err
}
