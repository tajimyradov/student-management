package admin

import (
	"go.uber.org/zap"
	"student-management/internal/models"
	repository "student-management/internal/repositories/admin"
)

type EmployeeService struct {
	repo   *repository.EmployeeRepository
	logger *zap.Logger
}

func NewEmployeeService(repo *repository.EmployeeRepository, logger *zap.Logger) *EmployeeService {
	return &EmployeeService{
		repo:   repo,
		logger: logger,
	}
}

func (e *EmployeeService) AddEmployeeRate(employee models.EmployeeRate) error {
	err := e.repo.AddEmployeeRate(employee)
	if err != nil {
		e.logger.Info(`failed to add employee rate`, zap.Error(err))
		return err
	}
	return nil
}

func (e *EmployeeService) DeleteEmployeeRate(id int) error {
	err := e.repo.DeleteEmployeeRate(id)
	if err != nil {
		e.logger.Info(`failed to delete employee rate`, zap.Error(err))
		return err
	}
	return nil
}

func (e *EmployeeService) GetAllEmployeeRate() ([]models.EmployeeRate, error) {
	res, err := e.repo.GetAllEmployeeRates()
	if err != nil {
		e.logger.Info(`failed to get all employee rates`, zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (e *EmployeeService) UpdateEmployeeRate(employee models.EmployeeRate) error {
	err := e.repo.UpdateEmployeeRate(employee)
	if err != nil {
		e.logger.Info(`failed to update employee rate`, zap.Error(err))
		return err
	}
	return nil
}

func (e *EmployeeService) GetAllPositions() ([]models.Position, error) {
	res, err := e.repo.GetAllPositions()
	if err != nil {
		e.logger.Info(`failed to get positions`, zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (e *EmployeeService) GetEmployeeRatByID(id int) (models.EmployeeRate, error) {
	res, err := e.repo.GetEmployeeRatByID(id)
	if err != nil {
		e.logger.Info(`failed to get employee rate by id`, zap.Error(err))
		return models.EmployeeRate{}, err
	}
	return res, nil
}
