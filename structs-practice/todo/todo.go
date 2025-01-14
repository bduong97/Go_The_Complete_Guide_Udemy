package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo) //struct fields need to be capitalized in order to be marshalled by external package
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
func (todo Todo) Display() {
	fmt.Printf(todo.Text)
}
func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid input")

	}

	return Todo{
		Text: content,
	}, nil
}
