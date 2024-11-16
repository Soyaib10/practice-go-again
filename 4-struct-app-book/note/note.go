package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Note is a struct
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// New is the constructor for the struct Note with validation
func New(title, content string) (Note, error) {
	// validation
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	// return struct
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

// Display method shows the content of Note struct
func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content: \n%v\n", note.Title, note.Content)
}

// Save mehtod saves string in the local system
func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
