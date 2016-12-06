package main

//go:generate cat "HELLO"

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	err := exec.Command("cmd", "/C", "start "+"http://localhost:7749").Run()
	if err != nil {
		log.Fatalf("open web error: %v", err)
	}
	log.Fatal(http.ListenAndServe(":7749", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gopher, time right now = %s", time.Now())
}
