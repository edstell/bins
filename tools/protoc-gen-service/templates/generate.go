package templates

//go:generate go-bindata -modtime 1 -pkg templates -ignore=\.go$ -ignore=.DS_Store .
//go:generate go fmt bindata.go
