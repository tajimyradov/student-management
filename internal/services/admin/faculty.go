package admin

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"student-management/internal/models"
	"student-management/internal/repositories/admin"
)

type FacultyService struct {
	repo   *admin.FacultyRepository
	logger *zap.Logger
}

func NewFacultyService(repo *admin.FacultyRepository, logger *zap.Logger) *FacultyService {
	return &FacultyService{repo: repo, logger: logger}
}

func (f *FacultyService) AddFaculty(input models.Faculty) (models.Faculty, error) {
	res, err := f.repo.AddFaculty(input)
	if err != nil {
		f.logger.Info("Faculty add failed", zap.Error(err))
		return models.Faculty{}, err
	}
	return res, err
}

func (f *FacultyService) UpdateFaculty(input models.Faculty) error {
	err := f.repo.UpdateFaculty(input)
	if err != nil {
		f.logger.Info("Faculty update failed", zap.Error(err))
		return err
	}
	return nil
}

func (f *FacultyService) DeleteFaculty(id int) error {
	err := f.repo.DeleteFaculty(id)
	if err != nil {
		f.logger.Info("Faculty delete failed", zap.Error(err))
		return err
	}
	return nil
}

func (f *FacultyService) GetFacultyByID(id int) (models.Faculty, error) {
	faculty, err := f.repo.GetFacultyByID(id)
	if err != nil {
		f.logger.Info("Faculty get failed", zap.Error(err))
		return models.Faculty{}, err
	}
	faculty.Files, err = f.repo.GetAllFiles(faculty.ID)
	if err != nil {
		f.logger.Info("Faculty get failed", zap.Error(err))
	}
	return faculty, nil
}

func (f *FacultyService) GetFaculties(input models.FacultySearch) (models.FacultiesAndPagination, error) {
	facultiesAndPagination, err := f.repo.GetFaculties(input)
	if err != nil {
		f.logger.Info("Faculties get failed", zap.Error(err))
		return models.FacultiesAndPagination{}, err
	}
	for i, faculty := range facultiesAndPagination.Faculties {
		facultiesAndPagination.Faculties[i].Files, err = f.repo.GetAllFiles(faculty.ID)
		if err != nil {
			f.logger.Info("Faculty get failed", zap.Error(err))
		}
	}
	return facultiesAndPagination, nil
}

func (f *FacultyService) UploadFile(id int, rFile *multipart.FileHeader, name string) error {
	ext := filepath.Ext(rFile.Filename)

	fileName := "files/" + uuid.New().String() + ext

	out, err := os.Create(fileName)
	if err != nil {
		f.logger.Info("Faculty create failed", zap.Error(err))
		_ = os.Remove(fileName)
		return err
	}
	defer out.Close()

	ff, err := rFile.Open()
	if err != nil {
		f.logger.Info("Faculty open failed", zap.Error(err))
		_ = os.Remove(fileName)
		return err
	}
	defer ff.Close()

	_, err = io.Copy(out, ff)
	if err != nil {
		f.logger.Info("Faculty copy failed", zap.Error(err))
		_ = os.Remove(fileName)
		return err
	}

	err = f.repo.AddFile(id, fileName, name)
	if err != nil {
		f.logger.Info("Faculty add failed", zap.Error(err))
		_ = os.Remove(fileName)
		return err
	}

	return nil
}

func (f *FacultyService) DeleteFile(fileID int) error {
	file, err := f.repo.GetFileByID(fileID)
	if err != nil {
		f.logger.Info("Faculty get file failed", zap.Error(err))
		return err
	}

	err = os.Remove(file)
	if err != nil {
		f.logger.Info("Faculty remove file failed", zap.Error(err))
		return err
	}

	err = f.repo.DeleteFile(fileID)
	if err != nil {
		f.logger.Info("Faculty delete failed", zap.Error(err))
		return err
	}
	return nil
}

func (f *FacultyService) GetFacultyInfo(facultyID int) (models.FacultyInfo, error) {
	professions, err := f.repo.GetProfessions(facultyID)
	if err != nil {
		f.logger.Info("Faculty get professions failed", zap.Error(err))
		return models.FacultyInfo{}, err
	}

	departments, err := f.repo.GetDepartments(facultyID)
	if err != nil {
		f.logger.Info("Faculty get departments failed", zap.Error(err))
		return models.FacultyInfo{}, err
	}

	teachers, err := f.repo.GetTeachers(facultyID)
	if err != nil {
		f.logger.Info("Faculty get teachers failed", zap.Error(err))
		return models.FacultyInfo{}, err
	}

	studentsCount, err := f.repo.GetStudentsCount(facultyID)
	if err != nil {
		f.logger.Info("Faculty get students count failed", zap.Error(err))
		return models.FacultyInfo{}, err
	}

	groupCount, err := f.repo.GetGroupCount(facultyID)
	if err != nil {
		f.logger.Info("Faculty get group count failed", zap.Error(err))
		return models.FacultyInfo{}, err
	}

	return models.FacultyInfo{
		Professions:  professions,
		Departments:  departments,
		Teachers:     teachers,
		StudentCount: studentsCount,
		GroupCount:   groupCount,
	}, nil
}
