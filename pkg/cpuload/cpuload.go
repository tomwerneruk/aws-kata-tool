package cpuload

import (
	"crypto/rand"
	"math/big"
	mathrand "math/rand"
)

func RunCPULoad(minimum int, variance int) *big.Int {
	min := minimum
	max := min + variance
	primelen := mathrand.Intn(max-min+1) + min
	prime, _ := rand.Prime(rand.Reader, primelen)
	return prime
}
