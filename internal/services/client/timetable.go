package client

import (
	"go.uber.org/zap"
	"student-management/internal/models"
	repository "student-management/internal/repositories/client"
	"time"
)

type TimetableService struct {
	repo   *repository.TimetableRepository
	logger *zap.Logger
}

func NewTimetableService(repo *repository.TimetableRepository, logger *zap.Logger) *TimetableService {
	return &TimetableService{
		repo:   repo,
		logger: logger,
	}
}

func (t *TimetableService) getWeekdayInt(s string) int {
	switch s {
	case time.Monday.String():
		return 1
	case time.Tuesday.String():
		return 2
	case time.Wednesday.String():
		return 3
	case time.Thursday.String():
		return 4
	case time.Friday.String():
		return 5
	case time.Saturday.String():
		return 6
	case time.Sunday.String():
		return 7
	}
	return 0
}

func (t *TimetableService) GetTimetable(userId, roleId int, weekday string) ([]models.Timetable, error) {
	weekdayInt := t.getWeekdayInt(weekday)
	var timetables []models.Timetable
	var err error
	if roleId == 1 {
		timetables, err = t.repo.GetTimetableOfStudent(userId, weekdayInt)
	} else if roleId == 2 || roleId == 3 {
		timetables, err = t.repo.GetTimetableOfTeacher(userId, weekdayInt)
	}
	if err != nil {
		t.logger.Info(`failed to get timetable`, zap.Error(err))
		return nil, err
	}

	return timetables, nil
}
