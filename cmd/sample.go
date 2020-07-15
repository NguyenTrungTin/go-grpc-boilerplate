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
var sampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "Load samples data into database",
	Long:  `Load samples data into database`,
	Run: func(cmd *cobra.Command, args []string) {
		Sample()
	},
}

func init() {
	dbCmd.AddCommand(sampleCmd)
}

func Sample() {
	db.Sample()
}
