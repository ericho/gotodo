package main

import (
	"flag"
	"fmt"
	"os"
)

func parseArgs(args []string) (action int, err error) {
	if args == nil || flag.NArg() == 0 {
		return 0, nil
	}

	switch args[0] {
	case "add":
		return addNewTask(args)
	case "list":
		return listTasks()
	case "done":
		return doneTask(args)
	case "clear":
		return clearTasks()
	default:
		return 0, fmt.Errorf("Unknown option provided: %s", args[0])
	}
}

func addNewTask(args []string) (int, error) {
	if len(args) != 2 {
		return 0, fmt.Errorf("Incomplete arguments provided")
	}
	err := AddTaskToFile(args[1])
	fmt.Printf("Adding '%s' task. Not implemented yet\n", args[1])
	// Receive one parameter, a quoted string that will be added
	// into the file .gotodo.txt
	// Will read the ID and assign a new one for this task
	return 0, err
}

func listTasks() (int, error) {
	fmt.Printf("Listing all tasks. Not implemented yet\n")
	// Will read the file content and print the list of tasks
	return 0, nil
}

func doneTask(args []string) (int, error) {
	if len(args) != 2 {
		return 0, fmt.Errorf("Incomplete arguments provided")
	}
	fmt.Printf("Task '%s' marked as done. Not implemented yet.\n ", args[1])
	// Receive an id number as a parameter, this tasks will be put
	// done with the [x] mark
	return 0, nil
}

func clearTasks() (int, error) {
	fmt.Printf("Clearing all the tasks. Not implemented yet.\n")
	// Delete all the file content and create a new header.
	return 0, nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [add] <task> | [list] | [done] <task id> | [clear]\n", os.Args[0])
	os.Exit(1)
}

func main() {

	flag.Parse()

	if flag.NArg() < 1 {
		usage()
	}

	args := flag.Args()

	InitTaskFile()
	_, err := parseArgs(args)
	if err != nil {
		fmt.Println(err)
		usage()
		os.Exit(1)
	}

}
