/*
 * Copyright (c) 2020 Nguyen Trung Tin. All rights reserved.
 * Use of this source code is governed by a BSD-style
 *  license that can be found in the LICENSE file.
 */

package connect

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Connection struct {
	Name       string
	Address    string
	TLS        bool
	TLSOptions *TLSOptions
	Conn       *grpc.ClientConn
}

type TLSOptions struct {
	CACert     string
	ClientCert string
	ClientKey  string
}

// Connect connect to gRPC server
// Remember to defer.Close() to avoid memory leak
func (c *Connection) Connect() (*Connection, error) {
	// TLS
	var options grpc.DialOption
	if c.TLS {
		tlsCredentials, err := loadClientTLSCredentials(c.TLSOptions)
		if err != nil {
			return nil, fmt.Errorf("cannot load TLS credentials")
		}

		options = grpc.WithTransportCredentials(tlsCredentials)
	} else {
		options = grpc.WithInsecure()
	}

	// gRPC client
	cc, err := grpc.Dial(c.Address, options)
	if err != nil {
		return nil, fmt.Errorf("failed to connect gRPC server")
	}

	c.Conn = cc

	return c, nil
}

// Close close the gRPC connection
func (c *Connection) Close() error {
	return c.Conn.Close()
}

func loadClientTLSCredentials(opts *TLSOptions) (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	caCert, err := ioutil.ReadFile(opts.CACert)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair(opts.ClientCert, opts.ClientKey)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}
