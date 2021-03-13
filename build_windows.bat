
@echo off
for /F %%i in ('go env GOOS') do ( set os=%%i)
for /F %%i in ('go env GOARCH') do ( set arch=%%i)
for /F %%i in ('go env GOVERSION') do ( set goversion=%%i)
for /F %%i in ('git rev-parse --short HEAD') do ( set commitid=%%i)
for /F %%i in ('git log --pretty^=format:"%%an" -1') do ( set account=%%i)
for /F "tokens=* delims=" %%i in ('git branch --show-current') do ( set branch=%%i)
for /f "tokens=* delims=" %%i in ('echo %date:~0,4%-%date:~5,2%-%date:~8,2%.%time:~1,1%:%time:~3,2%:%time:~6,2%') do ( set nowtime="%%i")
go build -ldflags "-X github.com/yuansudong/http_dns/cmdline._GitBranch=%branch% -X github.com/yuansudong/http_dns/cmdline._OS=%os% -X github.com/yuansudong/http_dns/cmdline._Arch=%arch% -X github.com/yuansudong/http_dns/cmdline._GoVersion=%goversion% -X github.com/yuansudong/http_dns/cmdline._GitCommit=%commitid% -X github.com/yuansudong/http_dns/cmdline._GitAccount=%account% -X github.com/yuansudong/http_dns/cmdline._DateTime=%nowtime%" -o http_dns.exe
