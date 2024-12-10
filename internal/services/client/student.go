package client

import (
	"go.uber.org/zap"
	"student-management/internal/models"
	repository "student-management/internal/repositories/client"
)

type StudentService struct {
	repo   *repository.StudentRepository
	logger *zap.Logger
}

func NewStudentService(repo *repository.StudentRepository, logger *zap.Logger) *StudentService {
	return &StudentService{
		repo:   repo,
		logger: logger,
	}
}

func (s *StudentService) GetStudents(userID, roleID, lessonID, groupID, typeID int) ([]models.Student, error) {
	var students []models.Student
	var err error
	if roleID == 2 || roleID == 3 {
		students, err = s.repo.GetStudentsForTeacher(userID, lessonID, groupID, typeID)
	}

	if err != nil {
		s.logger.Info("failed to get students", zap.Error(err))
		return nil, err
	}

	return students, nil
}

func (s *StudentService) CheckForExistence(input models.Absence) error {
	err := s.repo.CheckForAbsence(input)
	if err != nil {
		s.logger.Info("failed to check if student exists", zap.Error(err))
		return err
	}
	return nil
}
