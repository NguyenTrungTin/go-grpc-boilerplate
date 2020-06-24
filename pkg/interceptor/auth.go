/*
 * Copyright (c) 2020 Nguyen Trung Tin. All rights reserved.
 * Use of this source code is governed by a BSD-style
 *  license that can be found in the LICENSE file.
 */

package interceptor

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/nguyentrungtin/go-grpc-boilerplate/pkg/connect"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type AuthInterceptor struct {
	db        *gorm.DB
	connector connect.Connector
}

// NewAuthInterceptor create new AuthInterceptor
func NewAuthInterceptor(db *gorm.DB, cnt connect.Connector) *AuthInterceptor {
	return &AuthInterceptor{
		db:        db,
		connector: cnt,
	}
}

// Unary is unary interceptor which validate token and RBAC permissions
func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.WithField("method", info.FullMethod).Info("Unary interceptor")

		// get authorization token
		token, err := interceptor.GetToken(ctx, info.FullMethod)
		if err != nil {
			// @TODO: return error
			//return nil, err
			// temporary pass for demo
			return handler(ctx, req)

		}
		// @TODO: use token to check auth instead of log to stdout
		log.WithField("token", token).Info("TOKEN")

		cnn, err := interceptor.connector.Connect("ANOTHER-GRPC")
		defer cnn.Close()
		// @TODO: Call gRPC to PAVE-ID to check auth

		return handler(ctx, req)
	}
}

// Stream is stream interceptor which validate token and RBAC permissions
func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.WithField("method", info.FullMethod).Info("Stream interceptor")

		// get authorization token
		token, err := interceptor.GetToken(stream.Context(), info.FullMethod)
		if err != nil {
			// @TODO: return error
			//return err
			// temporary pass for demo
			return handler(srv, stream)
		}
		// @TODO: use token to check auth instead of log to stdout
		log.WithField("token", token).Info("TOKEN")

		cnn, err := interceptor.connector.Connect("ANOTHER-GRPC")
		defer cnn.Close()
		// @TODO: Call gRPC to PAVE-ID to check auth

		return handler(srv, stream)
	}
}

// GetToken get jwt token on metadata
func (interceptor *AuthInterceptor) GetToken(ctx context.Context, method string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	token := values[0]

	return token, nil
}
