/*
 * Copyright (c) 2020 Nguyen Trung Tin. All rights reserved.
 * Use of this source code is governed by a BSD-style
 *  license that can be found in the LICENSE file.
 */

package auth

import (
	log "github.com/sirupsen/logrus"

	"github.com/casbin/casbin/v2"
)

type RBACRequest struct {
	id      int64
	domain  string
	service string
	action  string
}

// RBAC (role base access control)
func RBAC(r *RBACRequest) error {
	_, err := casbin.NewEnforcer("config/rbac.conf", "config/policy.csv")
	if err != nil {
		log.WithError(err).Error("Failed to load RBAC config")
	}

	// @TODO: implement RBAC check
	return nil

}
