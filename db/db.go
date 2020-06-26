/*
 * Copyright (c) 2020 Nguyen Trung Tin. All rights reserved.
 * Use of this source code is governed by a BSD-style
 *  license that can be found in the LICENSE file.
 */

package db

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Init is used to init the connection to the database
func Connect() *gorm.DB {
	sslmode := viper.GetBool("DB_SSLMODE")
	var dsn string
	if sslmode {
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=require sslrootcert=certs/crdb/ca.crt sslkey=certs/crdb/client.pave.key sslcert=certs/crdb/client.pave.crt", viper.Get("DB_HOST"), viper.Get("DB_PORT"), viper.Get("DB_USER"), viper.Get("DB_NAME"), viper.Get("DB_PASSWORD"))
	} else {
		dsn = fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable", viper.Get("DB_HOST"), viper.Get("DB_PORT"), viper.Get("DB_NAME"))
	}

	// CockroachDB use postgres dialect
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.WithError(err).Fatal("Failed to connect to the database")
	}
	log.Info("Connected to database successfully")

	return db
}
