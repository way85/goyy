// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xsql_test

import (
	_ "github.com/mattn/go-oci8"
	_ "github.com/ziutek/mymysql/godrv"
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/xsql"
)

var factory xsql.Factory
var session xsql.Session

func init() {
	factory, err := xsql.New(&dialect.Oracle{}, "db")
	if err != nil {
		panic(err)
	}
	session, err = factory.Session()
	if err != nil {
		panic(err)
	}
}
