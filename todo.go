package main

import (
	"flag"
	"fmt"
	"os"
)

func addNewTask(args []string) (int, error) {
	if len(args) != 2 {
		return 0, fmt.Errorf("Incomplete arguments provided")
	}
	err := AddTaskToFile(args[1])
	fmt.Printf("Adding '%s' task.\n", args[1])
	return 0, err
}

func listTasks() (int, error) {
	err := ListAllTasks()
	return 0, err
}

func doneTask(args []string) (int, error) {
	if len(args) != 2 {
		return 0, fmt.Errorf("Incomplete arguments provided")
	}
	MarkTaskAsDone(args[1])
	fmt.Printf("Task '%s' marked as done.\n ", args[1])
	return 0, nil
}

func clearTasks() (int, error) {
	err := ClearTaskFile()
	fmt.Printf("Task file was reset.\n")
	return 0, err
}

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
