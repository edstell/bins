# protoc-gen-router
## Usage
```
go generate ./templates
go install ./protoc-gen-router
protoc --plugin protoc-gen-router --router_out=./ example.proto
```
