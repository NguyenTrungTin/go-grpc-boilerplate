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
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed base data into database",
	Long:  `Seed base data into database`,
	Run: func(cmd *cobra.Command, args []string) {
		Seed()
	},
}

func init() {
	dbCmd.AddCommand(seedCmd)
}

func Seed() {
	db.Seed()
}
