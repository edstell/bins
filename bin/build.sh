env GOOS=linux GOARCH=amd64 go build -o /tmp/main $1
zip -j /tmp/main.zip /tmp/main
