package magicgate

import (
	"fmt"
	"net/http"
	"time"
)

const greeting = `Hello gopher
Nice to meet you all`

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s, time now %s", greeting, time.Now())
}
