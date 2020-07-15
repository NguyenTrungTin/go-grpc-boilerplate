/*
 * Copyright (c) 2020 Nguyen Trung Tin. All rights reserved.
 * Use of this source code is governed by a BSD-style
 *  license that can be found in the LICENSE file.
 */

package cmd

import (
	"github.com/nguyentrungtin/go-grpc-boilerplate/db"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate table schema into database",
	Long:  `Migrate table schema into database`,
	Run: func(cmd *cobra.Command, args []string) {
		Migrate()
	},
}

func init() {
	dbCmd.AddCommand(migrateCmd)
}

func Migrate() {
	db.Migrate()
}
