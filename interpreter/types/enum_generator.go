//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"text/template"
)

var codeTemplate = template.Must(template.New("").Parse(`// Code generated by enum_generator; DO NOT EDIT.
//go:generate go run enum_generator.go
package types

import (
	"fmt"

	"go.starlark.net/starlark"
	"gopkg.in/gographics/imagick.v2/imagick"
)

type {{ .Type }} struct {
	Value     imagick.{{ .Type }}
	stringVal string
}

func (v {{ .Type }}) String() string {
	return fmt.Sprintf("{{ .Type }} %s", v.stringVal)
}

func (v {{ .Type }}) Type() string {
	return "{{ .Type }}"
}

func (v {{ .Type }}) Freeze() {
	return
}

func (v {{ .Type }}) Truth() starlark.Bool {
	return true
}

func (v {{ .Type }}) Hash() (uint32, error) {
	return uint32(v.Value), nil
}

var _ starlark.Value = (*{{ .Type }})(nil)

var _{{ .Type }}Map = map[string]{{ .Type }} {
	{{- range $val := .Values }}
	"{{ $val }}": { imagick.{{ $val }}, "{{ $val }}" },
	{{- end }}
}

var _{{ .Type }}Names = []string{
	{{- range $val := .Values }}
	"{{ $val }}",
	{{- end }}
}

type {{ .Type }}Enum struct {}

func (v {{ .Type }}Enum) String() string {
	return "{{ .Type }}Enum"
}

func (v {{ .Type }}Enum) Type() string {
	return "{{ .Type }}Enum"
}

func (v {{ .Type }}Enum) Freeze() {
	return
}

func (v {{ .Type }}Enum) Truth() starlark.Bool {
	return true
}

func (v {{ .Type }}Enum) Hash() (uint32, error) {
	return 0, nil
}

func (v {{ .Type }}Enum) Attr(name string) (starlark.Value, error) {
	val, found := _{{ .Type }}Map[name]
	if !found {
		return nil, starlark.NoSuchAttrError(fmt.Sprintf("unknown {{ .Type }} %s", name))
	}
	return val, nil
}

func (v {{ .Type }}Enum) AttrNames() []string {
	return _{{ .Type }}Names
}

var _ starlark.Value = (*{{ .Type }}Enum)(nil)
var _ starlark.HasAttrs = (*{{ .Type }}Enum)(nil)
`))

func main() {
	fileName := os.Getenv("GOFILE")

	path, err := build.Default.Import("gopkg.in/gographics/imagick.v2/imagick", ".", build.FindOnly)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(
		fset,
		fmt.Sprintf("%s/%s", path.Dir, fileName),
		nil,
		parser.ParseComments,
	)
	if err != nil {
		panic(err)
	}

	var values []string
	var foundTypeName string

	ast.Inspect(file, func(node ast.Node) bool {
		switch node.(type) {
		case *ast.ValueSpec:
			value := node.(*ast.ValueSpec)
			values = append(values, value.Names[0].Name)
		case *ast.TypeSpec:
			typeSpec := node.(*ast.TypeSpec)
			foundTypeName = typeSpec.Name.Name
		}
		return true
	})

	outBuffer := new(bytes.Buffer)

	codeTemplate.Execute(
		outBuffer,
		struct {
			Type     string
			FileName string
			Values   []string
		}{
			Type:     foundTypeName,
			FileName: fileName,
			Values:   values,
		},
	)

	formattedBytes, err := format.Source(outBuffer.Bytes())
	if err != nil {
		panic(err)
	}

	outFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}

	_, err = outFile.Write(formattedBytes)
	if err != nil {
		panic(err)
	}
}
