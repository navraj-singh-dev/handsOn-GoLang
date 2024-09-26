package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Text string `json:"text"`
}

func NewTodo(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("title or content cannot be empty")
	}
	return Todo{
		Text: text,
	}, nil
}

func (t Todo) Display() {
	fmt.Printf("\nThe text of your note is: \n%v\n", t.Text)
}

func (t Todo) Save() error {
	jsonByte, err := json.Marshal(t)
	if err != nil {
		return err
	}
	fileName := strings.ReplaceAll(t.Text, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	return os.WriteFile(fileName, jsonByte, 0644)
}
