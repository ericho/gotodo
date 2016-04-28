package main

import (
	"fmt"
)

func InitTaskFile() error {
	return nil
}

func AddTaskToFile(task string) error {
	if len(task) == 0 {
		return fmt.Errorf("Empty tasks provided.")
	}
	fmt.Printf("Adding the task...")
	return nil
}

func MarkTaskAsDone(id int) error {
	return nil
}

func ListAllTasks() error {
	return nil
}
