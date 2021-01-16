package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/edstell/lambda/tools/protoc-gen-service/generator"
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
		panic(fmt.Sprint("failed to unmarshal the generation request:", err))
	}
	funcs := map[string]interface{}{
		"toLower": strings.ToLower,
		"rhs": func(s string) string {
			ss := strings.Split(s, ".")
			return ss[len(ss)-1]
		},
	}
	generators := []generator.Generator{
		generator.FromTemplate("router", funcs),
		generator.FromTemplate("client", funcs),
	}
	files := make([]*plugin.CodeGeneratorResponse_File, 0, len(req.ProtoFile)*len(generators))
	for _, protoFile := range req.ProtoFile {
		if len(protoFile.Service) == 0 {
			continue
		}
		for _, generator := range generators {
			file, err := generator.Generate(protoFile)
			if err != nil {
				panic(fmt.Sprint("failed to generate file", err))
			}
			files = append(files, file)
		}
	}
	data, err = proto.Marshal(&plugin.CodeGeneratorResponse{
		File: files,
	})
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(data)
}
