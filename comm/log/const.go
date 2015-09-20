// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

// Priority used for identifying the severity of an event.
const (
	Pall = iota
	Ptrace
	Pdebug
	Pinfo
	Pwarn
	Perror
	Pcritical
	Pprint
	Poff
)

// These layouts define which text to prefix to each log entry generated by the Logger.
const (
	// Debug 2012/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	Ldate      = 1 << iota                 // the date: 2012/01/23
	Ltime                                  // the time: 01:23:23
	Lmicrosec                              // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                              // full file name and line number: /a/b/c/d.go:23
	Lshortfile                             // final file name element and line number: d.go:23. overrides Llongfile
	Lpriority                              // the priority: Debug
	Lstd       = Ldate | Ltime | Lpriority // initial values for the standard logger
)

// sets the outputs destination for the logger.
const (
	Oconsole     = 1 << iota             // the console logger
	Odailyfile                           // the daily file logger
	Orollingfile                         // the rolling file logger
	Ostd         = Oconsole | Odailyfile // the console|daily file logger
)

const (
	logDir = "logs"
)
