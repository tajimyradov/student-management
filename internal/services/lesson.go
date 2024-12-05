package services

import (
	"go.uber.org/zap"
	"student-management/internal/models"
	"student-management/internal/repositories"
)

type LessonService struct {
	repo   *repositories.LessonRepository
	logger *zap.Logger
}

func NewLessonService(repo *repositories.LessonRepository, logger *zap.Logger) *LessonService {
	return &LessonService{repo: repo, logger: logger}
}

func (s *LessonService) AddLesson(input models.Lesson) (models.Lesson, error) {
	res, err := s.repo.AddLesson(input)
	if err != nil {
		s.logger.Info("Lesson add failed", zap.Error(err))
		return models.Lesson{}, err
	}
	return res, nil
}

func (s *LessonService) UpdateLesson(input models.Lesson) error {
	err := s.repo.UpdateLesson(input)
	if err != nil {
		s.logger.Info("Lesson update failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *LessonService) DeleteLesson(id int) error {
	err := s.repo.DeleteLesson(id)
	if err != nil {
		s.logger.Info("Lesson delete failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *LessonService) GetLessonByID(id int) (models.Lesson, error) {
	res, err := s.repo.GetLessonByID(id)
	if err != nil {
		s.logger.Info("Lesson get by ID failed", zap.Error(err))
		return models.Lesson{}, err
	}
	return res, nil
}

func (s *LessonService) GetLessons(input models.LessonSearch) (models.LessonsAndPagination, error) {
	res, err := s.repo.GetLessons(input)
	if err != nil {
		s.logger.Info("Lessons get failed", zap.Error(err))
		return models.LessonsAndPagination{}, err
	}
	return res, nil
}
