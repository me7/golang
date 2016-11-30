package main

import (
	"database/sql"
	"log"

	_ "github.com/mxk/go-sqlite/sqlite3"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	c, err := sql.Open("sqlite3", "db.sqlite3")
	if c == nil || err != nil {
		log.Fatalf("sql.Open() error: %v", err)
	}
	defer c.Close()

	sql := "CREATE TABLE x(a,b,c)"
	r, err := c.Exec(sql)
	if r == nil || err != nil {
		log.Fatalf("c.Exec(%q) error: %v", sql, err)
	}

	sql = "INSERT INTO x VALUES (?,?,?)"
	s, err := c.Prepare(sql)
	if s == nil || err != nil {
		log.Fatalf("c.Prepare(%q) error: %v", sql, err)
	}

	r, err = s.Exec(1, 2.2, "test")
	if err != nil {
		log.Fatalf("s.Exec(%q) error: %v", sql, err)
	}

	r, err = s.Exec(3, []byte{4}, "blank")
	if err != nil {
		log.Fatalf("s.Exec(%q) error: %v", sql, err)
	}

	var (
		id int
		ca string
		cb string
		cc string
	)
	sql = "SELECT rowid, * FROM x ORDER BY rowid"
	rows, err := c.Query(sql)
	if rows == nil || err != nil {
		log.Fatalf("c.Query() error: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &ca, &cb, &cc)
		if err != nil {
			log.Fatalf("scan error: %v", err)
		}
		log.Println(ca, cb, cc)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

}
