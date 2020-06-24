/*
 * Copyright (c) 2020 Nguyen Trung Tin. All rights reserved.
 * Use of this source code is governed by a BSD-style
 *  license that can be found in the LICENSE file.
 */

package lib

import (
	log "github.com/sirupsen/logrus"
)

// HandleErr check error, if err not nil, log to stdout
func HandleErr(err error) {
	if err != nil {
		log.Error(err)
	}
}

// MustErr check error, if err not nil, log and fatal
func MustErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
