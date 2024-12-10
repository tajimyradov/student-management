package repositories

import (
	"github.com/jmoiron/sqlx"
	"student-management/internal/config"
	"student-management/internal/repositories/admin"
	"student-management/internal/repositories/client"
)

type Repositories struct {
	AdminRepos  *admin.Repository
	ClientRepos *client.Repository
}

func NewRepositories(studentDB *sqlx.DB, appConfig *config.AppConfig) *Repositories {
	return &Repositories{
		AdminRepos:  admin.NewRepository(studentDB, appConfig),
		ClientRepos: client.NewRepository(studentDB, appConfig),
	}
}
