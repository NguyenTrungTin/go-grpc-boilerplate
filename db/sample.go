/*
 * Copyright (c) 2020 Nguyen Trung Tin. All rights reserved.
 * Use of this source code is governed by a BSD-style
 *  license that can be found in the LICENSE file.
 */

package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Seed import sample data into database
func Sample() {
	sslmode := viper.GetBool("DB_SSLMODE")
	var dsn string
	if sslmode {
		dsn = fmt.Sprintf("cockroachdb://%s:%s@%s:%s/%s?sslmode=require&sslrootcert=%s&sslkey=%s&sslcert=%s", viper.Get("DB_USER"), viper.Get("DB_USER"), viper.Get("DB_HOST"), viper.Get("DB_PORT"), viper.Get("DB_NAME"), viper.Get("DB_CA"), viper.Get("DB_KEY"), viper.Get("DB_CERT"))
	} else {
		dsn = fmt.Sprintf("cockroachdb://%s:%s@%s:%s/%s?sslmode=disable", viper.Get("DB_USER"), viper.Get("DB_USER"), viper.Get("DB_HOST"), viper.Get("DB_PORT"), viper.Get("DB_NAME"))
	}

	m, err := migrate.New(
		"file://db/samples",
		dsn,
	)
	if err != nil {
		log.WithError(err).Error("Error happen when creation migration instance")
	}
	if err := m.Up(); err != nil {
		log.Info(err)
	}
}
