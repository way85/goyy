// generated by xgen -- DO NOT EDIT
package post

import (
	"bytes"
	"gopkg.in/goyy/goyy.v0/data/entity"
)

func NewMenuEntities(v int) *MenuEntities {
	entities := &MenuEntities{}
	entities.Make(v)
	return entities
}

type MenuEntities struct {
	datas []*MenuEntity
}

func (me *MenuEntities) Make(cap int) {
	me.datas = make([]*MenuEntity, 0, cap)
}

func (me *MenuEntities) New() entity.Interface {
	return NewMenuEntity()
}

func (me *MenuEntities) Append(v entity.Interface) {
	me.datas = append(me.datas, v.(*MenuEntity))
}

func (me *MenuEntities) Len() int {
	return len(me.datas)
}

func (me *MenuEntities) Cap() int {
	return cap(me.datas)
}

func (me *MenuEntities) Index(v int) entity.Interface {
	return me.datas[v]
}

func (me *MenuEntities) Slice() interface{} {
	return me.datas
}

func (me *MenuEntities) Value(v int) *MenuEntity {
	return me.datas[v]
}

func (me *MenuEntities) Values() []*MenuEntity {
	return me.datas
}

func (me *MenuEntities) JSON() string {
	var b bytes.Buffer
	b.WriteString(`{"datas":[`)
	for i := 0; i < me.Len(); i++ {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString(me.Index(i).JSON())
	}
	b.WriteString("]}")
	return b.String()
}
