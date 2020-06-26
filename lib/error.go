/*
 * Copyright (c) 2020 Nguyen Trung Tin. All rights reserved.
 * Use of this source code is governed by a BSD-style
 *  license that can be found in the LICENSE file.
 */

package lib

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleErr check error, if err not nil, log to stdout
func LogErr(err error) {
	if err != nil {
		log.Error(err)
	}
}

// Must check error, if err not nil, log and fatal
func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ContextErr check context error and return status and code base on current error
func ContextErr(ctx context.Context) error {
	LogErr(ctx.Err())

	switch ctx.Err() {
	case context.Canceled:
		return status.Error(codes.Canceled, "request is canceled")
	case context.DeadlineExceeded:
		return status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	default:
		return nil
	}
}
