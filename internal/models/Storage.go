package models

import (
	"errors"
	"fmt"
	"strconv"
)

type Storage struct {
	listOfToDos         []ToDo
	numberOfActiveTasks int
}

func NewStorage(todos []ToDo) Storage {
	return Storage{listOfToDos: todos, numberOfActiveTasks: 0}
}

func (storage *Storage) ListOfToDos() []ToDo {
	return storage.listOfToDos
}

func (storage *Storage) Add(do ToDo) {
	do.isActive = true
	storage.listOfToDos = append(storage.listOfToDos, do)
	storage.numberOfActiveTasks++
}

func (storage *Storage) FindByNumber(number int) (*ToDo, error) {
	if len(storage.listOfToDos) == 0 {
		return nil, errors.New("no to-dos found")
	}
	if number <= 0 || number > len(storage.listOfToDos) {
		return nil, errors.New("invalid to-do number")
	}
	return &storage.listOfToDos[number-1], nil
}

func (storage *Storage) DeleteByNumber(number int) error {
	if len(storage.listOfToDos) == 0 {
		return errors.New("no to-dos found")
	}
	if number <= 0 || number > len(storage.listOfToDos) {
		return errors.New("invalid to-do number")
	}
	storage.listOfToDos = append(storage.listOfToDos[:number-1], storage.listOfToDos[number:]...)
	return nil
}

func (storage *Storage) PrintActiveToDos() {
	for idx, elem := range storage.listOfToDos {
		if elem.IsActive() {
			fmt.Println(strconv.Itoa(idx+1)+".", elem.ToPrint())
		}
	}
	fmt.Print("\n")
}

func (storage *Storage) PrintCompletedToDos() {
	for idx, elem := range storage.listOfToDos {
		if !elem.IsActive() {
			fmt.Println(strconv.Itoa(idx+1)+".", elem.ToPrint())
		}
	}
	fmt.Print("\n")
}

func (storage *Storage) ClearAllToDos() {
	storage.listOfToDos = []ToDo{}
}

func (storage *Storage) ClearCompletedToDos() {
	for idx, toDo := range storage.listOfToDos {
		if !toDo.isActive {
			storage.listOfToDos = append(
				storage.listOfToDos[0:idx],
				storage.listOfToDos[idx+1:]...,
			)
		}
	}
}

func (storage *Storage) Tick(number int) error {
	toDo, err := storage.FindByNumber(number)
	if err != nil {
		return err
	}
	err = toDo.Tick()
	if err != nil {
		return err
	}
	storage.numberOfActiveTasks--
	return nil
}
