package services

import (
	"go.uber.org/zap"
	"student-management/internal/models"
	"student-management/internal/repositories"
)

type TimeService struct {
	repo   *repositories.TimeRepository
	logger *zap.Logger
}

func NewTimeService(repo *repositories.TimeRepository, logger *zap.Logger) *TimeService {
	return &TimeService{repo: repo, logger: logger}
}

func (s *TimeService) AddTime(input models.Time) (models.Time, error) {
	res, err := s.repo.AddTime(input)
	if err != nil {
		s.logger.Info("Time add failed", zap.Error(err))
		return models.Time{}, err
	}
	return res, nil
}

func (s *TimeService) UpdateTime(input models.Time) error {
	err := s.repo.UpdateTime(input)
	if err != nil {
		s.logger.Info("Time update failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *TimeService) DeleteTime(id int) error {
	err := s.repo.DeleteTime(id)
	if err != nil {
		s.logger.Info("Time delete failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *TimeService) GetTimeByID(id int) (models.Time, error) {
	res, err := s.repo.GetTimeByID(id)
	if err != nil {
		s.logger.Info("Time get by ID failed", zap.Error(err))
		return models.Time{}, err
	}
	return res, nil
}

func (s *TimeService) GetTimes(input models.TimeSearch) (models.TimesAndPagination, error) {
	res, err := s.repo.GetTimes(input)
	if err != nil {
		s.logger.Info("Times get failed", zap.Error(err))
		return models.TimesAndPagination{}, err
	}
	return res, nil
}
