env GOOS=linux GOARCH=amd64 go build -o /tmp/main ./service.api.recycling-services/main.go
zip -j /tmp/main.zip /tmp/main
