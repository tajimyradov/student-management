package admin

import (
	"go.uber.org/zap"
	"student-management/internal/models"
	repository "student-management/internal/repositories/admin"
)

type StatisticsService struct {
	repo   *repository.StatisticsRepository
	logger *zap.Logger
}

func NewStatisticsService(repo *repository.StatisticsRepository, logger *zap.Logger) *StatisticsService {
	return &StatisticsService{
		repo:   repo,
		logger: logger,
	}
}

func (s *StatisticsService) GetStatisticsByGender(facultyID int) ([]models.FacultyGender, error) {
	res, err := s.repo.GetStatisticsByGender(facultyID)
	if err != nil {
		s.logger.Info("failed to get statistics by gender", zap.Int("faculty_id", facultyID), zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *StatisticsService) GetStatisticsProfession(facultyID int) ([]models.FacultyProfession, error) {
	res, err := s.repo.GetStatisticsProfession(facultyID)
	if err != nil {
		s.logger.Info(`failed to get statistics profession`, zap.Int("faculty_id", facultyID), zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *StatisticsService) GetStatisticsByAge(facultyID int) ([]models.FacultyAge, error) {
	res, err := s.repo.GetStatisticsByAge(facultyID)
	if err != nil {
		s.logger.Info(`failed to get statistics by age`, zap.Int("faculty_id", facultyID), zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *StatisticsService) GetStatisticsByRegions(facultyID int) ([]models.FacultyRegion, error) {
	res, err := s.repo.GetStatisticsByRegions(facultyID)
	if err != nil {
		s.logger.Info(`failed to get statistics by region`, zap.Int("faculty_id", facultyID), zap.Error(err))
		return nil, err
	}
	return res, nil
}
