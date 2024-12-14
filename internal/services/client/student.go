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

func (s *StudentService) GetStudents(roleID, groupID int) ([]models.Student, error) {
	var students []models.Student
	var err error
	if roleID == 2 || roleID == 3 {
		students, err = s.repo.GetStudentsForTeacher(groupID)
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

func (s *StudentService) GetFaculties() ([]models.Faculty, error) {
	res, err := s.repo.GetFaculties()
	if err != nil {
		s.logger.Info("failed to get faculties", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *StudentService) GetDepartments(facultyID int) ([]models.Department, error) {
	res, err := s.repo.GetDepartments(facultyID)
	if err != nil {
		s.logger.Info("failed to get departments", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *StudentService) GetGroups(departmentID int) ([]models.Group, error) {
	res, err := s.repo.GetGroups(departmentID)
	if err != nil {
		s.logger.Info("failed to get groups", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *StudentService) GetLessons() ([]models.Lesson, error) {
	res, err := s.repo.GetLessons()
	if err != nil {
		s.logger.Info("failed to get lessons", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *StudentService) GetTypes() ([]models.Type, error) {
	res, err := s.repo.GetTypes()
	if err != nil {
		s.logger.Info("failed to get types", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *StudentService) GetTimes() ([]models.Time, error) {
	res, err := s.repo.GetTimes()
	if err != nil {
		s.logger.Info("failed to get times", zap.Error(err))
		return nil, err
	}
	return res, nil
}
