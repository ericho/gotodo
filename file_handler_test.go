package main

import "testing"

func TestAddTaskToFile(t *testing.T) {
	s := ""
	err := AddTaskToFile(s)
	if err == nil {
		t.Fatal("Error expected not found")
	}
}

func TestFileExistInvalidfile(t *testing.T) {
	filename := "/some/incorrect/path.txt"
	if exist := fileExist(filename); exist {
		t.Fatal("Incorrect file exists.")
	}
}

func TestFileExistValidFile(t *testing.T) {
	filename := "/proc/cpuinfo"
	if exist := fileExist(filename); !exist {
		t.Fatal("Correct file doesn't exist.")
	}
}
