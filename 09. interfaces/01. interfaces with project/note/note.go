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

func New(title string, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("title or content cannot be empty")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) Display() {
	fmt.Printf("\nThe title of your note is: \n%v\nThe content of your note is: \n%v\n", n.Title, n.Content)
}

func (n Note) Save() error {

	jsonByte, err := json.Marshal(n)
	if err != nil {
		return err
	}

	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	return os.WriteFile(fileName, jsonByte, 0644)
}
