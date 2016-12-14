package main

import (
	"log"
	"os/exec"
)

func main() {
	err := exec.Command("cmd", "/c start http://localhost:7749/pkg/github.com/me7").Run()
	if err != nil {
		log.Fatalf("open web error: %v", err)
	}

	println("Press Ctrl+C to stop application")

	exec.Command("godoc", "-http=localhost:7749").Run()
	if err != nil {
		log.Fatalf("open godoc error: %v", err)
	}

}
