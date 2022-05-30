package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Uint32 returns random uint32 value [0, max).
func Uint32(max uint32) uint32 {
	return rand.Uint32() % max
}
