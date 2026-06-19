package main

// LeetCode #2166: Design Bitset
// https://leetcode.com/problems/design-bitset/
// Difficulty: Medium
// Time: O(1) per operation | Space: O(n)

import (
	"fmt"
	"strings"
)

type Bitset struct {
	bits  []bool
	flip  bool
	ones  int
	size  int
}

func Constructor(size int) Bitset {
	return Bitset{
		bits: make([]bool, size),
		size: size,
	}
}

func (this *Bitset) Fix(idx int) {
	if this.flip {
		if this.bits[idx] {
			this.bits[idx] = false
			this.ones++
		}
	} else {
		if !this.bits[idx] {
			this.bits[idx] = true
			this.ones++
		}
	}
}

func (this *Bitset) Unfix(idx int) {
	if this.flip {
		if !this.bits[idx] {
			this.bits[idx] = true
			this.ones--
		}
	} else {
		if this.bits[idx] {
			this.bits[idx] = false
			this.ones--
		}
	}
}

func (this *Bitset) Flip() {
	this.flip = !this.flip
	this.ones = this.size - this.ones
}

func (this *Bitset) All() bool {
	return this.ones == this.size
}

func (this *Bitset) One() bool {
	return this.ones > 0
}

func (this *Bitset) Count() int {
	return this.ones
}

func (this *Bitset) ToString() string {
	var sb strings.Builder
	for i := 0; i < this.size; i++ {
		bit := this.bits[i]
		if this.flip {
			bit = !bit
		}
		if bit {
			sb.WriteByte('1')
		} else {
			sb.WriteByte('0')
		}
	}
	return sb.String()
}

func main() {
	// Test case 1
	bs := Constructor(5)
	bs.Fix(3)
	bs.Fix(1)
	fmt.Println("Test 1 All:", bs.All())
	fmt.Println("Test 1 toString:", bs.ToString())
	// Expected: false, "01010"

	// Test case 2
	bs2 := Constructor(5)
	bs2.Fix(0)
	bs2.Fix(1)
	bs2.Flip()
	fmt.Println("Test 2 All:", bs2.All())
	fmt.Println("Test 2 toString:", bs2.ToString())
	// Expected: false, "00111"

	// Test case 3
	bs3 := Constructor(2)
	bs3.Flip()
	fmt.Println("Test 3 All:", bs3.All())
	fmt.Println("Test 3 toString:", bs3.ToString())
	// Expected: true, "11"
}
