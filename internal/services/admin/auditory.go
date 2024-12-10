package admin

import (
	"go.uber.org/zap"
	"student-management/internal/models"
	"student-management/internal/repositories/admin"
)

type AuditoryService struct {
	repo   *admin.AuditoryRepository
	logger *zap.Logger
}

func NewAuditoryService(repo *admin.AuditoryRepository, logger *zap.Logger) *AuditoryService {
	return &AuditoryService{repo: repo, logger: logger}
}

func (s *AuditoryService) AddAuditory(input models.Auditory) (models.Auditory, error) {
	res, err := s.repo.AddAuditory(input)
	if err != nil {
		s.logger.Info("Auditory add failed", zap.Error(err))
		return models.Auditory{}, err
	}
	return res, nil
}

func (s *AuditoryService) UpdateAuditory(input models.Auditory) error {
	err := s.repo.UpdateAuditory(input)
	if err != nil {
		s.logger.Info("Auditory update failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *AuditoryService) DeleteAuditory(id int) error {
	err := s.repo.DeleteAuditory(id)
	if err != nil {
		s.logger.Info("Auditory delete failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *AuditoryService) GetAuditoryByID(id int) (models.Auditory, error) {
	res, err := s.repo.GetAuditoryByID(id)
	if err != nil {
		s.logger.Info("Auditory get by ID failed", zap.Error(err))
		return models.Auditory{}, err
	}
	return res, nil
}

func (s *AuditoryService) GetAuditories(input models.AuditorySearch) (models.AuditoriesAndPagination, error) {
	res, err := s.repo.GetAuditories(input)
	if err != nil {
		s.logger.Info("Auditories get failed", zap.Error(err))
		return models.AuditoriesAndPagination{}, err
	}
	return res, nil
}
