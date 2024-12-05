package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"student-management/internal/models"
)

type AuditoryRepository struct {
	db *sqlx.DB
}

func NewAuditoryRepository(db *sqlx.DB) *AuditoryRepository {
	return &AuditoryRepository{
		db: db,
	}
}

// AddAuditory adds a new auditory to the database
func (r *AuditoryRepository) AddAuditory(auditory models.Auditory) (models.Auditory, error) {
	query := `INSERT INTO auditories(name) VALUES ($1) RETURNING id`
	err := r.db.QueryRow(query, auditory.Name).Scan(&auditory.ID)
	return auditory, err
}

// UpdateAuditory updates the name of an auditory
func (r *AuditoryRepository) UpdateAuditory(auditory models.Auditory) error {
	query := `UPDATE auditories SET name=$1 WHERE id=$2`
	_, err := r.db.Exec(query, auditory.Name, auditory.ID)
	return err
}

// DeleteAuditory deletes an auditory by ID
func (r *AuditoryRepository) DeleteAuditory(id int) error {
	query := `DELETE FROM auditories WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}

// GetAuditoryByID fetches an auditory by its ID
func (r *AuditoryRepository) GetAuditoryByID(id int) (models.Auditory, error) {
	var auditory models.Auditory
	query := `SELECT id, name FROM auditories WHERE id=$1`
	err := r.db.Get(&auditory, query, id)
	return auditory, err
}

// GetAuditories retrieves auditories with optional search and pagination
func (r *AuditoryRepository) GetAuditories(input models.AuditorySearch) (models.AuditoriesAndPagination, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name ILIKE '%%%s%%'", input.Name))
	}

	if input.ID != 0 {
		setValues = append(setValues, fmt.Sprintf("id = $%d", argID))
		args = append(args, input.ID)
		argID++
	}

	queryArgs := strings.Join(setValues, " AND ")

	var query string
	if len(setValues) > 0 {
		query = fmt.Sprintf("SELECT id, name FROM auditories WHERE %s", queryArgs)
	} else {
		query = "SELECT id, name FROM auditories"
	}

	// Pagination
	paginationQuery := fmt.Sprintf("SELECT COUNT(*) FROM (%s) AS count_query", query)
	pagination, offset, err := r.getPagination(paginationQuery, input.Limit, input.Page)
	if err != nil {
		return models.AuditoriesAndPagination{}, err
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", input.Limit, offset)

	var auditories []models.Auditory
	err = r.db.Select(&auditories, query, args...)
	if err != nil {
		return models.AuditoriesAndPagination{}, err
	}

	return models.AuditoriesAndPagination{
		Auditories: auditories,
		Pagination: pagination,
	}, nil
}

// Helper: Pagination logic
func (r *AuditoryRepository) getPagination(query string, limit, page int) (models.Pagination, int, error) {
	var count, add int
	err := r.db.QueryRow(query).Scan(&count)
	if err != nil {
		return models.Pagination{}, 0, err
	}

	if count%limit > 0 {
		add = 1
	}

	return models.Pagination{
		TotalPages:  count/limit + add,
		TotalCount:  count,
		Limit:       limit,
		CurrentPage: page,
	}, (page - 1) * limit, nil
}
