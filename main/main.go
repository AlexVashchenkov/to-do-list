package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"to-do-list/internal/models"
)

func main() {
	var controller models.Controller

	for str, _ := bufio.NewReader(os.Stdin).ReadString('\n'); ; str, _ = bufio.NewReader(os.Stdin).ReadString('\n') {
		str := strings.Trim(strings.ToLower(str), "\n")
		if str == "exit" {
			break
		}

		nice, err := controller.Parse(str)
		if err != nil {
			fmt.Print(nice, err)
		}
	}
}
