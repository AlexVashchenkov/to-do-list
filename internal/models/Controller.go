package models

import (
	"fmt"
)

type Controller struct {
	Storage Storage
}

func (controller *Controller) AddToDo(name, desc string) error {
	if desc == "" {
		controller.Storage.Add(ToDo{Name: name})
	} else {
		controller.Storage.Add(ToDo{Name: name, Description: desc})
	}
	fmt.Print("ToDo has been added successfully\n\n")
	return nil
}

func (controller *Controller) ClearAllToDos() error {
	controller.Storage.ClearAllToDos()
	return nil
}

func (controller *Controller) ClearCompletedToDos() error {
	controller.Storage.ClearCompletedToDos()
	return nil
}

func (controller *Controller) DeleteByNumber(number int) error {
	return controller.Storage.DeleteByNumber(number)
}

func (controller *Controller) PrintToDos() error {
	err1 := controller.PrintActiveToDos()
	if err1 != nil {
		return err1
	}
	err2 := controller.PrintCompletedToDos()
	if err2 != nil {
		return err2
	}
	return nil
}

func (controller *Controller) PrintActiveToDos() error {
	if controller.Storage.numberOfActiveTasks > 0 {
		fmt.Print("Active to-do's:\n")
		controller.Storage.PrintActiveToDos()
	} else {
		fmt.Print("No active to-do's yet\n\n")
	}
	return nil
}

func (controller *Controller) PrintCompletedToDos() error {
	if len(controller.Storage.listOfToDos)-controller.Storage.numberOfActiveTasks > 0 {
		fmt.Print("Completed to-do's:\n")
		controller.Storage.PrintCompletedToDos()
	} else {
		fmt.Print("No completed to-do's yet\n\n")
	}
	return nil
}

func (controller *Controller) Tick(number int) error {
	return controller.Storage.Tick(number)
}

func (controller *Controller) ViewToDo(number int) error {
	toDo, err := controller.Storage.FindByNumber(number)
	if err != nil {
		return err
	}
	fmt.Print(toDo.ToString(), "\n\n")
	return nil
}
