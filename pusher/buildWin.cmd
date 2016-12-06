setlocal
%~d0
cd %~p0

for %%* in (.) do set CURR=%%~nx*
echo %CURR%
set FLAGS=-ldflags "-s -w"

set GOOS=windows
set GOARCH=386
go build %FLAGS% -o %CURR%.exe
upx --best %CURR%.exe
set GOARCH=amd64
go build %FLAGS% -o %CURR%_x64.exe
upx --best %CURR%_x64.exe
endlocal