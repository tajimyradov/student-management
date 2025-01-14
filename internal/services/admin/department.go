package admin

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"student-management/internal/models"
	repository "student-management/internal/repositories/admin"
)

type DepartmentService struct {
	repo   *repository.DepartmentRepository
	logger *zap.Logger
}

func NewDepartmentService(repo *repository.DepartmentRepository, logger *zap.Logger) *DepartmentService {
	return &DepartmentService{
		repo:   repo,
		logger: logger,
	}
}

func (d *DepartmentService) AddDepartment(input models.Department) (models.Department, error) {
	result, err := d.repo.AddDepartment(input)
	if err != nil {
		d.logger.Info("add department failed", zap.Error(err))
		return models.Department{}, err
	}
	return result, nil
}

func (d *DepartmentService) UpdateDepartment(input models.Department) error {
	err := d.repo.UpdateDepartment(input)
	if err != nil {
		d.logger.Info("update department failed", zap.Error(err))
		return err
	}
	return nil
}

func (d *DepartmentService) DeleteDepartment(id int) error {
	err := d.repo.DeleteDepartment(id)
	if err != nil {
		d.logger.Info("delete department failed", zap.Error(err))
		return err
	}
	return nil
}

func (d *DepartmentService) GetDepartmentByID(id int) (models.Department, error) {
	res, err := d.repo.GetDepartmentById(id)
	if err != nil {
		d.logger.Info("get department by id failed", zap.Error(err))
		return models.Department{}, err
	}
	res.Files, err = d.repo.GetAllFiles(res.ID)
	if err != nil {
		d.logger.Info("get department by id failed", zap.Error(err))
	}
	return res, nil
}

func (d *DepartmentService) GetDepartments(input models.DepartmentSearch) (models.DepartmentAndPagination, error) {
	res, err := d.repo.GetDepartments(input)
	if err != nil {
		d.logger.Info("get departments failed", zap.Error(err))
		return models.DepartmentAndPagination{}, err
	}
	for i, department := range res.Departments {
		res.Departments[i].Files, err = d.repo.GetAllFiles(department.ID)
		if err != nil {
			d.logger.Info("get departments failed", zap.Error(err))
		}
	}
	return res, nil
}

func (d *DepartmentService) UploadFile(id int, rFile *multipart.FileHeader, name string) error {
	ext := filepath.Ext(rFile.Filename)

	fileName := "files/" + uuid.New().String() + ext

	out, err := os.Create(fileName)
	if err != nil {
		d.logger.Info("Department create failed", zap.Error(err))
		_ = os.Remove(fileName)
		return err
	}
	defer out.Close()

	ff, err := rFile.Open()
	if err != nil {
		d.logger.Info("Department open failed", zap.Error(err))
		_ = os.Remove(fileName)
		return err
	}
	defer ff.Close()

	_, err = io.Copy(out, ff)
	if err != nil {
		d.logger.Info("Department copy failed", zap.Error(err))
		_ = os.Remove(fileName)
		return err
	}

	err = d.repo.AddFile(id, fileName, name)
	if err != nil {
		d.logger.Info("Department add failed", zap.Error(err))
		_ = os.Remove(fileName)
		return err
	}

	return nil
}

func (d *DepartmentService) DeleteFile(id int) error {
	file, err := d.repo.GetFileByID(id)
	if err != nil {
		d.logger.Info("Department get file failed", zap.Error(err))
		return err
	}

	err = os.Remove(file)
	if err != nil {
		d.logger.Info("Department remove file failed", zap.Error(err))
		return err
	}

	err = d.repo.DeleteFile(id)
	if err != nil {
		d.logger.Info("Department delete failed", zap.Error(err))
		return err
	}
	return nil
}

func (d *DepartmentService) GetDepartmentInfo(id int) (models.DepartmentInfo, error) {
	professions, err := d.repo.GetProfessions(id)
	if err != nil {
		d.logger.Info("Department get professions failed", zap.Error(err))
		return models.DepartmentInfo{}, err
	}

	teachers, err := d.repo.GetTeachers(id)
	if err != nil {
		d.logger.Info("Department get teachers failed", zap.Error(err))
		return models.DepartmentInfo{}, err
	}

	studentsCount, err := d.repo.GetStudentsCount(id)
	if err != nil {
		d.logger.Info("Department get students count failed", zap.Error(err))
		return models.DepartmentInfo{}, err
	}

	groupCount, err := d.repo.GetGroupCount(id)
	if err != nil {
		d.logger.Info("Department get group count failed", zap.Error(err))
		return models.DepartmentInfo{}, err
	}

	return models.DepartmentInfo{
		Professions:  professions,
		Teachers:     teachers,
		StudentCount: studentsCount,
		GroupCount:   groupCount,
	}, nil
}
