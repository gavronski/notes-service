package main

import (
	"net/http"
	"notes-service/data"
)

type ReqPayload struct {
	Action string `json:"action"`
	NoteID int    `json:"noteID,omitempty"`
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
