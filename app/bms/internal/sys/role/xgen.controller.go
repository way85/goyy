// generated by xgen -- DO NOT EDIT
package role

import (
	"gopkg.in/goyy/goyy.v0/web/controller"
)

var ctl = &Controller{
	ClientController: controller.ClientController{
		Settings: controller.Settings{
			Project: "sys",
			Module:  "role",
			Title:   "",
		},
		DTO: func() interface{} {
			return &DTO{}
		},
		DTOs: func() interface{} {
			return &DTOs{}
		},
	},
}

type Controller struct {
	controller.ClientController
}
