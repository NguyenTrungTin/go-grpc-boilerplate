/*
 * Copyright (c) 2020 Nguyen Trung Tin. All rights reserved.
 * Use of this source code is governed by a BSD-style
 *  license that can be found in the LICENSE file.
 */

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"

	"github.com/nguyentrungtin/go-grpc-boilerplate/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-grpc",
	Short: "go-grpc is a cli to interact with gRPC service",
	Long:  `go-grpc is a boilerplate to start development with gRPC`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
// if you need to init some func before each commands, you can implement in here
func initConfig() {
	// get all config, fatal if required config is not set
	config.GetAll()

	// init logrus setting
	initLog()
}

func initLog() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	if viper.Get("APP_ENV") == "development" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}
