use https://github.com/cratonica/trayhost to create system tray app
To use this lib you will need gcc.exe to compile c code

also need https://github.com/cratonica/2goarray to convert icon file to go file
`2goarray iconData main < icon.ico > icon.go`

when build need to set flag that hide terminal window
`go build -ldflags -H=windowsgui`

see build.cmd file for more info of build