package cpuload

import (
	"crypto/rand"
	"fmt"
	mathrand "math/rand"
)

func RunCPULoad(minimum int, variance int) {
	min := minimum
	max := min + variance
	primelen := mathrand.Intn(max-min+1) + min
	prime, _ := rand.Prime(rand.Reader, primelen)
	fmt.Println(prime)
}
