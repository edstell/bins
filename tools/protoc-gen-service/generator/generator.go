package generator

import (
	google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

type Generator interface {
	Generate(*google_protobuf.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error)
}

type GeneratorFunc func(*google_protobuf.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error)

func (f GeneratorFunc) Generate(desc *google_protobuf.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error) {
	return f(desc)
}
