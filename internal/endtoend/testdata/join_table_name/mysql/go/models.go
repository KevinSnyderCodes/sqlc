// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package querytest

import ()

type Bar struct {
	ID uint64
}

type Foo struct {
	ID  uint64
	Bar uint64
}
