package note

import (
	"errors"
	"fmt"
)

type Saver interface {
	Save() error
}

type OutputAndSave interface {
	Display()
	Saver
}

func SaveInFile(s Saver) error {
	err := s.Save()
	if err != nil {
		fmt.Println(errors.New("error during saving of file"))
		return err
	}
	return nil
}

func DisplayAndSave(o OutputAndSave) error {
	o.Display()
	err := o.Save()
	if err != nil {
		fmt.Println(errors.New("error during saving of file"))
		return err
	}
	return nil
}
