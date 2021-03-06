// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

package gcimporter

import "go/types"

type types_Alias => types.Alias

func types_NewAlias => types.NewAlias

// TODO(gri) Consider exporting this functionality from go/types (issue 17730).
func original(obj types.Object) types.Object {
	if alias, ok := obj.(*types.Alias); ok {
		return alias.Orig()
	}
	return obj
}

const testfile = "exports18.go"
