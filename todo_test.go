package main

import (
	"testing"
)

func TestCommandLineParser(t *testing.T) {
	var s []string
	_, err := parseArgs(s)
	if err != nil {
		t.Failed()
	}
}

func TestCommandLineParserWithAdd(t *testing.T) {
	s := []string{"add", "somestring"}
	_, err := parseArgs(s)
	if err != nil {
		t.Fatal("Error received")
	}
}

func TestCommandLineParserWithList(t *testing.T) {
	s := []string{"list", "somestring"}
	_, err := parseArgs(s)
	if err != nil {
		t.Fatal("Error received")
	}
}

func TestCommandLineParserWithDone(t *testing.T) {
	s := []string{"done", "somestring"}
	_, err := parseArgs(s)
	if err == nil {
		t.Fatal("Error received")
	}
}

func TestCommandLineParserWithClear(t *testing.T) {
	s := []string{"clear", "somestring"}
	_, err := parseArgs(s)
	if err != nil {
		t.Fatal("Error received")
	}
}

func TestCommandLineParserWithInvalid(t *testing.T) {
	s := []string{"invalid", "somestring"}
	_, err := parseArgs(s)
	if err == nil {
		t.Fatal("Incorrect action received")
	}
}

func TestAddTaskIncompleteArgs(t *testing.T) {
	s := []string{"add"}
	_, err := addNewTask(s)
	if err == nil {
		t.Fatal("Error not received with incomplete arguments")
	}
}

func TestAddTaskIncompleteLongArgs(t *testing.T) {
	s := []string{"add", "Some text", "another text"}
	_, err := addNewTask(s)
	if err == nil {
		t.Fatal("Error not received with incomplete arguments")
	}
}

func TestAddTask(t *testing.T) {
	s := []string{"add", "This is a new task"}
	_, err := addNewTask(s)
	if err != nil {
		t.Fatal("Error not expected received")
	}
}

func TestDoneTaskIncompleteArgs(t *testing.T) {
	s := []string{"done"}
	_, err := doneTask(s)
	if err == nil {
		t.Fatal("Error not received with incomplete arguments")
	}
}

func TestDoneTask(t *testing.T) {
	s := []string{"done", "1"}
	_, err := doneTask(s)
	if err != nil {
		t.Fatal("Error not expected received")
	}
}

func TestUsage(t *testing.T) {
	usage()
}
