package data

import (
	"context"
	"database/sql"
	"time"
)

var db *sql.DB

const dbTimeout = time.Second * 3

type Models struct {
	Note Note
}

type Note struct {
	ID              int       `json:"note_id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	TextColor       string    `json:"text_color"`
	BackgroundColor string    `json:"background_color"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func NewModels(db *sql.DB) Models {
	db = db

	return Models{
		Note: Note{},
	}
}

func (note *Note) GetAllNotes() ([]*Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from notes;`

	var notes []*Note
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return notes, err
	}

	// prepare row for scaning to destination
	for rows.Next() {
		var note Note

		err := rows.Scan(
			&note.ID,
			&note.Name,
			&note.Description,
			&note.TextColor,
			&note.BackgroundColor,
			&note.CreatedAt,
			&note.UpdatedAt,
		)

		if err != nil {
			return notes, nil
		}

		notes = append(notes, &note)
	}

	return notes, nil
}
