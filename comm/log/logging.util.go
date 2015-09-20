// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/files"
	"os"
	"time"
)

func (me *Logging) setConsoleLogger() {
	me.console = NewLogger(os.Stderr)
	me.console.SetPriority(me.priority)
	if me.layouts > 0 {
		me.console.SetLayouts(me.layouts)
	}
	me.console.SetPrefix(me.prefix)
}

func (me *Logging) resetConsoleLogger() {
	if me.outputs&Oconsole != 0 {
		if me.console == nil {
			me.setConsoleLogger()
		}
	}
}

func (me *Logging) getDailyFileName() string {
	date := time.Now().Format("2006-01-02")
	format := "./%s/daily.%s.log"
	return fmt.Sprintf(format, logDir, date)
}

func (me *Logging) setDailyFileLogger() {
	me.dailyfilename = me.getDailyFileName()
	if files.IsExist(me.dailyfilename) {
		f, _ := os.OpenFile(me.dailyfilename, os.O_APPEND|os.O_RDWR, 0666)
		me.dailyfile = NewLogger(f)
	} else {
		os.Mkdir(logDir, 0666)
		f, _ := os.Create(me.dailyfilename)
		me.dailyfile = NewLogger(f)
	}
	me.dailyfile.SetPriority(me.priority)
	if me.layouts > 0 {
		me.dailyfile.SetLayouts(me.layouts)
	}
	me.dailyfile.SetPrefix(me.prefix)
}

func (me *Logging) resetDailyFileLogger() {
	if me.outputs&Odailyfile != 0 {
		if me.dailyfilename != me.getDailyFileName() {
			me.setDailyFileLogger()
		}
	}
}
