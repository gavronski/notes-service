package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"notes-service/data"
	"notes-service/notes"

	"google.golang.org/grpc"
)

type NoteServer struct {
	notes.UnimplementedNoteServiceServer
	Models data.Models
}

func (n *NoteServer) GetNotes(ctx context.Context, req *notes.NoteRequest) (*notes.NoteResponse, error) {
	input := req.GetPayload()
	note, err := n.Models.Note.GetNoteByID(int(input.Id))

	var noteData *notes.Note = &notes.Note{
		Id:              int32(note.ID),
		Name:            note.Name,
		Description:     note.Description,
		TextColor:       note.TextColor,
		BackgroundColor: note.BackgroundColor,
	}

	if err != nil {
		res := &notes.NoteResponse{
			Note: noteData,
		}
		return res, err
	}

	return &notes.NoteResponse{Note: noteData}, nil
}

func (app *Config) gRPCListen() {

	//  set up TCP network listener for a gRPC server.
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))

	if err != nil {
		log.Fatal("Faile to listen for gRPC: %v", err)
	}

	// create gRPC server
	srv := grpc.NewServer()

	// register service
	notes.RegisterNoteServiceServer(srv, &NoteServer{Models: app.Models})
	log.Printf("gRPC server started on port: %s", gRPCPort)

	// start grpc Server
	if err = srv.Serve(listen); err != nil {
		log.Fatal("Faile to listen for gRPC: %v", err)
	}
}
