set GOARCH=amd64
set GOOS=linux
go build

ps
$env:GOARCH="amd64"
$env:GOOS="linux"
go build