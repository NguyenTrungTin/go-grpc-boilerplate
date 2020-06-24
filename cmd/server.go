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
	"net"
	"net/http"

	"github.com/nguyentrungtin/go-grpc-boilerplate/pkg/connect"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc/reflection"

	"github.com/nguyentrungtin/go-grpc-boilerplate/pkg/interceptor"

	"github.com/nguyentrungtin/go-grpc-boilerplate/pkg/services/greet"
	"github.com/nguyentrungtin/go-grpc-boilerplate/pkg/services/user"

	"github.com/nguyentrungtin/go-grpc-boilerplate/db"

	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/nguyentrungtin/go-grpc-boilerplate/pb"
	"github.com/rs/cors"
	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start the gRPC and gRPC gateway server",
	Long:  `start the gRPC and REST API server`,
	Run: func(cmd *cobra.Command, args []string) {
		Server()
	},
}

var (
	grpcPort  string
	httpPort  string
	tlsEnable bool
)

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringVarP(&grpcPort, "port", "p", viper.GetString("GRPC_PORT"), "run gRPC server on port")

	serverCmd.Flags().StringVarP(&httpPort, "http", "r", viper.GetString("HTTP_PORT"), "run HTTP server on port")

	serverCmd.Flags().BoolVarP(&tlsEnable, "tls", "t", viper.GetBool("TLS_ENABLE"), "enable TLS")
}

// Server start the server
func Server() {
	svr := GRPC()
	HTTP(svr)
}

// GRPC start gRPC server
func GRPC() *grpc.Server {
	svrAddr := fmt.Sprintf("0.0.0.0:%s", grpcPort)
	lis, err := net.Listen("tcp", svrAddr)
	if err != nil {
		log.WithError(err).Fatal("Failed to start gRPC server")
	}

	// connect to database
	db := db.Connect()

	// Create connector to connect with other gRPC server
	connector := createConnector()

	// interceptor
	auth := interceptor.NewAuthInterceptor(db, connector)
	recovery := interceptor.NewRecoveryInterceptor()

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(recovery.Unary(), auth.Unary()),
		grpc.ChainStreamInterceptor(recovery.Stream(), auth.Stream()),
	}

	// TLS
	if tlsEnable {
		tlsCredentials, err := loadServerTLSCredentials()
		if err != nil {
			log.WithError(err).Fatal("cannot load TLS credentials")
		}

		opts = append(opts, grpc.Creds(tlsCredentials))
	}
	log.WithField("tls", tlsEnable).Info("gRPC server tls mode")

	// new gRPC server
	svr := grpc.NewServer(opts...)

	// register services
	pb.RegisterGreetServiceServer(svr, greet.NewGreetServer(db, connector))
	pb.RegisterUserServiceServer(svr, user.NewUserServer(db, connector))

	// gRPC reflection
	reflection.Register(svr)

	// Start gRPC server
	// must start gRPC server on goroutine to make gRPC and HTTP server running concurrently
	go func() {
		log.WithField("address", svrAddr).Info("gRPC server started")
		if err := svr.Serve(lis); err != nil {
			log.WithError(err).Fatal("Failed to start gRPC server")
		}
	}()

	return svr
}

// HTTP start http server
func HTTP(svr *grpc.Server) {
	// gRPC-web wrapper (call gRPC on browser)
	wmux := grpcWeb(svr)

	// gRPC-gateway (REST API support)
	rmux := restGateway()

	// apiDocs (serve swagger API docs)
	dmux := apiDocs()

	// Root serveMux
	mux := http.NewServeMux()
	mux.Handle("/grpc/", http.StripPrefix("/grpc", wmux))
	mux.Handle("/api/", http.StripPrefix("/api", rmux))
	mux.Handle("/docs/", http.StripPrefix("/docs/api", dmux))

	// Start HTTP server
	httpAddr := fmt.Sprintf("0.0.0.0:%s", httpPort)
	rootHandler := cors.Default().Handler(mux)
	log.WithField("address", httpAddr).Info("HTTP server started")
	log.WithError(http.ListenAndServe(httpAddr, rootHandler)).Fatal("Failed to start HTTP server")
}

// grpcWeb is wrapper and be a proxy for gRPC server to support gRPC on browser
func grpcWeb(svr *grpc.Server) http.Handler {
	wrappedServer := grpcweb.WrapServer(svr)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}
	wmux := http.NewServeMux()
	wmux.HandleFunc("/", handler)

	return wmux
}

// restGateway use grpc-gateway to support REST API base on gRPC definition
func restGateway() http.Handler {
	rmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	ctx := context.Background()

	grpcAddr := fmt.Sprintf("0.0.0.0:%s", grpcPort)
	err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, rmux, grpcAddr, opts)
	if err != nil {
		log.WithError(err).Error("Failed to start gRPC REST gateway")
	}

	return rmux
}

// apiDocs serve api docs (swagger API,...)
func apiDocs() http.Handler {
	dmux := http.NewServeMux()
	fs := http.FileServer(http.Dir("docs/api"))
	dmux.Handle("/", fs)

	return dmux
}

func loadServerTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed client's certificate
	caCert, err := ioutil.ReadFile(viper.GetString("CA_CERT"))
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(viper.GetString("SERVER_CERT"), viper.GetString("SERVER_KEY"))
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}

func createConnector() connect.Connector {
	connector := connect.Connector{}

	tlsOpts := connect.TLSOptions{
		CACert:     viper.GetString("CA_CERT"),
		ClientKey:  viper.GetString("CLIENT_KEY"),
		ClientCert: viper.GetString("CLIENT_CERT"),
	}

	paveStorage := connect.Connection{
		Name:       "ANOTHER-GRPC",
		Address:    viper.GetString("ANOTHER_GRPC"),
		TLS:        viper.GetBool("TLS_ENABLE"),
		TLSOptions: &tlsOpts,
	}

	return connector.Add(&paveStorage)
}
