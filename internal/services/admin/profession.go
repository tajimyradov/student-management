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

type ProfessionService struct {
	repo   *repository.ProfessionRepository
	logger *zap.Logger
}

func NewProfessionService(repo *repository.ProfessionRepository, logger *zap.Logger) *ProfessionService {
	return &ProfessionService{
		repo:   repo,
		logger: logger,
	}
}

func (p *ProfessionService) AddProfession(input models.Profession) (models.Profession, error) {
	result, err := p.repo.AddProfession(input)
	if err != nil {
		p.logger.Info("add profession failed", zap.Error(err))
		return models.Profession{}, err
	}
	return result, nil
}

func (p *ProfessionService) UpdateProfession(input models.Profession) error {
	err := p.repo.UpdateProfession(input)
	if err != nil {
		p.logger.Info("update profession failed", zap.Error(err))
		return err
	}
	return nil
}

func (p *ProfessionService) DeleteProfession(id int) error {
	err := p.repo.DeleteProfession(id)
	if err != nil {
		p.logger.Info("delete profession failed", zap.Error(err))
		return err
	}
	return nil
}

func (p *ProfessionService) GetProfessionByID(id int) (models.Profession, error) {
	res, err := p.repo.GetProfessionByID(id)
	if err != nil {
		p.logger.Info("get profession by id failed", zap.Error(err))
		return models.Profession{}, err
	}
	res.Files, err = p.repo.GetAllFiles(res.ID)
	if err != nil {
		p.logger.Info("get all files failed", zap.Error(err))
	}
	return res, nil
}

func (p *ProfessionService) GetProfessions(input models.ProfessionSearch) (models.ProfessionAndPagination, error) {
	res, err := p.repo.GetProfessions(input)
	if err != nil {
		p.logger.Info("get professions failed", zap.Error(err))
		return models.ProfessionAndPagination{}, err
	}
	for i, profession := range res.Professions {
		res.Professions[i].Files, err = p.repo.GetAllFiles(profession.ID)
		if err != nil {
			p.logger.Info("get all files failed", zap.Error(err))
		}
	}
	return res, nil
}

func (p *ProfessionService) UploadFile(id int, rFile *multipart.FileHeader, name string) error {
	ext := filepath.Ext(rFile.Filename)

	fileName := "files/" + uuid.New().String() + ext

	out, err := os.Create(fileName)
	if err != nil {
		p.logger.Info("Profession create failed", zap.Error(err))
		_ = os.Remove(fileName)
		return err
	}
	defer out.Close()

	ff, err := rFile.Open()
	if err != nil {
		p.logger.Info("Profession open failed", zap.Error(err))
		_ = os.Remove(fileName)
		return err
	}
	defer ff.Close()

	_, err = io.Copy(out, ff)
	if err != nil {
		p.logger.Info("Profession copy failed", zap.Error(err))
		_ = os.Remove(fileName)
		return err
	}

	err = p.repo.AddFile(id, fileName, name)
	if err != nil {
		p.logger.Info("Profession add failed", zap.Error(err))
		_ = os.Remove(fileName)
		return err
	}

	return nil
}

func (p *ProfessionService) DeleteFile(id int) error {
	file, err := p.repo.GetFileByID(id)
	if err != nil {
		p.logger.Info("Profession get file failed", zap.Error(err))
		return err
	}

	err = os.Remove(file)
	if err != nil {
		p.logger.Info("Profession remove file failed", zap.Error(err))
		return err
	}

	err = p.repo.DeleteFile(id)
	if err != nil {
		p.logger.Info("Profession delete failed", zap.Error(err))
		return err
	}
	return nil
}

func (p *ProfessionService) GetProfessionInfo(id int) (models.ProfessionInfo, error) {
	studentsCount, err := p.repo.GetStudents(id)
	if err != nil {
		p.logger.Info("get students failed", zap.Error(err))
		return models.ProfessionInfo{}, err
	}

	groups, err := p.repo.GetGroups(id)
	if err != nil {
		p.logger.Info("get groups failed", zap.Error(err))
		return models.ProfessionInfo{}, err
	}

	files, err := p.repo.GetAllFiles(id)
	if err != nil {
		p.logger.Info("get all files failed", zap.Error(err))
		return models.ProfessionInfo{}, err
	}

	return models.ProfessionInfo{
		StudentCount: studentsCount,
		Files:        files,
		Groups:       groups,
	}, nil
}
