package guark

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/guark/guark/internal/generator"

	"github.com/kylelemons/godebug/diff"
)

var tmpl = `package test_pkg
//
// ------ AUTO GENERATED FILE (DO NOT EDIT) ------
//

import (
	"github.com/guark/guark/internal/embed"
)

var Embeds app.Embed = app.Embed{
	Embeds: map[string]*[]byte{
		{{- range $name, $embed := .embeds }}
		"{{ $name }}": {{- stringify $embed }},
		{{- end }}
	},
}
`

var expecting = `package test_pkg
//
// ------ AUTO GENERATED FILE (DO NOT EDIT) ------
//

import (
	"github.com/guark/guark/internal/embed"
)

var Embeds app.Embed = app.Embed{
	Embeds: map[string]*[]byte{
		"file1":[]byte{104,101,108,108,111},
		"file2":[]byte{119,111,114,108,100},
	},
}
`

func TestEmbedGenerator(t *testing.T) {

	files, err := filepath.Glob(fmt.Sprintf("%s/*", "./testdata/assets"))

	if err != nil {
		t.Error(err)
	}

	bytes, err := generator.Generate(files, "testdata/assets/", tmpl)

	if err != nil {
		t.Error(err)
	}

	if string(bytes) != expecting {
		t.Log(diff.Diff(string(bytes), expecting))
		t.Error("unexpected results!")
	}
}
