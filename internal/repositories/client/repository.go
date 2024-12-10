package client

import (
	"github.com/jmoiron/sqlx"
	"student-management/internal/config"
)

type Repository struct {
	AuthRepository      *AuthRepository
	TimetableRepository *TimetableRepository
	StudentRepository   *StudentRepository
}

func NewRepository(studentDB *sqlx.DB, config *config.AppConfig) *Repository {
	return &Repository{
		AuthRepository:      NewAuthRepository(studentDB, config),
		TimetableRepository: NewTimetableRepository(studentDB),
		StudentRepository:   NewStudentRepository(studentDB),
	}
}
