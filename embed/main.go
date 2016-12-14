package main

import (
	"html/template"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", indexHandler)
	exec.Command("cmd", "/c start http://localhost:7749")
	log.Fatal(http.ListenAndServe(":7749", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := "Embeded template"

	//tpl := template.Must(template.ParseFiles("template/index.tpl"))
	idx, _ := Asset("template/index.tpl")
	tpl := template.New("index")
	tpl.Parse(string(idx))
	tpl.Execute(w, data)
}
