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
