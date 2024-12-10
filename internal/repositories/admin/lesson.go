package admin

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"student-management/internal/models"
)

type LessonRepository struct {
	db *sqlx.DB
}

func NewLessonRepository(db *sqlx.DB) *LessonRepository {
	return &LessonRepository{
		db: db,
	}
}

// AddLesson adds a new lesson to the database
func (r *LessonRepository) AddLesson(lesson models.Lesson) (models.Lesson, error) {
	query := `INSERT INTO lessons(name, code) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(query, lesson.Name, lesson.Code).Scan(&lesson.ID)
	return lesson, err
}

// UpdateLesson updates a lesson's name and code
func (r *LessonRepository) UpdateLesson(lesson models.Lesson) error {
	query := `UPDATE lessons SET name=$1, code=$2 WHERE id=$3`
	_, err := r.db.Exec(query, lesson.Name, lesson.Code, lesson.ID)
	return err
}

// DeleteLesson deletes a lesson by ID
func (r *LessonRepository) DeleteLesson(id int) error {
	query := `DELETE FROM lessons WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}

// GetLessonByID fetches a lesson by its ID
func (r *LessonRepository) GetLessonByID(id int) (models.Lesson, error) {
	var lesson models.Lesson
	query := `SELECT id, name, code FROM lessons WHERE id=$1`
	err := r.db.Get(&lesson, query, id)
	return lesson, err
}

// GetLessons retrieves lessons with optional search and pagination
func (r *LessonRepository) GetLessons(input models.LessonSearch) (models.LessonsAndPagination, error) {
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

	if input.Code != "" {
		setValues = append(setValues, fmt.Sprintf("code ILIKE '%%%s%%'", input.Code))
	}

	queryArgs := strings.Join(setValues, " AND ")

	var query string
	if len(setValues) > 0 {
		query = fmt.Sprintf("SELECT id, name, code FROM lessons WHERE %s", queryArgs)
	} else {
		query = "SELECT id, name, code FROM lessons"
	}

	// Pagination
	paginationQuery := fmt.Sprintf("SELECT COUNT(*) FROM (%s) AS count_query", query)
	pagination, offset, err := r.getPagination(paginationQuery, input.Limit, input.Page)
	if err != nil {
		return models.LessonsAndPagination{}, err
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", input.Limit, offset)

	var lessons []models.Lesson
	err = r.db.Select(&lessons, query, args...)
	if err != nil {
		return models.LessonsAndPagination{}, err
	}

	return models.LessonsAndPagination{
		Lessons:    lessons,
		Pagination: pagination,
	}, nil
}

// Helper: Pagination logic
func (r *LessonRepository) getPagination(query string, limit, page int) (models.Pagination, int, error) {
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
