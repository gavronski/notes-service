package main

import (
	"log"
	"net/http"
	"notes-service/data"
	"time"
)

type ReqPayload struct {
	Action string `json:"action"`
	NoteID string `json:"note_id,omitempty"`
	Data   Note   `json:"data,omitempty"`
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

func (app *Config) GetNotesList(w http.ResponseWriter, r *http.Request) {
	var notes []*data.Note

	notes, err := app.Models.Note.GetAllNotes()

	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error: false,
		Data:  notes,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetNoteByID(w http.ResponseWriter, r *http.Request) {
	var requestPayload ReqPayload
	err := app.readJSON(w, r, &requestPayload)

	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	id := requestPayload.Data.ID
	note, err := app.Models.Note.GetNoteByID(id)

	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error: false,
		Data:  note,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) PostNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload ReqPayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	var note data.Note = data.Note(requestPayload.Data)

	err = note.InsertNote()

	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Note's been successfully added.",
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) UpdateNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload ReqPayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var note data.Note = data.Note(requestPayload.Data)
	log.Println(note)
	err = note.UpdateNote()
	log.Println(err)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Note's been successfully updated.",
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) DeleteNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload ReqPayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	var note data.Note = data.Note(requestPayload.Data)

	err = note.DeleteNote()

	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Note's been successfully deleted.",
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
