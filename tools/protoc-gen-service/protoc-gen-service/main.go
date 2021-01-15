package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"github.com/edstell/lambda/tools/protoc-gen-router/templates"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	req := &plugin.CodeGeneratorRequest{}
	if err := proto.Unmarshal(data, req); err != nil {
		panic(err)
	}

	data, err = templates.Asset("router.gotmpl")
	if err != nil {
		panic(err)
	}

	t, err := template.New("router.gotmpl").Funcs(map[string]interface{}{
		"toLower": strings.ToLower,
		"rhs": func(s string) string {
			ss := strings.Split(s, ".")
			return ss[len(ss)-1]
		},
	}).Parse(string(data))
	if err != nil {
		panic(err)
	}

	files := make([]*plugin.CodeGeneratorResponse_File, 0, len(req.ProtoFile))
	for _, protoFile := range req.ProtoFile {
		name := strings.Replace(*protoFile.Name, "proto", "router.go", -1)
		file := &plugin.CodeGeneratorResponse_File{
			Name: &name,
		}

		var b bytes.Buffer
		if err := t.Execute(&b, protoFile); err != nil {
			panic(err)
		}

		content := b.String()
		file.Content = &content

		files = append(files, file)
	}

	data, err = proto.Marshal(&plugin.CodeGeneratorResponse{
		File: files,
	})
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(data)
}
