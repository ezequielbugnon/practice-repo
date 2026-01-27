package bitmanipulation

import "math/bits"

func MSB(n uint) uint {
	if n == 0 {
		return 0
	}
	return 1 << (bits.Len(n) - 1)
}
