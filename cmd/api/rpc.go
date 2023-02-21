package main

import (
	"encoding/json"
	"notes-service/data"
)

type RPCServer struct{}

type RPCPayload struct {
	Action string
	ID     int
}

// ReadAction is RPCServer method that is invoked by RPC client from the broker-service.
func (r *RPCServer) ReadAction(payload RPCPayload, resp *string) error {
	switch payload.Action {
	case "get-notes-list":
		response := r.GetNotesList()
		jsonData, _ := json.Marshal(response)

		*resp = string(jsonData)

	case "get-note-by-id":
		response := r.GetNoteByID(payload.ID)
		jsonData, _ := json.Marshal(response)

		*resp = string(jsonData)
	}

	return nil
}

func (r *RPCServer) GetNotesList() jsonResponse {
	var notes []*data.Note
	notes, err := app.Models.Note.GetAllNotes()

	response := jsonResponse{
		Error: false,
		Data:  notes,
	}

	if err != nil {
		response.Error = true
		return response
	}

	return response
}

func (r *RPCServer) GetNoteByID(id int) jsonResponse {
	note, err := app.Models.Note.GetNoteByID(id)
	response := jsonResponse{
		Error: false,
		Data:  note,
	}

	if err != nil {
		response.Error = true
		return response
	}

	return response
}
