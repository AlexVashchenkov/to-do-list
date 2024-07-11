package models

import (
	"errors"
	"fmt"
)

type ToDo struct {
	Name        string
	Description string
	isActive    bool
}

func NewToDo(name, description string, isActive bool) ToDo {
	return ToDo{Name: name, Description: description, isActive: isActive}
}

func (t *ToDo) IsActive() bool {
	return t.isActive
}

func (t *ToDo) Status() string {
	if t.isActive {
		return "Active"
	} else {
		return "Completed"
	}
}

func (t *ToDo) ToPrint() string {
	if t.Description != "" {
		return "\"" + t.Name + "\": " + t.Description + ", status: " + t.Status()
	}
	return "\"" + t.Name + "\", status: " + t.Status()
}

func (t *ToDo) ToString() string {
	return "\"" + t.Name + "\": " + t.Description + ", status: " + t.Status()
}

func (t *ToDo) Tick() error {
	if t.isActive {
		t.isActive = false
		fmt.Print("\"" + t.Name + "\" has been ticked\n\n")
		return nil
	}
	return errors.New("todo is already completed")
}
