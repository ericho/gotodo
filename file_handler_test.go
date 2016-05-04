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

func TestReadFileIntoArray(t *testing.T) {
	filename := "/proc/cpuinfo"
	if l, _ := readFileIntoArray(filename); len(l) == 0 {
		t.Fatal("File cannot be read")
	}
}

func TestReadFileIntoArrayInvalidArgs(t *testing.T) {
	filename := "Somefile.txt"
	if l, err := readFileIntoArray(filename); len(l) != 0 && err != nil {
		t.Fatal("File cannot be read")
	}
}
