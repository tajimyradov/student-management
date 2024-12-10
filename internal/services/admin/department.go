package admin

import (
	"go.uber.org/zap"
	"student-management/internal/models"
	repository "student-management/internal/repositories/admin"
)

type DepartmentService struct {
	repo   *repository.DepartmentRepository
	logger *zap.Logger
}

func NewDepartmentService(repo *repository.DepartmentRepository, logger *zap.Logger) *DepartmentService {
	return &DepartmentService{
		repo:   repo,
		logger: logger,
	}
}

func (d *DepartmentService) AddDepartment(input models.Department) (models.Department, error) {
	result, err := d.repo.AddDepartment(input)
	if err != nil {
		d.logger.Info("add department failed", zap.Error(err))
		return models.Department{}, err
	}
	return result, nil
}

func (d *DepartmentService) UpdateDepartment(input models.Department) error {
	err := d.repo.UpdateDepartment(input)
	if err != nil {
		d.logger.Info("update department failed", zap.Error(err))
		return err
	}
	return nil
}

func (d *DepartmentService) DeleteDepartment(id int) error {
	err := d.repo.DeleteDepartment(id)
	if err != nil {
		d.logger.Info("delete department failed", zap.Error(err))
		return err
	}
	return nil
}

func (d *DepartmentService) GetDepartmentByID(id int) (models.Department, error) {
	res, err := d.repo.GetDepartmentById(id)
	if err != nil {
		d.logger.Info("get department by id failed", zap.Error(err))
		return models.Department{}, err
	}
	return res, nil
}

func (d *DepartmentService) GetDepartments(input models.DepartmentSearch) (models.DepartmentAndPagination, error) {
	res, err := d.repo.GetDepartments(input)
	if err != nil {
		d.logger.Info("get departments failed", zap.Error(err))
		return models.DepartmentAndPagination{}, err
	}
	return res, nil
}
