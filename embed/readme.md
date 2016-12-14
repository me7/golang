example for embed resource in your execute file (e.g. embed template folder)
refer to http://blog.ralch.com/tutorial/golang-embedded-resources/

1. `go-bindata template` to create bindata.go
2. instead of parsefile from filename directly `tpl := template.Must(template.ParseFiles("template/index.tpl"))`
3. become
  * `idx, _ := Asset("template/index.tpl")` to get []byte from bindata.go
  * `tpl := template.New("index")` make new template
  * `tpl.Parse(string(idx))` convert []byte to string then parse it
4. `tpl.Execute(w, data)` to execute template

template data = 5k
build without template = 1535488 byte
build with template =    1539072 byte