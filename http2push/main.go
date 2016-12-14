package main

import (
	"fmt"
	"log"
	"net/http"
)

const indexHTML = `<html>
<head>
	<title>Hello</title>
	<script src="/main.js"></script>
</head>
<body>
</body>
</html>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		pusher, ok := w.(http.Pusher)
		if ok { // Push is supported. Try pushing rather than waiting for the browser.
			if err := pusher.Push("/main.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		fmt.Fprintf(w, indexHTML)
	})

	//start browser and go to https://127.0.0.1:7749/

	// Run crypto/tls/generate_cert.go to generate cert.pem and key.pem.
	// See https://golang.org/src/crypto/tls/generate_cert.go
	log.Fatal(http.ListenAndServeTLS(":7749", "cert.pem", "key.pem", nil))
}
