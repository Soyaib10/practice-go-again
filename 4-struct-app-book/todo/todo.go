package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Note is a struct
type Todo struct {
	Text string `json:"text"`
}

// Display method shows the content of Todo struct
func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

// Save mehtod saves string in the local system
func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

// New is the constructor for the struct Todo with validation
func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("invalid input")
	}
	return Todo{
		Text: text,
	}, nil
}