package main

import (
	"errors"
	"fmt"
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

// empty interface means any data type as argument is accepted as argument
func printAnything(i interface{}) {
	fmt.Println(i) // println also accept every data-type as argument
}

func main() {
	noteInstance, _ := New("I am title", "I am content")

	// now this function below will print (errors might occurs for some scenarios) 
	// any custom struct or any data type you give as argument
	// I gave the "note" struct as argument and it successfully printed it
	printAnything(noteInstance)
}

// There are many problems that occurs with empty interface
// so those problems are discussed in the .txt file of this 
// folder.
// So i have explained how empty interface can be used standlone
// in the next folder i will include the switch statement to solve the
// problems it brings.
