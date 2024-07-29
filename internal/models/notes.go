package models

import (
	"database/sql"
	"errors"
	"fmt"
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
	// Write the SQL statement we want to execute. Again, I've split it over two
	// lines for readability.
	stmt := `SELECT id, title, content, created, updated FROM notes
	WHERE id = ?`

	// Use the QueryRow() method on the connection pool to execute our
	// SQL statement, passing in the untrusted id variable as the value for the
	// placeholder parameter. This returns a pointer to a sql.Row object which
	// holds the result from the database.
	row := m.DB.QueryRow(stmt, id)

	// Initialize a pointer to a new zeroed Snippet struct.
	n := &Note{}

	// Use row.Scan() to copy the values from each field in sql.Row to the
	// corresponding field in the Snippet struct. Notice that the arguments
	// to row.Scan are *pointers* to the place you want to copy the data into,
	// and the number of arguments must be exactly the same as the number of
	// columns returned by your statement.
	err := row.Scan(&n.ID, &n.Title, &n.Content, &n.Created, &n.Updated)
	if err != nil {
		// If the query returns no rows, then row.Scan() will return a
		// sql.ErrNoRows error. We use the errors.Is() function check for that
		// error specifically, and return our own ErrNoRecord error
		// instead (we'll create this in a moment).
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	fmt.Println("GET")

	// If everything went OK then return the Snippet object.
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
