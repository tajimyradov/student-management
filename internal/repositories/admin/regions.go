package admin

import (
	"github.com/jmoiron/sqlx"
	"student-management/internal/models"
)

type RegionsRepository struct {
	studentDB *sqlx.DB
}

func NewRegionsRepository(db *sqlx.DB) *RegionsRepository {
	return &RegionsRepository{
		studentDB: db,
	}
}

func (r *RegionsRepository) GetRegions() ([]models.Region, error) {
	var regions []models.Region
	err := r.studentDB.Select(&regions, "SELECT id,name FROM regions")
	return regions, err
}
