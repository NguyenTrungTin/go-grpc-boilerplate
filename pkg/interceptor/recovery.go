/*
 * Copyright (c) 2020 Nguyen Trung Tin. All rights reserved.
 * Use of this source code is governed by a BSD-style
 *  license that can be found in the LICENSE file.
 */

package interceptor

import (
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecoveryInterceptor struct {
}

// NewRecoveryInterceptor create new RecoveryInterceptor
func NewRecoveryInterceptor() *RecoveryInterceptor {
	return &RecoveryInterceptor{}
}

// Unary is unary interceptor to recovery process from panic
func (interceptor *RecoveryInterceptor) Unary() grpc.UnaryServerInterceptor {
	opts := recoveryOpts()
	return grpc_recovery.UnaryServerInterceptor(opts...)
}

// Unary is stream interceptor to recovery process from panic
func (interceptor *RecoveryInterceptor) Stream() grpc.StreamServerInterceptor {
	opts := recoveryOpts()
	return grpc_recovery.StreamServerInterceptor(opts...)
}

func recoveryOpts() []grpc_recovery.Option {
	var recoveryFunc grpc_recovery.RecoveryHandlerFunc = func(p interface{}) (err error) {
		log.Error(p)
		return status.Error(codes.Internal, "internal server error")
	}
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(recoveryFunc),
	}
	return opts
}
