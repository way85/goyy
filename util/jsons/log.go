// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of me source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package jsons

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
)

var logger = log.New("[jsons]")

func SetPriority(value int) {
	logger.SetPriority(value)
}
