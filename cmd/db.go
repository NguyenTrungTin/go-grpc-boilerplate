/*
 * Copyright (c) 2020 Nguyen Trung Tin. All rights reserved.
 * Use of this source code is governed by a BSD-style
 *  license that can be found in the LICENSE file.
 */

package cmd

import (
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Interact with database",
	Long:  `Interact with database`,
}

func init() {
	rootCmd.AddCommand(dbCmd)
}
