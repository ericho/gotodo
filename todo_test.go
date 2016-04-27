package main

import "testing"

func TestCommandLineParser(t *testing.T) {
	var s []string
	_, err := readargs(s)
	if err != nil {
		t.Failed()
	}
}

func TestCommandLineParserWithAdd(t *testing.T) {
	s := []string{"add", "somestring"}
	res, err := readargs(s)
	if err != nil {
		t.Fatal("Error received")
	}
	if res != ADD {
		t.Fatal("Incorrect result value")
	}
}

func TestCommandLineParserWithList(t *testing.T) {
	s := []string{"list", "somestring"}
	res, err := readargs(s)
	if err != nil {
		t.Fatal("Error received")
	}
	if res != LIST {
		t.Fatal("Incorrect result value")
	}
}

func TestCommandLineParserWithDone(t *testing.T) {
	s := []string{"done", "somestring"}
	res, err := readargs(s)
	if err != nil {
		t.Fatal("Error received")
	}
	if res != DONE {
		t.Fatal("Incorrect result value")
	}
}

func TestCommandLineParserWithClear(t *testing.T) {
	s := []string{"clear", "somestring"}
	res, err := readargs(s)
	if err != nil {
		t.Fatal("Error received")
	}
	if res != CLEAR {
		t.Fatal("Incorrect result value")
	}
}

func TestCommandLineParserWithInvalid(t *testing.T) {
	s := []string{"invalid", "somestring"}
	res, err := readargs(s)
	if res != 0 {
		t.Fatal("Incorrect action received")
	}
	if err == nil {
		t.Fatal("Expected error not received")
	}
}

func TestExecuteOptionInvalid(t *testing.T) {
	res, err := executeoption(100, nil)
	if err == nil {
		t.Fatal("Expected error not found")
	}
	if res == 0 {
		t.Fatal("Invalid return value")
	}
}
