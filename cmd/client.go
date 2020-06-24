/*
 * Copyright (c) 2020 Nguyen Trung Tin. All rights reserved.
 * Use of this source code is governed by a BSD-style
 *  license that can be found in the LICENSE file.
 */

package cmd

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"google.golang.org/grpc/credentials"

	"github.com/nguyentrungtin/go-grpc-boilerplate/pb"
	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "start gRPC client",
	Long:  `start gRPC client`,
	Run: func(cmd *cobra.Command, args []string) {
		Client()
	},
}

var (
	grpcAddress     string
	tlsClientEnable bool = viper.GetBool("TLS_ENABLE")
)

func init() {
	rootCmd.AddCommand(clientCmd)

	clientCmd.Flags().StringVarP(&grpcAddress, "address", "a", viper.GetString("GRPC_PORT"), "run gRPC client and dial to gRPC server address")

	clientCmd.Flags().BoolVarP(&tlsClientEnable, "tls", "t", viper.GetBool("TLS_ENABLE"), "enable Client TLS")
}

func Client() {
	svrAddr := fmt.Sprintf("0.0.0.0:%s", grpcAddress)

	// TLS
	var opts grpc.DialOption
	if tlsClientEnable {
		tlsCredentials, err := loadClientTLSCredentials()
		if err != nil {
			log.WithError(err).Fatal("cannot load TLS credentials")
		}

		opts = grpc.WithTransportCredentials(tlsCredentials)
	} else {
		opts = grpc.WithInsecure()
	}
	log.WithField("tls", tlsClientEnable).Info("gRPC client tls mode")

	// gRPC client
	cc, err := grpc.Dial(svrAddr, opts)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect gRPC server")
	}
	defer cc.Close()

	c := pb.NewGreetServiceClient(cc)

	GetUser(c)
}

func loadClientTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	caCert, err := ioutil.ReadFile(viper.GetString("CA_CERT"))
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair(viper.GetString("CLIENT_CERT"), viper.GetString("CLIENT_KEY"))
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

func GetUser(c pb.GreetServiceClient) {
	req := &pb.GreetRequest{
		Name: &pb.Name{
			FirstName: "Tin",
			LastName:  "Nguyen",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.WithError(err).Error("Error while calling Greet GRPC")
	}
	log.WithField("result", res.Result).Info("Response from Greet: ")
}
