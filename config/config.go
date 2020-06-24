/*
 * Copyright (c) 2020 Nguyen Trung Tin. All rights reserved.
 * Use of this source code is governed by a BSD-style
 *  license that can be found in the LICENSE file.
 */

package config

import (
	log "github.com/sirupsen/logrus"

	"io/ioutil"

	"github.com/spf13/viper"
)

// init will be invoked automatically when config package is imported by other package
// env will be set following precedence order: default -> .env -> env on runtime (docker/k8s)
func init() {

	// load default env
	setDefaultConfig()

	// load env from .env file
	viper.SetConfigFile(".env")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		// do nothing, if cannot get from file, it'll get from ENV on runtime
	}

	// read in environment variables on runtime
	viper.AutomaticEnv()

}

// GetAll get config from ENV in .env, parse content of file if needed
// env will be set following precedence order: default -> .env -> env on runtime (docker/k8s)
// if required config is not set, fatal and exit process
func GetAll() {

	// read in environment variables on runtime
	viper.AutomaticEnv()

	if err := getJwtKeyFromFile(); err != nil {
		log.WithError(err).Fatal("Failed to load JWT key")
	}

	// these ENV must be set, if not, log with fatal
	requiredConfig("JWT_PRIVATE", "JWT_PUBLIC", "DB_NAME", "DB_USER", "DB_PASSWORD")
}

// setDefaultConfig set default ENV when no config file is set
func setDefaultConfig() {
	viper.SetDefault("APP_NAME", "GRPC-BOILERPLATE")
	viper.SetDefault("APP_ENV", "development")

	viper.SetDefault("GRPC_PORT", "9901")
	viper.SetDefault("HTTP_PORT", "9902")

	viper.SetDefault("TLS_ENABLE", false)
	viper.SetDefault("CA_CERT", "certs/grpc/ca-cert.pem")
	viper.SetDefault("SERVER_KEY", "certs/grpc/server-key.pem")
	viper.SetDefault("SERVER_CERT", "certs/grpc/server-cert.pem")
	viper.SetDefault("CLIENT_KEY", "certs/grpc/client-key.pem")
	viper.SetDefault("CLIENT_CERT", "certs/grpc/client-cert.pem")

	viper.SetDefault("JWT_PRIVATE", "certs/jwt/jwt-private.pem")
	viper.SetDefault("JWT_PUBLIC", "certs/jwt/jwt-public.pem")
	viper.SetDefault("JWT_EXP", 72)

	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "26257")
	viper.SetDefault("DB_SSLMODE", false)
	viper.SetDefault("DB_CA", "/certs/crdb/ca.crt")
	viper.SetDefault("DB_KEY", "/certs/crdb/client.dev.key")
	viper.SetDefault("DB_CERT", "/certs/crdb/client.dev.crt")
}

// getJwtKeyFromFile get jwt public/private key from local files
// In production, viper will get it from env for security reason
func getJwtKeyFromFile() error {
	private, err := ioutil.ReadFile(viper.GetString("JWT_PRIVATE"))
	if err != nil {
		return err
	}

	viper.Set("JWT_PRIVATE", string(private))

	public, err := ioutil.ReadFile(viper.GetString("JWT_PUBLIC"))
	if err != nil {
		return err
	}

	viper.Set("JWT_PUBLIC", string(public))

	return nil
}

func requiredConfig(cfgs ...string) {
	for _, cfg := range cfgs {
		if !viper.IsSet(cfg) {
			log.WithField("missing", cfg).Fatal("ENV is missing!")
		}
	}

}
