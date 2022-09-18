// -*- Mode: Go; indent-tabs-mode: t -*-
//go:build !linux
// +build !linux

/*
 * Copyright (C) 2019-2022 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package preseed

import (
	"errors"
)

var preseedNotAvailableError = errors.New("preseed mode not available for systems other than linux")

func Classic(chrootDir string) error {
	return preseedNotAvailableError
}

func Core20(opts *CorePreseedOpts) error {
	return preseedNotAvailableError
}
