package main

import (
	//	"flag"
	"errors"
	"fmt"
	"os"
)

const (
	ADD = 1 << iota
	LIST
	DONE
	CLEAR
)

func readactionfromargs(arg string) (action int, err error) {
	switch {
	case "add" == arg:
		return ADD, nil
	case "list" == arg:
		return LIST, nil
	case "done" == arg:
		return DONE, nil
	case "clear" == arg:
		return CLEAR, nil
	}
	return 0, errors.New("Undefined option received.")
}

func readargs(args []string) (action int, err error) {

	if args == nil || len(args) == 0 {
		return 0, nil
	}

	action, err = readactionfromargs(args[0])
	return action, err
}

func executeoption(action int, args []string) (int, error) {
	switch {
	case action == ADD:
		return add_task(args)
	case action == LIST:
		return list_tasks()
	case action == DONE:
		return done_task(args)
	case action == CLEAR:
		return clear_tasks()
	}
	return 1, errors.New("Invalid option provided")
}

func add_task(args []string) (int, error) {
	// Receive one parameter, a quoted string that will be added
	// into the file .gotodo.txt
	// Will read the ID and assign a new one for this task
	return 0, nil
}

func list_tasks() (int, error) {
	// Will read the file content and print the list of tasks
	return 0, nil
}

func done_task(args []string) (int, error) {
	// Receive an id number as a parameter, this tasks will be put
	// done with the [x] mark
	return 0, nil
}

func clear_tasks() (int, error) {
	// Delete all the file content and create a new header.
	return 0, nil
}

func main() {
	fmt.Println(len(os.Args))
	fmt.Println(os.Args[1:])
	action, err := readargs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}
	result, err := executeoption(action, os.Args[2:])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Lets print something")
	fmt.Println(action)
	fmt.Println(result)
}
