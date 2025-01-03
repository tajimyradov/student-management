package admin

import (
	"github.com/jmoiron/sqlx"
	"student-management/internal/models"
)

type EmployeeRepository struct {
	db *sqlx.DB
}

func NewEmployeeRepository(db *sqlx.DB) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

func (e *EmployeeRepository) GetAllEmployeeRates() ([]models.EmployeeRate, error) {
	query := `select er.id, er.position_id,p.name as position_name, er.first_name, er.last_name, er."0.25",er."0.50",er."0.75",er."1.00",er.partial from employee_rate as er join positions as p  on er.position_id = p.id`
	rows, err := e.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var employeeRates []models.EmployeeRate
	for rows.Next() {
		var employeeRate models.EmployeeRate
		err = rows.Scan(&employeeRate.ID, &employeeRate.PositionID, &employeeRate.PositionName, &employeeRate.FirstName, &employeeRate.LastName, &employeeRate.S025, &employeeRate.S050, &employeeRate.S075, &employeeRate.S100, &employeeRate.Partial)
		if err != nil {
			return nil, err
		}
		employeeRates = append(employeeRates, employeeRate)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return employeeRates, nil
}

func (e *EmployeeRepository) AddEmployeeRate(employeeRate models.EmployeeRate) error {
	query := `insert into employee_rate (position_id, first_name, last_name, "0.25","0.50","0.75","1.00",partial) values ($1,$2,$3,$4,$5,$6,$7,$8) `

	_, err := e.db.Exec(query, employeeRate.PositionID, employeeRate.FirstName, employeeRate.LastName, employeeRate.S025, employeeRate.S050, employeeRate.S075, employeeRate.S100, employeeRate.Partial)
	if err != nil {
		return err
	}

	return nil
}

func (e *EmployeeRepository) DeleteEmployeeRate(id int) error {
	query := `delete from employee_rate where id=$1`
	_, err := e.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (e *EmployeeRepository) UpdateEmployeeRate(employeeRate models.EmployeeRate) error {
	query := `update employee_rate set position_id=$1, first_name=$2, last_name=$3, "0.25"=$4,"0.50"=$5,"0.75"=$6,"1.00"=$7,partial=$8 where id=$9`

	_, err := e.db.Exec(query, employeeRate.PositionID, employeeRate.FirstName, employeeRate.LastName, employeeRate.S025, employeeRate.S050, employeeRate.S075, employeeRate.S100, employeeRate.Partial, employeeRate.ID)
	if err != nil {
		return err
	}
	return nil
}

func (e *EmployeeRepository) GetAllPositions() ([]models.Position, error) {
	query := `select id, name from positions`
	var positions []models.Position
	err := e.db.Select(&positions, query)
	if err != nil {
		return nil, err
	}
	return positions, nil
}

func (e *EmployeeRepository) GetEmployeeRatByID(id int) (models.EmployeeRate, error) {

	query := `select er.id, er.position_id,p.name as position_name, er.first_name, er.last_name, er."0.25",er."0.50",er."0.75",er."1.00",er.partial from employee_rate as er join positions as p  on er.position_id = p.id where er.id=$1`
	var employeeRate models.EmployeeRate
	err := e.db.QueryRow(query, id).Scan(&employeeRate.ID, &employeeRate.PositionID, &employeeRate.PositionName, &employeeRate.FirstName, &employeeRate.LastName, &employeeRate.S025, &employeeRate.S050, &employeeRate.S075, &employeeRate.S100, &employeeRate.Partial)

	return employeeRate, err

}
