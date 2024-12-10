package admin

import (
	"go.uber.org/zap"
	"student-management/internal/models"
	repository "student-management/internal/repositories/admin"
)

type GroupService struct {
	repo   *repository.GroupRepository
	logger *zap.Logger
}

func NewGroupService(repo *repository.GroupRepository, logger *zap.Logger) *GroupService {
	return &GroupService{
		repo:   repo,
		logger: logger,
	}
}

func (g *GroupService) AddGroup(input models.Group) (models.Group, error) {
	result, err := g.repo.AddGroup(input)
	if err != nil {
		g.logger.Info("add group failed", zap.Error(err))
		return models.Group{}, err
	}
	return result, nil
}

func (g *GroupService) UpdateGroup(input models.Group) error {
	err := g.repo.UpdateGroup(input)
	if err != nil {
		g.logger.Info("update group failed", zap.Error(err))
		return err
	}
	return nil
}

func (g *GroupService) DeleteGroup(groupID int) error {
	err := g.repo.DeleteGroup(groupID)
	if err != nil {
		g.logger.Info("delete group failed", zap.Error(err))
		return err
	}
	return nil
}

func (g *GroupService) GetGroupByID(id int) (models.Group, error) {
	result, err := g.repo.GetGroupByID(id)
	if err != nil {
		g.logger.Info("get group failed", zap.Error(err))
		return models.Group{}, err
	}
	return result, nil
}

func (g *GroupService) GetGroups(input models.GroupSearch) (models.GroupsAndPagination, error) {
	result, err := g.repo.GetGroups(input)
	if err != nil {
		g.logger.Info("get groups failed", zap.Error(err))
		return models.GroupsAndPagination{}, err
	}
	return result, nil
}
