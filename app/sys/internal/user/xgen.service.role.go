// generated by xgen -- DO NOT EDIT
package user

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/service"
)

var RoleMgr = &RoleManager{
	Manager: service.Manager{
		Entity: func() entity.Interface {
			return NewRoleEntity()
		},
		Entities: func() entity.Interfaces {
			return NewRoleEntities(10)
		},
	},
}

type RoleManager struct {
	service.Manager
}
