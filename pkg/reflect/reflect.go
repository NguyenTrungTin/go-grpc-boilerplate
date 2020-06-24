/*
 * Copyright (c) 2020 Nguyen Trung Tin. All rights reserved.
 * Use of this source code is governed by a BSD-style
 *  license that can be found in the LICENSE file.
 */

package reflect

import (
	"context"

	gr "github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

type Reflector interface {
	ListServices() ([]string, error)
}

type ClientReflect struct {
	Conn *grpc.ClientConn
}

// NewClientReflect create new ClientReflect
func NewClientReflect(c *grpc.ClientConn) *ClientReflect {
	return &ClientReflect{
		Conn: c,
	}
}

// LisServices list all gRPC services
func (cr *ClientReflect) ListServices() ([]string, error) {
	rf := gr.NewClient(context.Background(), grpc_reflection_v1alpha.NewServerReflectionClient(cr.Conn))
	svcs, err := rf.ListServices()
	if err != nil {
		return nil, err
	}
	return svcs, nil
}

type ServerReflect struct {
	svr *grpc.Server
}

// NewServerReflect create ServerReflect
func NewServerReflect(svr *grpc.Server) *ServerReflect {
	return &ServerReflect{
		svr: svr,
	}
}

// LisServices list all gRPC services
func (sr *ServerReflect) ListServices() ([]string, error) {
	sds, err := gr.LoadServiceDescriptors(sr.svr)
	if err != nil {
		return nil, err
	}

	var svcs []string
	for _, sd := range sds {
		svcs = append(svcs, sd.GetName())
	}

	return svcs, nil
}
