// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/web/session"
	"html/template"
)

var Conf = &conf{
	Addr:    ":9090",
	Actives: []string{profile.DEV},
	Api: &apiOptions{
		URL: "/apis",
	},
	Asset: &staticOptions{
		Enable: false,
		Dir:    "static",
		URL:    "/assets",
	},
	Developer: &staticOptions{
		Enable: false,
		Dir:    "/assets/devs",
		URL:    "/devs",
	},
	Operation: &staticOptions{
		Enable: false,
		Dir:    "/assets/oprs",
		URL:    "/oprs",
	},
	Upload: &uploadOptions{
		Enable:  false,
		Dir:     "/assets/upls",
		URL:     "/upls",
		MaxSize: 5242880,
	},
	Html: &htmlOptions{
		Enable: false,
	},
	Session: &sessionOptions{
		Enable: true,
		Addr:   ":6379",
		Options: &session.Options{
			Path:     "",
			Domain:   "",
			MaxAge:   30 * 60,
			Secure:   false,
			HttpOnly: true,
		},
	},
	Secure: &secureOptions{
		Enable:     true,
		LoginUrl:   "/login",
		SuccessUrl: "/",
		Filters: []xtype.Map{
			{"/**", "anon"},
		},
	},
	Template: &templateOptions{
		Enable:     true,
		Dir:        "templates",
		Extensions: []string{"html"},
		Funcs:      []template.FuncMap{},
		Delims: templateDelims{
			Left:  "{{",
			Right: "}}",
		},
		Reloaded: true,
	},
}

type conf struct {
	Addr      string           // the TCP network address
	Actives   []string         // Active profile
	Api       *apiOptions      // Api options
	Asset     *staticOptions   // Static resource options
	Developer *staticOptions   // Developer static resource options
	Operation *staticOptions   // Operation static resource options
	Upload    *uploadOptions   // Upload options
	Html      *htmlOptions     // Html resource options
	Session   *sessionOptions  // the session TCP network address
	Secure    *secureOptions   // url secure options
	Template  *templateOptions // template options
}

type apiOptions struct {
	URL string // APIs URL prefix
}

type staticOptions struct {
	Enable bool   // Whether service is enabled
	Dir    string // Static resource directory
	URL    string // Static resource URL prefix
}

type uploadOptions struct {
	Enable  bool   // Whether service is enabled
	Dir     string // Upload directory
	URL     string // Upload URL prefix
	MaxSize int    // Max upload size
}

type htmlOptions struct {
	Enable bool // Whether service is enabled
}

type sessionOptions struct {
	Enable bool // Whether service is enabled
	*session.Options
	Addr string
}

type secureOptions struct {
	Enable     bool // Whether service is enabled
	LoginUrl   string
	SuccessUrl string
	Filters    []xtype.Map
}

// Options is a struct for specifying configuration options for the html render
type templateOptions struct {
	// Whether service is enabled
	Enable bool
	// Directory to load templates. Default is "templates"
	Dir string
	// Extensions to parse template files from. Defaults to ["tmpl"]
	Extensions []string
	// Funcs is a slice of FuncMaps to apply to the template upon compilation.
	// This is useful for helper functions. Defaults to [].
	Funcs []template.FuncMap
	// Delims sets the action delimiters to the specified strings in the templateDelims struct.
	Delims templateDelims
	// Reloaded sets up the template for each reload
	Reloaded bool
}

// templateDelims represents a set of Left and Right delimiters for HTML template rendering
type templateDelims struct {
	// Left delimiter, defaults to {{
	Left string
	// Right delimiter, defaults to }}
	Right string
}
