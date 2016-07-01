package main

import "testing"

func TestParseIdFromLine(t *testing.T) {
	id, err := parseIdFromTaskLine("10 [x] Some task")
	if err != nil {
		t.Errorf("Some error parsing regexp")
	}
	if id != 10 {
		t.Errorf("Incorrect id number")
	}
}

func TestParseIdFromLine1000(t *testing.T) {
	id, err := parseIdFromTaskLine("1000 [x] Some task")
	if err != nil {
		t.Errorf("Some error parsing regexp")
	}
	if id != 1000 {
		t.Errorf("Incorrect id number")
	}
}

func TestParseIdFromLineEmptyID(t *testing.T) {
	id, err := parseIdFromTaskLine("[x] Some task")
	if err == nil {
		t.Errorf("Some error parsing regexp")
	}
	if id != 0 {
		t.Errorf("Incorrect id number %d", id)
	}
}

func TestGetLastId(t *testing.T) {
	id := getLastId()
	if id == 0 {
		t.Errorf("Bad last Id %d", id)
	}
}

func TestMarkLineAsDoneValidString(t *testing.T) {
	_, err := markLineAsDone("1     Some string")
	if err != nil {
		t.Errorf("Error not expected received")
	}
}

func TestMarkLineAsDoneValidString2(t *testing.T) {
	_, err := markLineAsDone("1 Some string")
	if err != nil {
		t.Errorf("Error not expected received")
	}
}

func TestMarkLineAsDoneInvalidString(t *testing.T) {
	_, err := markLineAsDone("1SomeString")
	if err == nil {
		t.Errorf("Error expected not received")
	}
}

func TestIsDoneLine(t *testing.T) {
	res := isDoneLine("1   Some string")
	if res {
		t.Errorf("Undone line was found as done")
	}
}

func TestIsDoneLineDone(t *testing.T) {
	res := isDoneLine("1 [x]  Some string")
	if !res {
		t.Errorf("Undone line was found as done")
	}
}

func TestIsDoneInvalidDone(t *testing.T) {
	res := isDoneLine("1[x] SomeString")
	if res {
		t.Errorf("Invalid done detected as valid")
	}
}

func TestInsertDoneMark(t *testing.T) {
	line := insertDoneMark("1 Some string")
	if line != "1 [x] Some string\n" {
		t.Errorf("Wrong string received: %s", line)
	}
}

func TestInsertDoneMark2(t *testing.T) {
	line := insertDoneMark("1                  Some string")
	if line != "1 [x] Some string\n" {
		t.Errorf("Wrong string received: %s", line)
	}
}

func TestInsertDoneMark3(t *testing.T) {
	line := insertDoneMark("1Some string")
	if line != "1Some string" {
		t.Errorf("Wrong string received")
	}
}

func TestClearTaskFile(t *testing.T) {
	InitTaskFile()
	err := ClearTaskFile()
	if err != nil {
		t.Errorf("Error clearing task file")
	}
}

func TestListAllTasks(t *testing.T) {
	err := ListAllTasks()
	if err != nil {
		t.Errorf("Error listing tasks")
	}
}
