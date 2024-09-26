package main

import (
	"fmt"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {

	// create a custom value
	note := Note{
		Title:     "title",
		Content:   "Content",
		CreatedAt: time.Now(),
	}

	// this will work because b, c arguments are both int
	// otherwise the + operator would have caused issues
	// if b and c were different data types. Eventhough this
	// function let b, c accept different data types
	fmt.Println(add_K_And_L(note, 1000, 1000))
}

// now this generic function do not need empty interface at all
// jsut define your accepted data-types with arguments in the []
// syntax
func add_K_And_L[J any, K int | float64 | string](a J, b K, c K) (J, K) {
	// a can be any data type
	// b can be int or string or float
	// c can be int or string or float
	// Now i can also return these placeholders to tell go
	// that my return types are also not hard coded like int, float etc. instead
	// my return types can be whatever data-type that my placeholders J, K represent

	// One crucial thing to remember is that operators work differently with different
	// data-types. Like here + operator is only gonna work if b and c are same data types.
	// If i call this function and then pass string in "b" and int in "c" then this function
	// will fail because + cannot add a string with int. So make sure to use "ok" syntax and
	// other coding logic when working with generics to make sure no silly errors occurs.
	// generics are not magic they are just simple way to make a function accept flexible
	// amount of data-types dynamically for a standalone argument or arguments. But handling
	// arguments and the logic is in your hands.

	return a, (b + c)
}
