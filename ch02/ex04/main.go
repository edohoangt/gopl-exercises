package main

import (
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i / 2] + byte(i & 1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] + // byte(n) truncates n by doing n % 256
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	count := 0
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i * 8))])
	}
	return count
}

func PopCountShift(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		if x & 1 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}

func main() {
	testNum := 124
	fmt.Printf("PopCount - Num set bits of %d: %d\n", testNum, PopCount(uint64(testNum)))
	fmt.Printf("PopCountLoop - Num set bits of %d: %d\n", testNum, PopCountLoop(uint64(testNum)))
	fmt.Printf("PopCountShift - Num set bits of %d: %d\n", testNum, PopCountShift(uint64(testNum)))
}