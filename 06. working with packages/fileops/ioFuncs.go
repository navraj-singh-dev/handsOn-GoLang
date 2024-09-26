package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFileFloat(filename string) (float64, error) {
	
	// file -> []byte
	contentInByte, err := os.ReadFile(filename)
	if err != nil {
		return 0, errors.New("either file not present or cannot read it")
	}

	// []byte -> string
	contentInString := string(contentInByte)

	// string -> float64
	contentInFloat, err := strconv.ParseFloat(contentInString, 64)
	if err != nil {
		return 0, errors.New("cannot parse the file content to float")
	}

	return contentInFloat, nil
}

func WriteFileFloat(filename string, inputForFile float64) {

	// input -> string
	var inputInString string = fmt.Sprint(inputForFile)

	// string -> []byte
	var inputInByte []byte = []byte(inputInString)

	// []byte -> file
	os.WriteFile(filename, inputInByte, 0644)
}
