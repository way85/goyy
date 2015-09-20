// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package webs

import (
	"bytes"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"net/http"
	"net/url"
)

// Form contains the parsed form data, including both the URL
// field's query parameters and the POST or PUT form data.
func ToParams(values url.Values, prefix ...string) map[string]string {
	result := make(map[string]string, 0)
	for k, v := range values {
		if _, ok := strings.HasAnyPrefix(k, prefix...); ok {
			result[k] = strings.JoinIgnoreBlank(v, ",")
		}
	}
	return result
}

// Form contains the parsed form data, including both the URL
// field's query parameters and the POST or PUT form data.
func Params(req *http.Request, prefix ...string) (map[string]string, error) {
	vs, err := Values(req)
	if err != nil {
		return nil, err
	}
	result := make(map[string]string, 0)
	for k, v := range vs {
		if _, ok := strings.HasAnyPrefix(k, prefix...); ok {
			result[k] = strings.JoinIgnoreBlank(v, ",")
		}
	}
	return result, nil
}

// Form contains the parsed form data, including both the URL
// field's query parameters and the POST or PUT form data.
func Values(req *http.Request) (values url.Values, err error) {
	if req.Method == "GET" {
		values = req.Form
	} else {
		err := req.ParseForm()
		if err != nil {
			return nil, err
		}
		values = req.PostForm
	}
	if values == nil {
		values = url.Values{}
	}
	u, err := url.Parse(req.RequestURI)
	if err != nil {
		return nil, err
	}
	if strings.IsNotBlank(u.RawQuery) {
		vs, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			return nil, err
		}
		for k, v := range vs {
			for _, s := range v {
				if strings.IsNotBlank(s) {
					values.Add(k, s)
				}
			}
		}
	}
	return
}

// RemoteIP returns the Remote IP
func RemoteIP(req *http.Request) string {
	unknown := "unknown"
	ip := req.Header.Get("x-forwarded-for")
	if strings.IsBlank(ip) || unknown == strings.ToLower(ip) {
		ip = req.Header.Get("X-Forwarded-For")
	}
	if strings.IsBlank(ip) || unknown == strings.ToLower(ip) {
		ip = req.Header.Get("Proxy-Client-IP")
	}
	if strings.IsBlank(ip) || unknown == strings.ToLower(ip) {
		ip = req.Header.Get("WL-Proxy-Client-IP")
	}
	if strings.IsBlank(ip) || unknown == strings.ToLower(ip) {
		ip = req.Header.Get("HTTP_CLIENT_IP")
	}
	if strings.IsBlank(ip) || unknown == strings.ToLower(ip) {
		ip = req.Header.Get("HTTP_X_FORWARDED_FOR")
	}
	if strings.IsBlank(ip) || unknown == strings.ToLower(ip) {
		ip = req.RemoteAddr
	}
	if strings.IsNotBlank(ip) && strings.Index(ip, ",") != -1 {
		ips := strings.Split(ip, ",")
		for i := 0; i < len(ips); i++ {
			if strings.IsNotBlank(ips[i]) && unknown != strings.ToLower(ip) {
				ip = ips[i]
				break
			}
		}
		if "0:0:0:0:0:0:1" == ip {
			ip = "127.0.0.1"
		}
	}
	if strings.IsNotBlank(ip) && strings.Index(ip, ":") != -1 {
		ip = strings.Before(ip, ":")
	}
	return ip
}

// ParseQuery returns the url.Query by the url.Values.
func ParseQuery(params url.Values) string {
	if params == nil {
		return ""
	}
	var b bytes.Buffer
	i := 0
	for k, v := range params {
		param := strings.JoinIgnoreBlank(v, ",")
		if strings.IsNotBlank(param) {
			if i > 0 {
				b.WriteString("&")
			}
			b.WriteString(k + "=" + param)
			i++
		}
	}
	return b.String()
}
