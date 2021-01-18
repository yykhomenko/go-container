package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("unknown command: " + os.Args[1])
	}
}

func run() {
	fmt.Printf("running %v\n", os.Args[2:])
}
