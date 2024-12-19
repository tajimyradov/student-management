package admin

import (
	"crypto/sha1"
	"fmt"
	"go.uber.org/zap"
	"image"
	"image/jpeg"
	"os"
	"strconv"
	"student-management/internal/config"
	"student-management/internal/models"
	repository "student-management/internal/repositories/admin"
)

type StudentService struct {
	repo   *repository.StudentRepository
	logger *zap.Logger
	config *config.AppConfig
}

func NewStudentService(repo *repository.StudentRepository, logger *zap.Logger, config *config.AppConfig) *StudentService {
	return &StudentService{
		repo:   repo,
		config: config,
		logger: logger,
	}
}

func (s *StudentService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(s.config.Secrets.AccessSecret)))
}

func (s *StudentService) AddStudent(student models.Student) (models.Student, error) {
	student.Password = s.generatePasswordHash(student.Password)
	res, err := s.repo.AddStudent(student)
	if err != nil {
		s.logger.Info("add student failed", zap.Error(err))
		return models.Student{}, err
	}

	return res, err
}

func (s *StudentService) DeleteStudent(id int) error {
	err := s.repo.DeleteStudent(id)
	if err != nil {
		s.logger.Info("delete student failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *StudentService) UpdateStudent(student models.Student) error {
	err := s.repo.UpdateStudent(student)
	if err != nil {
		s.logger.Info("update student failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *StudentService) GetStudents(input models.StudentSearch) (models.StudentsAndPagination, error) {
	students, err := s.repo.GetStudents(input)
	if err != nil {
		s.logger.Info("get students failed", zap.Error(err))
		return models.StudentsAndPagination{}, err
	}
	return students, nil
}

func (s *StudentService) GetStudentByID(id int) (models.Student, error) {
	res, err := s.repo.GetStudentByID(id)
	if err != nil {
		s.logger.Info("get student failed", zap.Error(err))
		return models.Student{}, err
	}
	return res, nil
}

func (s *StudentService) UploadImageOfStudent(image image.Image, id int) error {
	// Create a file path
	filePath := fmt.Sprintf("images/student/%d.jpg", id)

	// Create a file
	file, err := os.Create(filePath)
	if err != nil {
		s.logger.Info("create file failed", zap.Error(err))
		return err
	}
	defer file.Close()

	// Encode the image to JPEG format and write it to the file
	err = jpeg.Encode(file, image, &jpeg.Options{Quality: 90})
	if err != nil {
		s.logger.Info("encode jpeg failed", zap.Error(err))
		return err
	}

	err = s.repo.UpdateStudentsImage(s.config.Domains.Image+"images/student/"+strconv.Itoa(id)+".jpg", id)
	if err != nil {
		s.logger.Info("update teachers failed", zap.Error(err))
		return err
	}

	return nil
}
