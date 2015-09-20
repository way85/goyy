// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplService = `// generated by xgen -- DO NOT EDIT
package {{.PackageName}}

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/service"
){{range $i, $e := .Entities}}{{with $name := mgrname $e.Name $e.Relationship}}

var {{mgrvar $e.Name $e.Relationship}} = &{{$name}}{
	Manager: service.Manager{
		Entity: func() entity.Interface {
			return New{{$e.Name}}()
		},
		Entities: func() entity.Interfaces {
			return New{{entities $e.Name}}(10)
		},
	},
}

type {{$name}} struct {
	service.Manager
}{{end}}{{end}}
`
