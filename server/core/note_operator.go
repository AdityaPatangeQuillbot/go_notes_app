package core

import (
	"log"
	"main/db"
	"main/models"
)

type NotesOperator interface {
	// GetNoteByID returns a note by its ID
	GetNoteByID(id int) (*models.Note, error)
	// GetNotes returns all notes
	GetNotes() ([]*models.Note, error)
	// CreateNote creates a new note
	CreateNote(note *models.Note) error
	// UpdateNote updates a note
	UpdateNote(note *models.Note) error
	// DeleteNote deletes a note by its ID
	DeleteNote(id int) error
}

type AppNotesOperator struct {
	id     string
	client db.PrismaClient
}

func NewNotesOperator() NotesOperator {
	client := db.NewClient()

	err := client.Prisma.Connect()

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &AppNotesOperator{
		id:     "AppNotesOperator",
		client: *client,
	}
}

func (n *AppNotesOperator) GetNoteByID(id int) (*models.Note, error) {

	return nil, nil
}

func (n *AppNotesOperator) GetNotes() ([]*models.Note, error) {
	return nil, nil
}

func (n *AppNotesOperator) CreateNote(note *models.Note) error {

	return nil
}

func (n *AppNotesOperator) UpdateNote(note *models.Note) error {

	return nil
}

func (n *AppNotesOperator) DeleteNote(id int) error {

	return nil
}
