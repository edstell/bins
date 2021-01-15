package generator

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/edstell/lambda/tools/protoc-gen-service/templates"
	google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

func FromTemplate(name string, funcs map[string]interface{}) Generator {
	data, err := templates.Asset(fmt.Sprintf("%s.gotmpl", name))
	if err != nil {
		panic(err)
	}
	t, err := template.New(name).Funcs(funcs).Parse(string(data))
	if err != nil {
		panic(err)
	}
	return GeneratorFunc(func(desc *google_protobuf.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error) {
		name := strings.Replace(*desc.Name, "proto", fmt.Sprintf("%s.go", name), -1)
		var b bytes.Buffer
		if err := t.Execute(&b, desc); err != nil {
			return nil, err
		}
		content := b.String()
		return &plugin.CodeGeneratorResponse_File{
			Name:    &name,
			Content: &content,
		}, nil
	})
}
