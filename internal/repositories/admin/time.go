package admin

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"student-management/internal/models"
)

type TimeRepository struct {
	db *sqlx.DB
}

func NewTimeRepository(db *sqlx.DB) *TimeRepository {
	return &TimeRepository{
		db: db,
	}
}

// AddTime adds a new time record to the database
func (r *TimeRepository) AddTime(time models.Time) (models.Time, error) {
	query := `INSERT INTO times(start_time, end_time) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(query, time.StartTime, time.EndTime).Scan(&time.ID)
	return time, err
}

// UpdateTime updates an existing time record
func (r *TimeRepository) UpdateTime(time models.Time) error {
	query := `UPDATE times SET start_time=$1, end_time=$2 WHERE id=$3`
	_, err := r.db.Exec(query, time.StartTime, time.EndTime, time.ID)
	return err
}

// DeleteTime deletes a time record by ID
func (r *TimeRepository) DeleteTime(id int) error {
	query := `DELETE FROM times WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}

// GetTimeByID fetches a time record by its ID
func (r *TimeRepository) GetTimeByID(id int) (models.Time, error) {
	var time models.Time
	query := `SELECT id, start_time, end_time FROM times WHERE id=$1`
	err := r.db.Get(&time, query, id)
	return time, err
}

// GetTimes retrieves time records with optional search and pagination
func (r *TimeRepository) GetTimes(input models.TimeSearch) (models.TimesAndPagination, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.StartTime != "" {
		setValues = append(setValues, fmt.Sprintf("start_time >= $%d", argID))
		args = append(args, input.StartTime)
		argID++
	}

	if input.EndTime != "" {
		setValues = append(setValues, fmt.Sprintf("end_time <= $%d", argID))
		args = append(args, input.EndTime)
		argID++
	}

	if input.ID != 0 {
		setValues = append(setValues, fmt.Sprintf("id = $%d", argID))
		args = append(args, input.ID)
		argID++
	}

	queryArgs := strings.Join(setValues, " AND ")

	var query string
	if len(setValues) > 0 {
		query = fmt.Sprintf("SELECT id, start_time, end_time FROM times WHERE %s", queryArgs)
	} else {
		query = "SELECT id, start_time, end_time FROM times"
	}

	// Pagination
	paginationQuery := fmt.Sprintf("SELECT COUNT(*) FROM (%s) AS count_query", query)
	pagination, offset, err := r.getPagination(paginationQuery, input.Limit, input.Page)
	if err != nil {
		return models.TimesAndPagination{}, err
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", input.Limit, offset)

	var times []models.Time
	err = r.db.Select(&times, query, args...)
	if err != nil {
		return models.TimesAndPagination{}, err
	}

	return models.TimesAndPagination{
		Times:      times,
		Pagination: pagination,
	}, nil
}

// Helper: Pagination logic
func (r *TimeRepository) getPagination(query string, limit, page int) (models.Pagination, int, error) {
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
