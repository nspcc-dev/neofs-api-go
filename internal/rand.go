package internal

import (
	"math/rand"
	"time"
)

// RandUint32 returns random uint32 value [0, max).
func RandUint32(max uint32) uint32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Uint32() % max
}
