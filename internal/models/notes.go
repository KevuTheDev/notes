package models

import (
	"database/sql"
	"errors"
	"time"
)

type Note struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Updated time.Time
}

type NoteModel struct {
	DB *sql.DB
}

type NoteModelInterface interface {
	Get(id int) (*Note, error)
	Latest() ([]*Note, error)
}

func (m *NoteModel) Get(id int) (*Note, error) {
	stmt := `SELECT id, title, content, created, updated FROM notes
	WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	n := &Note{}

	err := row.Scan(&n.ID, &n.Title, &n.Content, &n.Created, &n.Updated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return n, nil
}

func (m *NoteModel) Latest() ([]*Note, error) {
	stmt := `SELECT id, title, content, created, updated FROM notes
	ORDER BY updated DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notes := []*Note{}
	for rows.Next() {
		n := &Note{}

		err = rows.Scan(&n.ID, &n.Title, &n.Content, &n.Created, &n.Updated)
		if err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil

}
