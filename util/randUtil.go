package util

import "math/rand"

func initRand(u *Utils) *rand.Rand {
	return rand.New(rand.NewSource(u.GetCurTs()))
}
