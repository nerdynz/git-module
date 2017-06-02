// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

type ObjectType string

const (
	OBJECT_COMMIT ObjectType = "commit"
	OBJECT_TREE   ObjectType = "tree"
	OBJECT_BLOB   ObjectType = "blob"
	OBJECT_TAG    ObjectType = "tag"
)

func (ot ObjectType) String() string {
	switch ot {
	case OBJECT_COMMIT:
		return "commit"
	case OBJECT_TREE:
		return "tree"
	case OBJECT_BLOB:
		return "blob"
	case OBJECT_TAG:
		return "tag"
	default:
		return ""
	}
}
