package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"struct-app/note"
	"struct-app/todo"
)

// getUserInput takes input for title and content
func getUserInput(promt string) string {
	fmt.Print(promt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

// getNoteData contains title and content
func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func main() {
	title, content := getNoteData()         // retlated to note.go
	todoText := getUserInput("Todo text: ") // related to todo.go

	// todo section
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = todo.Save()
	if err != nil {
		fmt.Println("Saving todo failed!")
		return
	}
	fmt.Println("Saving the todo succeeded!")

	// note section
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving note failed!")
		return
	}
	fmt.Println("Saving the note succeeded!")
}
