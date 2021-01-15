# protoc-gen-service
## Usage
After making a change:
1. Re-generate templates (if you changed them).
2. Re-install proto binary using go install.
```
go generate ./templates
go install ../protoc-gen-service
```
