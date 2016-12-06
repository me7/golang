setlocal
%~d0
cd %~p0

for %%* in (.) do set CURR=%%~nx*
echo %CURR%
set FLAGS=-ldflags "-s -w"

set GOOS=windows
set GOARCH=386
go build %FLAGS% -o bin/windows_x86/%CURR%.exe
upx --best bin/windows_x86/%CURR%.exe
set GOARCH=amd64
go build %FLAGS% -o bin/windows_x64/%CURR%.exe
upx --best bin/windows_x64/%CURR%.exe

set GOOS=linux
set GOARCH=386
go build %FLAGS% -o bin/linux_x86/%CURR%
upx --best bin/linux_x86/%CURR%
set GOARCH=amd64
go build %FLAGS% -o bin/linux_x64/%CURR%
upx --best bin/linux_x64/%CURR%

set GOOS=darwin
set GOARCH=386
go build %FLAGS% -o bin/darwin_x86/%CURR%
upx --best bin/darwin_x86/%CURR%
set GOARCH=amd64
go build %FLAGS% -o bin/darwin_x64/%CURR%
upx --best bin/darwin_x64/%CURR%

endlocal