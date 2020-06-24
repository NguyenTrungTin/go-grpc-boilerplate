/*
 * Copyright (c) 2020 Nguyen Trung Tin. All rights reserved.
 * Use of this source code is governed by a BSD-style
 *  license that can be found in the LICENSE file.
 */

package greet

import (
	"context"
	"fmt"

	"github.com/nguyentrungtin/go-grpc-boilerplate/pkg/connect"

	"gorm.io/gorm"

	"github.com/nguyentrungtin/go-grpc-boilerplate/pb"
)

type GreetServer struct {
	db        *gorm.DB
	connector connect.Connector
}

// NewGreetServer create new GreetServer
func NewGreetServer(db *gorm.DB, cnt connect.Connector) *GreetServer {
	return &GreetServer{
		db:        db,
		connector: cnt,
	}
}

// Greet handle GreetRequest and return GreetResponse
func (server *GreetServer) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	firstName := req.Name.FirstName
	lastName := req.Name.LastName
	result := fmt.Sprintf("Hi, %s %s", firstName, lastName)
	res := &pb.GreetResponse{
		Result: result,
	}
	return res, nil
}
