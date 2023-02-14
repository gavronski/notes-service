package data

import (
	"context"
	"database/sql"
	"log"
	"time"
)

var newDB *sql.DB

const dbTimeout = time.Second * 3

type Models struct {
	Note Note
}

type Note struct {
	ID              int       `json:"note_id,string,omitempty"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	TextColor       string    `json:"text_color"`
	BackgroundColor string    `json:"background_color"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func NewModels(db *sql.DB) Models {
	newDB = db

	return Models{
		Note: Note{},
	}
}

// GetAllNotes return list of notes
func (n *Note) GetAllNotes() ([]*Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from notes;`

	var notes []*Note
	rows, err := newDB.QueryContext(ctx, query)

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
			return notes, err
		}

		notes = append(notes, &note)
	}

	return notes, nil
}

// GetNoteByID returs single note
func (n *Note) GetNoteByID(id int) (*Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from notes where id = $1;`

	row := newDB.QueryRowContext(ctx, query, id)

	var note Note
	err := row.Scan(
		&note.ID,
		&note.Name,
		&note.Description,
		&note.TextColor,
		&note.BackgroundColor,
		&note.CreatedAt,
		&note.UpdatedAt,
	)

	if err != nil {
		return &note, err
	}

	return &note, nil
}

func (note *Note) InsertNote() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	insert into 
		notes(id, name, description, text_color, background_color, created_at, updated_at)
	values
		(nextval('notes_sequence'), $1, $2, $3, $4, $5, $6);`

	_, err := newDB.ExecContext(ctx, query,
		note.Name,
		note.Description,
		note.TextColor,
		note.BackgroundColor,
		time.Now(),
		time.Now())

	if err != nil {
		return err
	}
	return nil
}

func (note *Note) UpdateNote() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	log.Println(note)
	query := `
	update notes 
		set name = $1, description = $2, text_color = $3, background_color = $4, updated_at = $5 
	where id = $6;`
	log.Println(note)
	_, err := newDB.ExecContext(ctx, query,
		note.Name,
		note.Description,
		note.TextColor,
		note.BackgroundColor,
		time.Now(),
		note.ID,
	)

	if err != nil {
		return err
	}
	return nil
}

func (note *Note) DeleteNote() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `delete from notes where id = $1;`

	_, err := newDB.ExecContext(ctx, query, note.ID)

	if err != nil {
		return err
	}
	return nil
}
