package client

import (
	"github.com/jmoiron/sqlx"
	"student-management/internal/config"
)

type AuthRepository struct {
	config    *config.AppConfig
	studentDB *sqlx.DB
}

func NewAuthRepository(studentDB *sqlx.DB, config *config.AppConfig) *AuthRepository {
	return &AuthRepository{
		studentDB: studentDB,
		config:    config,
	}
}

func (a *AuthRepository) LoginAsStudent(username, password string) (int, error) {
	query := `select id from students where username = $1 and password = $2`
	var id int
	err := a.studentDB.QueryRow(query, username, password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthRepository) LoginAsTeacher(username, password string) (int, int, error) {
	query := `select id,role_id from teachers where username = $1 and password = $2`
	var id, roleId int
	err := a.studentDB.QueryRow(query, username, password).Scan(&id, &roleId)
	if err != nil {
		return 0, 0, err
	}
	return id, roleId, nil
}
