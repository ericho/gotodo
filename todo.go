package main

import (
	//	"flag"
	"fmt"
	"os"
)

const (
	ADD = 1 << iota
	LIST
	DONE
	CLEAR
)

func readactionfromargs(arg string) (action int) {
	switch {
	case "add" == arg:
		action = ADD
	case "list" == arg:
		action = LIST
	case "done" == arg:
		action = DONE
	case "clear" == arg:
		action = CLEAR
	}
	return action
}

func readargs(args []string) (action int, err error) {

	if args == nil {
		return 0, nil
	}

	action = readactionfromargs(args[0])
	return action, err
}

func main() {
	readargs(os.Args[1:])
	fmt.Println("Lets print something")
}
