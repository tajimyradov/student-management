package admin

import (
	"go.uber.org/zap"
	"student-management/internal/models"
	"student-management/internal/repositories/admin"
)

type FacultyService struct {
	repo   *admin.FacultyRepository
	logger *zap.Logger
}

func NewFacultyService(repo *admin.FacultyRepository, logger *zap.Logger) *FacultyService {
	return &FacultyService{repo: repo, logger: logger}
}

func (f *FacultyService) AddFaculty(input models.Faculty) (models.Faculty, error) {
	res, err := f.repo.AddFaculty(input)
	if err != nil {
		f.logger.Info("Faculty add failed", zap.Error(err))
		return models.Faculty{}, err
	}
	return res, err
}

func (f *FacultyService) UpdateFaculty(input models.Faculty) error {
	err := f.repo.UpdateFaculty(input)
	if err != nil {
		f.logger.Info("Faculty update failed", zap.Error(err))
		return err
	}
	return nil
}

func (f *FacultyService) DeleteFaculty(id int) error {
	err := f.repo.DeleteFaculty(id)
	if err != nil {
		f.logger.Info("Faculty delete failed", zap.Error(err))
		return err
	}
	return nil
}

func (f *FacultyService) GetFacultyByID(id int) (models.Faculty, error) {
	faculty, err := f.repo.GetFacultyByID(id)
	if err != nil {
		f.logger.Info("Faculty get failed", zap.Error(err))
		return models.Faculty{}, err
	}
	return faculty, nil
}

func (f *FacultyService) GetFaculties(input models.FacultySearch) (models.FacultiesAndPagination, error) {
	facultiesAndPagination, err := f.repo.GetFaculties(input)
	if err != nil {
		f.logger.Info("Faculties get failed", zap.Error(err))
		return models.FacultiesAndPagination{}, err
	}
	return facultiesAndPagination, nil
}
