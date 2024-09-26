package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// all keys are exported because json.Marshal funtion
// do not convert un-exported fields for safety
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title string, content string) (Note, error) {

	// validate arguments and return error
	if title == "" || content == "" {
		return Note{}, errors.New("title or content cannot be empty")
	}

	// return the Note instance & nil as error for success
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

// Display the note on console
func (n Note) Display() {
	fmt.Printf("\nThe title of your note is: \n%v\nThe content of your note is: \n%v\n", n.Title, n.Content)
}

// creates a json file and saves the note there,
// title of the note will be "Note.Title"
func (n Note) Save() error {

	// struct -> marshall
	jsonByte, err := json.Marshal(n)
	if err != nil {
		return err
	}

	// Note.Title will be used as the fileName
	// sanitize & correct the "Note.Title" before saving
	// replace all spaces in titles with "_"
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	// lowercase the fileName, add .json file extension to the end
	fileName = strings.ToLower(fileName) + ".json"

	// create | save a new json file with "Note.Title" field value
	return os.WriteFile(fileName, jsonByte, 0644)
}
