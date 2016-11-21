* download and app engine sdk for go
* create simple http.ListenAndServe app (go run main.go)

### modify to run on app engine
* create app.yaml
* modify main.go
`windows have bug on dev_appserver.py, need to comment line 53-54 of file_watcher.py`
* go app serve --> should run OK on localhost:8000 now
