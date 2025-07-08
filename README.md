Change line 13 of auth/jwt.go to your secret key, then run
<br/>`go env -w GO111MODULE=auto`
<br/>`cd hr-api-go && go get -u`
<br/>`go build -o hr-api`/`go build -o hr-api.exe` or `go run api.go`
