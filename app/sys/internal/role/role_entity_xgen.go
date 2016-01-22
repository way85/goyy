// generated by xgen -- DO NOT EDIT
package role

import (
	"bytes"
	"fmt"

	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

var (
	ENTITY           = schema.TABLE("sys_role", "ROLE")
	ENTITY_ID        = ENTITY.PRIMARY("id", "ID")
	ENTITY_MEMO      = ENTITY.COLUMN("memo", "MEMO")
	ENTITY_CREATES   = ENTITY.COLUMN("creates", "CREATES")
	ENTITY_CREATER   = ENTITY.CREATER("creater", "CREATER")
	ENTITY_CREATED   = ENTITY.CREATED("created", "CREATED")
	ENTITY_MODIFIER  = ENTITY.MODIFIER("modifier", "MODIFIER")
	ENTITY_MODIFIED  = ENTITY.MODIFIED("modified", "MODIFIED")
	ENTITY_VERSION   = ENTITY.VERSION("version", "VERSION")
	ENTITY_DELETION  = ENTITY.DELETION("deletion", "DELETION")
	ENTITY_ARTIFICAL = ENTITY.COLUMN("artifical", "ARTIFICAL")
	ENTITY_HISTORY   = ENTITY.COLUMN("history", "HISTORY")
	ENTITY_NAME      = ENTITY.COLUMN("name", "NAME")
	ENTITY_CODE      = ENTITY.COLUMN("code", "CODE")
	ENTITY_GENRE     = ENTITY.COLUMN("genre", "GENRE")
	ENTITY_ORDINAL   = ENTITY.COLUMN("ordinal", "ORDINAL")
)

func NewEntity() *Entity {
	e := &Entity{}
	e.init()
	return e
}

func (me *Entity) Name() string {
	return me.name.Value()
}

func (me *Entity) SetName(v string) {
	me.name.SetValue(v)
}

func (me *Entity) Code() string {
	return me.code.Value()
}

func (me *Entity) SetCode(v string) {
	me.code.SetValue(v)
}

func (me *Entity) Genre() string {
	return me.genre.Value()
}

func (me *Entity) SetGenre(v string) {
	me.genre.SetValue(v)
}

func (me *Entity) Ordinal() string {
	return me.ordinal.Value()
}

func (me *Entity) SetOrdinal(v string) {
	me.ordinal.SetValue(v)
}

func (me *Entity) init() {
	me.table = ENTITY
	me.initSetDict()
	me.initSetColumn()
	me.initSetDefault()
	me.initSetField()
	me.initSetExcel()
}

func (me *Entity) initSetDict() {
}

func (me *Entity) initSetColumn() {
	if t, ok := me.Sys.Type("id"); ok {
		t.SetColumn(ENTITY_ID)
	}
	if t, ok := me.Sys.Type("memo"); ok {
		t.SetColumn(ENTITY_MEMO)
	}
	if t, ok := me.Sys.Type("creates"); ok {
		t.SetColumn(ENTITY_CREATES)
	}
	if t, ok := me.Sys.Type("creater"); ok {
		t.SetColumn(ENTITY_CREATER)
	}
	if t, ok := me.Sys.Type("created"); ok {
		t.SetColumn(ENTITY_CREATED)
	}
	if t, ok := me.Sys.Type("modifier"); ok {
		t.SetColumn(ENTITY_MODIFIER)
	}
	if t, ok := me.Sys.Type("modified"); ok {
		t.SetColumn(ENTITY_MODIFIED)
	}
	if t, ok := me.Sys.Type("version"); ok {
		t.SetColumn(ENTITY_VERSION)
	}
	if t, ok := me.Sys.Type("deletion"); ok {
		t.SetColumn(ENTITY_DELETION)
	}
	if t, ok := me.Sys.Type("artifical"); ok {
		t.SetColumn(ENTITY_ARTIFICAL)
	}
	if t, ok := me.Sys.Type("history"); ok {
		t.SetColumn(ENTITY_HISTORY)
	}
	me.name.SetColumn(ENTITY_NAME)
	me.code.SetColumn(ENTITY_CODE)
	me.genre.SetColumn(ENTITY_GENRE)
	me.ordinal.SetColumn(ENTITY_ORDINAL)
}

func (me *Entity) initSetDefault() {
	if t, ok := me.Sys.Type("created"); ok {
		t.SetDefault("-62135596800")
	}
	if t, ok := me.Sys.Type("modified"); ok {
		t.SetDefault("-62135596800")
	}
}

func (me *Entity) initSetField() {
	for _, c := range entity.SysColumns {
		if t, ok := me.Sys.Type(c); ok {
			t.SetField(entity.DefaultField())
		}
	}
	me.name.SetField(entity.DefaultField())
	me.code.SetField(entity.DefaultField())
	me.genre.SetField(entity.DefaultField())
	me.ordinal.SetField(entity.DefaultField())
}

func (me *Entity) initSetExcel() {
}

func (me Entity) New() entity.Interface {
	return NewEntity()
}

func (me *Entity) Get(column string) interface{} {
	switch column {
	case ENTITY_NAME.Name():
		return me.name.Value()
	case ENTITY_CODE.Name():
		return me.code.Value()
	case ENTITY_GENRE.Name():
		return me.genre.Value()
	case ENTITY_ORDINAL.Name():
		return me.ordinal.Value()
	}
	return me.Sys.Get(column)
}

func (me *Entity) GetPtr(column string) interface{} {
	switch column {
	case ENTITY_NAME.Name():
		return me.name.ValuePtr()
	case ENTITY_CODE.Name():
		return me.code.ValuePtr()
	case ENTITY_GENRE.Name():
		return me.genre.ValuePtr()
	case ENTITY_ORDINAL.Name():
		return me.ordinal.ValuePtr()
	}
	return me.Sys.GetPtr(column)
}

func (me *Entity) Table() schema.Table {
	return me.table
}

func (me *Entity) Type(column string) (entity.Type, bool) {
	switch column {
	case ENTITY_NAME.Name():
		return &me.name, true
	case ENTITY_CODE.Name():
		return &me.code, true
	case ENTITY_GENRE.Name():
		return &me.genre, true
	case ENTITY_ORDINAL.Name():
		return &me.ordinal, true
	}
	return me.Sys.Type(column)
}

func (me *Entity) Column(field string) (schema.Column, bool) {
	switch strings.ToLowerFirst(field) {
	case "name":
		return ENTITY_NAME, true
	case "code":
		return ENTITY_CODE, true
	case "genre":
		return ENTITY_GENRE, true
	case "ordinal":
		return ENTITY_ORDINAL, true
	}
	return me.Sys.Column(field)
}

func (me *Entity) Columns() []schema.Column {
	return []schema.Column{
		ENTITY_ID,
		ENTITY_MEMO,
		ENTITY_CREATES,
		ENTITY_CREATER,
		ENTITY_CREATED,
		ENTITY_MODIFIER,
		ENTITY_MODIFIED,
		ENTITY_VERSION,
		ENTITY_DELETION,
		ENTITY_ARTIFICAL,
		ENTITY_HISTORY,
		ENTITY_NAME,
		ENTITY_CODE,
		ENTITY_GENRE,
		ENTITY_ORDINAL,
	}
}

func (me *Entity) Names() []string {
	return []string{
		"id",
		"memo",
		"creates",
		"creater",
		"created",
		"modifier",
		"modified",
		"version",
		"deletion",
		"artifical",
		"history",
		"name",
		"code",
		"genre",
		"ordinal",
	}
}

func (me *Entity) Value() *Entity {
	return me
}

func (me *Entity) SetString(field, value string) error {
	switch strings.ToLowerFirst(field) {
	case "name":
		return me.name.SetString(value)
	case "code":
		return me.code.SetString(value)
	case "genre":
		return me.genre.SetString(value)
	case "ordinal":
		return me.ordinal.SetString(value)
	}
	return me.Sys.SetString(field, value)
}

func (me *Entity) Validate() error {
	return nil
}

func (me *Entity) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(fmt.Sprintf(`"id":%q`, me.Sys.Pk.Id()))
	b.WriteString(fmt.Sprintf(`,"memo":%q`, me.Sys.Memo()))
	b.WriteString(fmt.Sprintf(`,"creates":%q`, me.Sys.Creates()))
	b.WriteString(fmt.Sprintf(`,"creater":%q`, me.Sys.Creater()))
	b.WriteString(fmt.Sprintf(`,"created":%d`, me.Sys.Created()))
	b.WriteString(fmt.Sprintf(`,"modifier":%q`, me.Sys.Modifier()))
	b.WriteString(fmt.Sprintf(`,"modified":%d`, me.Sys.Modified()))
	b.WriteString(fmt.Sprintf(`,"version":%d`, me.Sys.Version()))
	b.WriteString(fmt.Sprintf(`,"deletion":%d`, me.Sys.Deletion()))
	b.WriteString(fmt.Sprintf(`,"artifical":%d`, me.Sys.Artifical()))
	b.WriteString(fmt.Sprintf(`,"history":%d`, me.Sys.History()))
	b.WriteString(fmt.Sprintf(`,"name":%q`, me.name.String()))
	b.WriteString(fmt.Sprintf(`,"code":%q`, me.code.String()))
	b.WriteString(fmt.Sprintf(`,"genre":%q`, me.genre.String()))
	b.WriteString(fmt.Sprintf(`,"ordinal":%q`, me.ordinal.String()))
	b.WriteString("}")
	return b.String()
}

func (me *Entity) ExcelColumns() []string {
	return nil
}
