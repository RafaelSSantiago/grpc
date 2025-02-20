package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name, description, CategoryId string) (*Course, error) {
	id := uuid.New().String()

	_, err := c.db.Exec(`INSERT INTO courses (id,name,description,category_id) VALUES($1, $2, $3) RETURNING id`,
		id, name, description, CategoryId)
	if err != nil {
		return nil, err
	}
	return &Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  CategoryId,
	}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM courses")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	cousers := []Course{}
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		cousers = append(cousers, Course{ID: id, Name: name, Description: description})
	}

	return cousers, nil
}
