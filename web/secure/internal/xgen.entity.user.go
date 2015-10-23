// generated by xgen -- DO NOT EDIT
package internal

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

var (
	USER            = schema.TABLE("sys_user")
	USER_ID         = USER.PRIMARY("id")
	USER_NAME       = USER.COLUMN("name")
	USER_LOGIN_NAME = USER.COLUMN("login_name")
)

func NewUser() *User {
	e := &User{}
	e.init()
	return e
}

func (me *User) Id() string {
	return me.id.Value()
}

func (me *User) SetId(v string) {
	me.id.SetValue(v)
}

func (me *User) Name() string {
	return me.name.Value()
}

func (me *User) SetName(v string) {
	me.name.SetValue(v)
}

func (me *User) LoginName() string {
	return me.loginName.Value()
}

func (me *User) SetLoginName(v string) {
	me.loginName.SetValue(v)
}

func (me *User) init() {
	me.table = USER
	me.id.SetColumn(USER_ID)
	me.name.SetColumn(USER_NAME)
	me.loginName.SetColumn(USER_LOGIN_NAME)
	me.id.SetField(entity.DefaultField())
	me.name.SetField(entity.DefaultField())
	me.loginName.SetField(entity.DefaultField())
}

func (me User) New() entity.Interface {
	return NewUser()
}

func (me *User) Get(column string) interface{} {
	switch column {
	case USER_ID.Name():
		return me.id.Value()
	case USER_NAME.Name():
		return me.name.Value()
	case USER_LOGIN_NAME.Name():
		return me.loginName.Value()
	}
	return nil
}

func (me *User) GetPtr(column string) interface{} {
	switch column {
	case USER_ID.Name():
		return me.id.ValuePtr()
	case USER_NAME.Name():
		return me.name.ValuePtr()
	case USER_LOGIN_NAME.Name():
		return me.loginName.ValuePtr()
	}
	return nil
}

func (me *User) Table() schema.Table {
	return me.table
}

func (me *User) Type(column string) (entity.Type, bool) {
	switch column {
	case USER_ID.Name():
		return &me.id, true
	case USER_NAME.Name():
		return &me.name, true
	case USER_LOGIN_NAME.Name():
		return &me.loginName, true
	}
	return nil, false
}

func (me *User) Column(field string) (schema.Column, bool) {
	switch strings.ToLowerFirst(field) {
	case "id":
		return USER_ID, true
	case "name":
		return USER_NAME, true
	case "loginName":
		return USER_LOGIN_NAME, true
	}
	return nil, false
}

func (me *User) Columns() []schema.Column {
	return []schema.Column{
		USER_ID,
		USER_NAME,
		USER_LOGIN_NAME,
	}
}

func (me *User) Names() []string {
	return []string{
		"id",
		"name",
		"loginName",
	}
}

func (me *User) Value() *User {
	return me
}

func (me *User) SetString(field, value string) error {
	switch strings.ToLowerFirst(field) {
	case "id":
		return me.id.SetString(value)
	case "name":
		return me.name.SetString(value)
	case "loginName":
		return me.loginName.SetString(value)
	}
	return nil
}

func (me *User) Validate() error {
	return nil
}
