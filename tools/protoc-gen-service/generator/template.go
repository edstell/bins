package generator

import (
	"bytes"
	"fmt"
	"html/template"
	"path"

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
		filepath := *desc.Name
		name := filepath[0:len(filepath)-len(path.Ext(filepath))] + fmt.Sprintf(".%s.go", name)
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
