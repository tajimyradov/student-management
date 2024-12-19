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

func (a *AuthRepository) LoginAsStudent(username, password string) (int, string, string, string, error) {
	query := `select id,first_name, last_name,image from students where username = $1 and password = $2`
	var id int
	var firstName, lastName, image string
	err := a.studentDB.QueryRow(query, username, password).Scan(&id, &firstName, &lastName, &image)
	if err != nil {
		return 0, "", "", "", err
	}
	return id, firstName, lastName, image, nil
}

func (a *AuthRepository) LoginAsTeacher(username, password string) (int, int, string, string, string, error) {
	query := `select id,role_id,first_name, last_name,image from teachers where username = $1 and password = $2`
	var id, roleId int
	var firstName, lastName, image string
	err := a.studentDB.QueryRow(query, username, password).Scan(&id, &roleId, &firstName, &lastName, &image)
	if err != nil {
		return 0, 0, "", "", "", err
	}
	return id, roleId, firstName, lastName, image, nil
}
