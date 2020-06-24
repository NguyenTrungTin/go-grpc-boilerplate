/*
 * Copyright (c) 2020 Nguyen Trung Tin. All rights reserved.
 * Use of this source code is governed by a BSD-style
 *  license that can be found in the LICENSE file.
 */

package model

type User struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	FirstName string `gorm:"not null" json:"firstname" validate:"required" conform:"trim"`
	LastName  string `gorm:"not null" json:"lastname" validate:"required" conform:"trim"`
	Username  string `gorm:"not null;unique" json:"username" validate:"required" conform:"trim"`
	Email     string `gorm:"not null" json:"email" validate:"required,email" conform:"trim"`
}
