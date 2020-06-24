/*
 * Copyright (c) 2020 Nguyen Trung Tin. All rights reserved.
 * Use of this source code is governed by a BSD-style
 *  license that can be found in the LICENSE file.
 */

package user

import (
	"context"

	"github.com/nguyentrungtin/go-grpc-boilerplate/pkg/connect"

	"github.com/nguyentrungtin/go-grpc-boilerplate/model"

	"gorm.io/gorm"

	"github.com/nguyentrungtin/go-grpc-boilerplate/pb"
)

type UserServer struct {
	db        *gorm.DB
	connector connect.Connector
}

// NewUserServer create new UserServer
func NewUserServer(db *gorm.DB, cnt connect.Connector) *UserServer {
	return &UserServer{
		db:        db,
		connector: cnt,
	}
}

// GetUser handle GetUserRequest and return GetUserResponse
func (server *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	id := req.GetId()
	var user model.User
	server.db.Where("id = ?", id).Find(&user)
	if user.ID == 0 {
		// @TODO: return error
		//return nil, status.Error(codes.NotFound, "User not found")

		// return example user
		user = model.User{
			ID:        id,
			FirstName: "Tin",
			LastName:  "Nguyen",
			Username:  "tin",
			Email:     "developer@example.com",
		}
	}
	res := &pb.GetUserResponse{
		User: &pb.User{
			Id:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		}}

	return res, nil
}
