// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplDtoMain = `// generated by xgen -- DO NOT EDIT
package {{.PackageName}}

import ({{if .IsExtend}}
	"gopkg.in/goyy/goyy.v0/data/dto"{{end}}{{if .IsTimeField}}
	"time"{{end}}
){{range $e := .Entities}}{{with $name := dtoname $e.Name $e.Relationship}}

type {{dtosname $e.Name $e.Relationship}} []{{$name}}

type {{$name}} struct {{"{"}}{{if eq $e.Extend "pk"}}
	dto.Pk{{else if eq $e.Extend "sys"}}
	dto.Sys{{else if eq $e.Extend "tree"}}
	dto.Tree{{end}}{{range $f := $e.Fields}}
	{{fname $f.Name $e.FieldMaxLen}} {{padright $f.Type $e.TypeMaxLen}} ` + "`" + `json:"{{$f.Name}}"` + "`" + `{{end}}
}{{end}}{{end}}
`
