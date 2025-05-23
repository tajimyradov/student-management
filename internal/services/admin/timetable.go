package admin

import (
	"go.uber.org/zap"
	"student-management/internal/models"
	repository "student-management/internal/repositories/admin"
	"time"
)

type TimetableService struct {
	lastSync time.Time
	repo     *repository.TimetableRepository
	logger   *zap.Logger
}

func NewTimetableService(repo *repository.TimetableRepository, logger *zap.Logger) *TimetableService {
	return &TimetableService{
		lastSync: time.Now(),
		logger:   logger,
		repo:     repo,
	}
}

func (t *TimetableService) AddTimetable(input models.Timetable) error {
	err := t.repo.AddTimetable(input)
	if err != nil {
		t.logger.Info("add timetable failed", zap.Error(err))
		return err
	}
	return nil
}

func (t *TimetableService) DeleteTimetable(id int) error {
	err := t.repo.DeleteTimetable(id)
	if err != nil {
		t.logger.Info("delete timetable failed", zap.Error(err))
		return err
	}
	return nil
}

func (t *TimetableService) GetTimetableOfGroup(groupId int) ([]models.Timetable, error) {
	timetables, err := t.repo.GetTimetableOfGroup(groupId)
	if err != nil {
		t.logger.Info("get timetable of group failed", zap.Error(err))
		return nil, err
	}
	return timetables, nil
}

func (t *TimetableService) GetStudentTeacherLessonBinding(teacherID, lessonID int) (models.LessonTeacherStudent, error) {
	res, err := t.repo.GetStudentTeacherLessonBinding(teacherID, lessonID)
	if err != nil {
		t.logger.Info("get student teacher lesson binding failed", zap.Error(err))
		return models.LessonTeacherStudent{}, err
	}
	return res, nil
}

func (t *TimetableService) GetAbsences(input models.AbsenceSearch) ([]models.Absence, error) {
	if time.Now().Sub(t.lastSync) >= time.Second*15 {
		err := t.repo.Sync()
		if err != nil {
			t.logger.Info("sync timetable failed", zap.Error(err))
			return nil, err
		}
	}

	res, err := t.repo.GetAbsences(input)
	if err != nil {
		t.logger.Info("get absences failed", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (t *TimetableService) UpdateAbsence(status, id int) error {
	err := t.repo.UpdateAbsence(status, id)
	if err != nil {
		t.logger.Info("update absence failed", zap.Error(err))
		return err
	}
	return nil
}

//func (t *TimetableService) Sync() error {
//	err := t.repo.Sync()
//	if err != nil {
//		t.logger.Info("sync timetable failed", zap.Error(err))
//		return err
//	}
//	return nil
//}

func (t *TimetableService) GetAbsenceByID(id int) (models.Absence, error) {
	res, err := t.repo.GetAbsenceByID(id)
	if err != nil {
		t.logger.Info("get absence failed", zap.Error(err))
		return models.Absence{}, err
	}
	return res, nil
}
