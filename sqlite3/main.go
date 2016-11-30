package main

import (
	"database/sql"
	"log"
)

func main() {
	c, err := sql.Open("sqlite3", "db.sqlite3")
	if c == nil || err != nil {
		log.Fatalf("sql.Open() error: %v", err)
	}
	defer c.Close()
}
