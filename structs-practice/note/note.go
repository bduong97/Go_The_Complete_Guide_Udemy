package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note) //struct fields need to be capitalized in order to be marshalled by external package
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.Title, note.Content)
}
func New(title, content string) (Note, error) {
	if content == "" || title == "" {
		return Note{}, errors.New("invalid input")

	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
