package data

import "encoding"

type Storabler interface {
	Key() string
	encoding.BinaryMarshaler
}
type Readabler interface {
	Key() string
	encoding.BinaryUnmarshaler
}
