package services

import (
	"go.uber.org/zap"
	"student-management/internal/models"
	repository "student-management/internal/repositories"
)

type ProfessionService struct {
	repo   *repository.ProfessionRepository
	logger *zap.Logger
}

func NewProfessionService(repo *repository.ProfessionRepository, logger *zap.Logger) *ProfessionService {
	return &ProfessionService{
		repo:   repo,
		logger: logger,
	}
}

func (d *ProfessionService) AddProfession(input models.Profession) (models.Profession, error) {
	result, err := d.repo.AddProfession(input)
	if err != nil {
		d.logger.Info("add profession failed", zap.Error(err))
		return models.Profession{}, err
	}
	return result, nil
}

func (d *ProfessionService) UpdateProfession(input models.Profession) error {
	err := d.repo.UpdateProfession(input)
	if err != nil {
		d.logger.Info("update profession failed", zap.Error(err))
		return err
	}
	return nil
}

func (d *ProfessionService) DeleteProfession(id int) error {
	err := d.repo.DeleteProfession(id)
	if err != nil {
		d.logger.Info("delete profession failed", zap.Error(err))
		return err
	}
	return nil
}

func (d *ProfessionService) GetProfessionByID(id int) (models.Profession, error) {
	res, err := d.repo.GetProfessionByID(id)
	if err != nil {
		d.logger.Info("get profession by id failed", zap.Error(err))
		return models.Profession{}, err
	}
	return res, nil
}

func (d *ProfessionService) GetProfessions(input models.ProfessionSearch) (models.ProfessionAndPagination, error) {
	res, err := d.repo.GetProfessions(input)
	if err != nil {
		d.logger.Info("get professions failed", zap.Error(err))
		return models.ProfessionAndPagination{}, err
	}
	return res, nil
}
