package admin

import (
	"go.uber.org/zap"
	"student-management/internal/models"
	repository "student-management/internal/repositories/admin"
)

type RegionService struct {
	repo   *repository.RegionsRepository
	logger *zap.Logger
}

func NewRegionService(repo *repository.RegionsRepository, logger *zap.Logger) *RegionService {
	return &RegionService{
		repo:   repo,
		logger: logger,
	}
}

func (r *RegionService) GetRegions() ([]models.Region, error) {
	res, err := r.repo.GetRegions()
	if err != nil {
		r.logger.Info(`failed to get regions`, zap.Error(err))
		return nil, err
	}

	return res, nil
}
