// generated by xgen -- DO NOT EDIT
package role

import (
	"bytes"
	"gopkg.in/goyy/goyy.v0/data/entity"
)

func NewPostEntities(v int) *PostEntities {
	entities := &PostEntities{}
	entities.Make(v)
	return entities
}

type PostEntities struct {
	datas []*PostEntity
}

func (me *PostEntities) Make(cap int) {
	me.datas = make([]*PostEntity, 0, cap)
}

func (me *PostEntities) New() entity.Interface {
	return NewPostEntity()
}

func (me *PostEntities) Append(v entity.Interface) {
	me.datas = append(me.datas, v.(*PostEntity))
}

func (me *PostEntities) Len() int {
	return len(me.datas)
}

func (me *PostEntities) Cap() int {
	return cap(me.datas)
}

func (me *PostEntities) Index(v int) entity.Interface {
	return me.datas[v]
}

func (me *PostEntities) Slice() interface{} {
	return me.datas
}

func (me *PostEntities) Value(v int) *PostEntity {
	return me.datas[v]
}

func (me *PostEntities) Values() []*PostEntity {
	return me.datas
}

func (me *PostEntities) JSON() string {
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
