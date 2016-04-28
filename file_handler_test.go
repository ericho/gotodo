package main

import "testing"

func TestAddTaskToFile(t *testing.T) {
	s := ""
	err := AddTaskToFile(s)
	if err == nil {
		t.Fatal("Error expected not found")
	}
}
