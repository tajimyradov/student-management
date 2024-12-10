package admin

import (
	"go.uber.org/zap"
	"student-management/internal/models"
	repository "student-management/internal/repositories/admin"
)

type TimetableService struct {
	repo   *repository.TimetableRepository
	logger *zap.Logger
}

func NewTimetableService(repo *repository.TimetableRepository, logger *zap.Logger) *TimetableService {
	return &TimetableService{
		logger: logger,
		repo:   repo,
	}
}

func (t *TimetableService) AddStudentTeacherLessonBinding(input models.LessonTeacherStudent) error {
	err := t.repo.AddStudentTeacherLessonBinding(input)
	if err != nil {
		t.logger.Info("add student teacher lesson binding failed", zap.Error(err))
		return err
	}
	return nil
}

func (t *TimetableService) DeleteStudentTeacherLessonBinding(input models.LessonTeacherStudent) error {
	err := t.repo.DeleteStudentTeacherLessonBinding(input)
	if err != nil {
		t.logger.Info("delete student teacher lesson binding failed", zap.Error(err))
		return err
	}
	return nil
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
