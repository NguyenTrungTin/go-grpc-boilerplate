/*
 * Copyright (c) 2020 Nguyen Trung Tin. All rights reserved.
 * Use of this source code is governed by a BSD-style
 *  license that can be found in the LICENSE file.
 */

package connect

import (
	"fmt"

	"google.golang.org/grpc"
)

type Connector map[string]*Connection

// Add add connection to connector
func (c Connector) Add(connections ...*Connection) Connector {
	for _, connection := range connections {
		c[connection.Name] = connection
	}
	return c
}

// Get get connection in the connector by name
func (c Connector) Get(name string) (*Connection, error) {
	connection, ok := c[name]
	if !ok {
		return nil, fmt.Errorf("connection name %s not found", name)
	}
	return connection, nil
}

// Delete delete connection in the connector
func (c Connector) Delete(connections ...*Connection) Connector {
	for _, connection := range connections {
		delete(c, connection.Name)
	}
	return c
}

// Connect connect to the gRPC server
// Remember to defer Close() to avoid memory leak
func (c Connector) Connect(name string) (*grpc.ClientConn, error) {
	connection, err := c.Get(name)
	if err != nil {
		return nil, err
	}
	cnn, err := connection.Connect()
	if err != nil {
		return nil, err
	}
	return cnn.Conn, nil
}
