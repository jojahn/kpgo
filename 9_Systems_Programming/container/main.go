package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("help")
	}
}

func must_(err error) {
	if err != nil {
		panic(err)
	}
}

func run_() {
	log.Printf("Running %v \n", os.Args[1:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(cmd.Run())
}
