package common

import (
	"time"
)

type Item interface {
	GetItem() uint
	SetID(uint)
	SetCreatedAt(time.Time)
}
