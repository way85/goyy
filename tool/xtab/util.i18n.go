// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"gopkg.in/goyy/goyy.v0/comm/i18n"
)

var i18N = i18n.New(locales, i18n.Locale_en_US)

var locales = map[string]map[string]string{
	i18n.Locale_zh_CN: zh_CN,
	i18n.Locale_en_US: en_US,
}

var zh_CN = map[string]string{
	"setting.empty":                     xsettings + ":configuration.settings:不能为空！",
	"setting.statement.empty":           xsettings + ":configuration.settings.statement:不能为空！",
	"setting.statement.seperator.empty": xsettings + ":configuration.settings.statement.@seperator:不能为空！",
	"setting.statement.case.empty":      xsettings + ":configuration.settings.statement.@case:不能为空！",
	"setting.statement.comment.empty":   xsettings + ":configuration.settings.statement.@comment:不能为空！",
	"project.empty":                     xprojects + ":configuration.projects:不能为空！",
	"project.project.empty":             xprojects + ":configuration.projects.project:不能为空！",
	"project.id.empty":                  xprojects + ":configuration.projects.project.@id:不能为空！",
	"project.database.empty":            xprojects + ":configuration.projects.project.@database:不能为空！",
	"module.empty":                      xmodules + ":configuration.modules:不能为空！",
	"module.module.empty":               xmodules + ":configuration.modules.module:不能为空！",
	"module.id.empty":                   xmodules + ":configuration.modules.module.@id:不能为空！",
	"module.project.empty":              xmodules + ":configuration.modules.module.@project:不能为空！",
	"module.project.errorf":             xmodules + ":configuration.modules.module.@project=%s:对应的项目不存在！",
	"domain.empty":                      xdomains + ":configuration.domains:不能为空！",
	"domain.domain.empty":               xdomains + ":configuration.domains.domain:不能为空！",
	"domain.id.empty":                   xdomains + ":configuration.domains.domain.@id:不能为空！",
	"domain.types.empty":                xdomains + ":configuration.domains.domain.@types:不能为空！",
	"domain.length.emptyf":              xdomains + ":configuration.domains.domain[@id=%s].@length:不能为空！",
	"domain.precision.emptyf":           xdomains + ":configuration.domains.domain[@id=%s].@precision:不能为空！",
	"column.empty":                      xcolumns + ":configuration.columns:不能为空！",
	"column.column.empty":               xcolumns + ":configuration.columns.column:不能为空！",
	"column.id.empty":                   xcolumns + ":configuration.columns.column.@id:不能为空！",
	"column.domain.empty":               xcolumns + ":configuration.columns.column.@domain:不能为空！",
	"column.domain.errorf":              xcolumns + ":configuration.columns.column.@domain=%s:对应的域不存在！",
	"table.emptyf":                      "%s:configuration.tables:不能为空！",
	"table.table.emptyf":                "%s:configuration.tables.table:不能为空！",
	"table.id.empty":                    xtables + ":configuration.tables.table.@id:不能为空！",
	"table.id.emptyf":                   "%s:configuration.tables.table.@id:不能为空！",
	"table.extends.errorf":              "%s:configuration.tables.table[@id=%s].@extends=%s:对应的表不存在！",
	"table.column.id.emptyf":            "%s:configuration.tables.table[@id=%s].column.@id:不能为空！",
	"table.column.extends.errorf":       "%s:configuration.tables.table[@id=%s].column.@extends=%s:对应的列不存在！",
	"table.column.domain.emptyf":        "%s:configuration.tables.table[@id=%s].column.@domain:不能为空！",
	"table.column.domain.errorf":        "%s:configuration.tables.table[@id=%s].column.@domain=%s:对应的域不存在！",
}

var en_US = map[string]string{
	"setting.empty":                     xsettings + ":configuration.settings:can not be empty!",
	"setting.statement.empty":           xsettings + ":configuration.settings.statement:can not be empty!",
	"setting.statement.seperator.empty": xsettings + ":configuration.settings.statement.@seperator:can not be empty!",
	"setting.statement.case.empty":      xsettings + ":configuration.settings.statement.@case:can not be empty!",
	"setting.statement.comment.empty":   xsettings + ":configuration.settings.statement.@comment:can not be empty!",
	"project.empty":                     xprojects + ":configuration.projects:can not be empty!",
	"project.project.empty":             xprojects + ":configuration.projects.project:can not be empty!",
	"project.id.empty":                  xprojects + ":configuration.projects.project.@id:can not be empty!",
	"project.database.empty":            xprojects + ":configuration.projects.project.@database:can not be empty!",
	"module.empty":                      xmodules + ":configuration.modules:can not be empty!",
	"module.module.empty":               xmodules + ":configuration.modules.module:can not be empty!",
	"module.id.empty":                   xmodules + ":configuration.modules.module.@id:can not be empty!",
	"module.project.empty":              xmodules + ":configuration.modules.module.@project:can not be empty!",
	"module.project.errorf":             xmodules + ":configuration.modules.module.@project=%s:match the project does not exist!",
	"domain.empty":                      xdomains + ":configuration.domains:can not be empty!",
	"domain.domain.empty":               xdomains + ":configuration.domains.domain:can not be empty!",
	"domain.id.empty":                   xdomains + ":configuration.domains.domain.@id:can not be empty!",
	"domain.types.empty":                xdomains + ":configuration.domains.domain.@types:can not be empty!",
	"domain.length.emptyf":              xdomains + ":configuration.domains.domain[@id=%s].@length:can not be empty!",
	"domain.precision.emptyf":           xdomains + ":configuration.domains.domain[@id=%s].@precision:can not be empty!",
	"column.empty":                      xcolumns + ":configuration.columns:can not be empty!",
	"column.column.empty":               xcolumns + ":configuration.columns.column:can not be empty!",
	"column.id.empty":                   xcolumns + ":configuration.columns.column.@id:can not be empty!",
	"column.domain.empty":               xcolumns + ":configuration.columns.column.@domain:can not be empty!",
	"column.domain.errorf":              xcolumns + ":configuration.columns.column.@domain=%s:match the domain does not exist!",
	"table.emptyf":                      "%s:configuration.tables:can not be empty!",
	"table.table.emptyf":                "%s:configuration.tables.table:can not be empty!",
	"table.id.empty":                    xtables + ":configuration.tables.table.@id:can not be empty!",
	"table.id.emptyf":                   "%s:configuration.tables.table.@id:can not be empty!",
	"table.extends.errorf":              "%s:configuration.tables.table[@id=%s].@extends=%s:match the table does not exist!",
	"table.column.id.emptyf":            "%s:configuration.tables.table[@id=%s].column.@id:can not be empty!",
	"table.column.extends.errorf":       "%s:configuration.tables.table[@id=%s].column.@extends=%s:match the column does not exist!",
	"table.column.domain.emptyf":        "%s:configuration.tables.table[@id=%s].column.@domain:can not be empty!",
	"table.column.domain.errorf":        "%s:configuration.tables.table[@id=%s].column.@domain=%s:match the domain does not exist!",
}
