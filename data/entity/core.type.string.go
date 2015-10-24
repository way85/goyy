// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"bytes"
)

type String struct {
	base
	value []byte
}

func (me *String) Value() string {
	return string(me.value)
}

func (me *String) ValuePtr() *[]byte {
	return &me.value
}

func (me *String) SetValue(v string) {
	me.value = []byte(v)
	me.field.SetModified(true)
}

func (me *String) SetDefault(v string) error {
	me.value = []byte(v)
	return nil
}

func (me *String) SetString(v string) error {
	if err := me.SetDefault(v); err != nil {
		return err
	} else {
		me.field.SetModified(true)
		return nil
	}
}

func (me *String) String() string {
	out := bytes.TrimRight(me.value, "\x00")
	return string(out)
}

func (me *String) Name() string {
	return "string"
}
