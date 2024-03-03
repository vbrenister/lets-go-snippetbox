package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

type SnippetModelInterface interface {
	Insert(title string, content string, expires int) (int, error)
	Get(id int) (Snippet, error)
	Latest() ([]Snippet, error)
}

type Snippet struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	ExpiresAt time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires) 
	VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + $3::interval) RETURNING id`

	expiresInterval := fmt.Sprintf("%d days", expires)

	var id int
	err := m.DB.QueryRow(stmt, title, content, expiresInterval).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	query := `SELECT id, title, content, created, expires 
	FROM snippets WHERE expires > CURRENT_TIMESTAMP AND id = $1`

	var s Snippet

	err := m.DB.QueryRow(query, id).Scan(&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.ExpiresAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		}

		return Snippet{}, err
	}

	s.Content = strings.Replace(s.Content, "\\n", "\n", -1)

	return s, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > CURRENT_TIMESTAMP ORDER BY id DESC LIMIT 10`

	var snippets []Snippet

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var s Snippet
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.ExpiresAt)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
