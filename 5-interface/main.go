package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"struct-app/note"
	"struct-app/todo"
)

// Saver interface connects save methods
type saver interface {
	Save() error
}

// embaded struct means any value of outputtable type must have all the methods of saver interface and in addition it must have a display method.
type outputtable interface { 
	saver
	Display()
}

func main() {
	printSomething("------WELCOME TO MY APP------")
	title, content := getNoteData()         // retlated to note.go
	todoText := getUserInput("Todo text: ") // related to todo.go

	// create todo and userNote
	todo, err := todo.New(todoText) // create
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content) // create
	if err != nil {
		fmt.Println(err)
		return
	}

	// display and save todo and note.
	err = outputData(todo)
	if err != nil {
		return
	}
	outputData(userNote)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving note failed!")
		return err
	}
	fmt.Println("Saving the note succeeded!")
	return nil
}

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


// pass any value truough a function
func printSomething(value interface{}) {
	fmt.Println(value)

	// check type of value
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer: ", intVal)
		return
	}
}