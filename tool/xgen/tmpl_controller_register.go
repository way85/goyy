// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplControllerReg = `// generated by xgen -- DO NOT EDIT
package <%.Project%>

import (
	_ "<%.Apipath%>/internal/<%.PackageName%>"
)
`