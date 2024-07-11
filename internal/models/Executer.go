package models

import (
	"errors"
	"strconv"
	"strings"
)

type Executer struct {
	controller *Controller
}

func (controller *Controller) Parse(str string) (bool, error) {
	str = strings.ToLower(str)

	if str == "" {
		return false, nil
	}

	args := strings.Fields(str)
	command := args[0]
	switch {
	case command == "list" && len(args) >= 1 && len(args) <= 2:
		return true, nil

	case command == "add":
		return true, nil

	case command == "clear":
		if len(args) == 1 {
			return true, controller.ClearAllToDos()
		} else {
			argument := args[1]
			if argument == "complete" {
				return true, controller.ClearCompletedToDos()
			} else {
				return false, errors.New("invalid number of arguments")
			}
		}

	case command == "tick":
		if len(args) != 2 {
			return false, errors.New("invalid number of arguments")
		} else {
			argument, _ := strconv.Atoi(args[1])
			err := controller.Tick(argument)
			if err != nil {
				return false, err
			}
		}

	case command == "delete":
		if len(args) != 2 {
			return false, errors.New("invalid number of arguments")
		} else {
			argument, _ := strconv.Atoi(args[1])
			err := controller.DeleteByNumber(argument)
			if err != nil {
				return false, err
			}
			return true, nil
		}
	case command == "exit":
		return true, nil
	case command == "view":
		if len(args) == 2 {
			argument, _ := strconv.Atoi(args[1])
			err := controller.ViewToDo(argument)
			if err != nil {
				return false, err
			}
		} else {
			return false, errors.New("invalid number of arguments")
		}
	default:
		return false, errors.New("invalid command")
	}
	return true, nil
}
