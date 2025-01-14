package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}
func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todoText := getUserInput("Todo text: ")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = saveData(todo)
	if err != nil {
		return
	}

	userNote.Display()
	err = saveData(userNote)
	if err != nil {
		return
	}
	
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("saving the note failed")
		return err
	}
	fmt.Println("Saving the note succeeded!")
	return nil
}
func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}
func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)  //reader that listens on the command line
	text, err := reader.ReadString('\n') //need single quotes
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") //windows specific
	return text
}
