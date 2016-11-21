* download and app engine sdk for go
* create simple http.ListenAndServe app (go run main.go)

### modify to run on app engine
* create app.yaml
* modify main.go
`windows have bug on dev_appserver.py, need to comment line 53-54 of file_watcher.py`
* go app serve --> should run OK on localhost:8000 now
* goto https://console.cloud.google.com/appengine and create new project name magic-gate
* goapp deploy
* see project at `your-app-name`.appspot.com e.g. https://magic-gate-150208.appspot.com/
