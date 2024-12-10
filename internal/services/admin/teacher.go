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

type TeacherService struct {
	repo   *repository.TeacherRepository
	logger *zap.Logger
	config *config.AppConfig
}

func NewTeacherService(repo *repository.TeacherRepository, logger *zap.Logger, config *config.AppConfig) *TeacherService {
	return &TeacherService{
		repo:   repo,
		logger: logger,
		config: config,
	}
}

func (t *TeacherService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(t.config.Secrets.AccessSecret)))
}

func (t *TeacherService) AddTeacher(teacher models.Teacher) (models.Teacher, error) {
	teacher.Password = t.generatePasswordHash(teacher.Password)
	res, err := t.repo.AddTeacher(teacher)
	if err != nil {
		t.logger.Info("add teacher failed", zap.Error(err))
	}

	return res, err
}

func (t *TeacherService) DeleteTeacher(id int) error {
	err := t.repo.DeleteTeacher(id)
	if err != nil {
		t.logger.Info("delete teacher failed", zap.Error(err))
		return err
	}
	return nil
}

func (t *TeacherService) UpdateTeacher(teacher models.Teacher) error {
	err := t.repo.UpdateTeacher(teacher)
	if err != nil {
		t.logger.Info("update teacher failed", zap.Error(err))
		return err
	}
	return nil
}

func (t *TeacherService) GetTeachers(input models.TeacherSearch) (models.TeachersAndPagination, error) {
	teachers, err := t.repo.GetTeachers(input)
	if err != nil {
		t.logger.Info("get teachers failed", zap.Error(err))
		return models.TeachersAndPagination{}, err
	}
	return teachers, nil
}

func (t *TeacherService) GetTeacherByID(id int) (models.Teacher, error) {
	res, err := t.repo.GetTeacherByID(id)
	if err != nil {
		t.logger.Info("get teacher failed", zap.Error(err))
		return models.Teacher{}, err
	}
	return res, nil
}

func (t *TeacherService) UploadImageOfTeacher(image image.Image, id int) error {
	// Create a file path
	filePath := fmt.Sprintf("images/%d.jpg", id)

	// Create a file
	file, err := os.Create(filePath)
	if err != nil {
		t.logger.Info("create file failed", zap.Error(err))
		return err
	}
	defer file.Close()

	// Encode the image to JPEG format and write it to the file
	err = jpeg.Encode(file, image, &jpeg.Options{Quality: 90})
	if err != nil {
		t.logger.Info("encode jpeg failed", zap.Error(err))
		return err
	}

	err = t.repo.UpdateTeachersImage(strconv.Itoa(id), id)
	if err != nil {
		t.logger.Info("update teachers failed", zap.Error(err))
		return err
	}

	return nil
}
