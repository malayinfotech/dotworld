// Copyright (C) 2021 Storx  Labs, Inc.
// See LICENSE for copying information.

package reference

//go:generate go run gen.go

import (
	"dotworld"
)

// WorldMap is a good default world map to use.
func WorldMap() *dotworld.Map {
	return worldMap.Copy()
}
