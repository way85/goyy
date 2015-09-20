// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"strconv"
)

type Uint struct {
	base
	value uint
}

func (me *Uint) Value() uint {
	return me.value
}

func (me *Uint) ValuePtr() *uint {
	return &me.value
}

func (me *Uint) SetValue(v uint) {
	me.value = v
	me.field.SetModified(true)
}

func (me *Uint) SetDefault(v string) error {
	if strings.IsBlank(v) {
		return errors.NewNotBlank("v")
	}
	val, err := strconv.ParseInt(v, 10, 0)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	me.value = uint(val)
	return nil
}

func (me *Uint) SetString(v string) error {
	if err := me.SetDefault(v); err != nil {
		return err
	} else {
		me.field.SetModified(true)
		return nil
	}
}

func (me *Uint) String() string {
	return strconv.Itoa(int(me.value))
}

func (me *Uint) Name() string {
	return "uint"
}
