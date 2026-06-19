package main

// LeetCode #1622: Fancy Sequence
// https://leetcode.com/problems/fancy-sequence/
// Difficulty: Hard

import "fmt"

const MOD = 1000000007

// modInv computes modular inverse using Fermat's little theorem (MOD is prime)
func modInv(a int) int {
	return modPow(a, MOD-2)
}

func modPow(a, b int) int {
	res := 1
	a %= MOD
	for b > 0 {
		if b&1 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		b >>= 1
	}
	return res
}

type Fancy struct {
	arr []int // stored base values (normalized)
	mul int   // global multiplier
	add int   // global adder
}

func Constructor() Fancy {
	return Fancy{
		arr: []int{},
		mul: 1,
		add: 0,
	}
}

// Append(val): append val to sequence
func (f *Fancy) Append(val int) {
	// If mul == 0, all existing elements have value = f.add.
	// We need to normalize: set all existing stored values to 0
	// and set mul=1, add=f.add (unchanged).
	if f.mul == 0 {
		for i := range f.arr {
			f.arr[i] = 0
		}
		f.mul = 1
		// f.add stays the same
	}

	// Normalize: find stored_val such that stored_val * mul + add = val
	// stored_val = (val - add) * inv(mul) (mod MOD)
	v := (val - f.add + MOD) % MOD
	v = v * modInv(f.mul) % MOD
	f.arr = append(f.arr, v)
}

// AddAll(inc): add inc to every element
func (f *Fancy) AddAll(inc int) {
	f.add = (f.add + inc) % MOD
}

// MultAll(m): multiply every element by m
func (f *Fancy) MultAll(m int) {
	f.mul = f.mul * m % MOD
	f.add = f.add * m % MOD
}

// GetIndex(idx): return value at index, -1 if out of bounds
func (f *Fancy) GetIndex(idx int) int {
	if idx >= len(f.arr) {
		return -1
	}
	if f.mul == 0 {
		return f.add
	}
	return (f.arr[idx]*f.mul + f.add) % MOD
}

// FancySequence tests the Fancy struct operations and returns results
func FancySequence() []int {
	f := Constructor()
	results := []int{}

	f.Append(2)      // seq = [2]
	f.AddAll(3)      // seq = [5]
	f.Append(7)      // seq = [5, 7]
	f.MultAll(2)     // seq = [10, 14]
	results = append(results, f.GetIndex(0)) // 10
	f.AddAll(3)      // seq = [13, 17]
	f.Append(10)     // seq = [13, 17, 10]
	f.MultAll(2)     // seq = [26, 34, 20]
	results = append(results, f.GetIndex(0)) // 26
	results = append(results, f.GetIndex(1)) // 34
	results = append(results, f.GetIndex(2)) // 20

	return results
}

func main() {
	// Test case: sequence of operations
	f := Constructor()
	f.Append(2)
	f.AddAll(3)
	f.Append(7)
	f.MultAll(2)
	fmt.Printf("GetIndex(0) = %d (expected 10)\n", f.GetIndex(0))
	f.AddAll(3)
	f.Append(10)
	f.MultAll(2)
	fmt.Printf("GetIndex(0) = %d (expected 26)\n", f.GetIndex(0))
	fmt.Printf("GetIndex(1) = %d (expected 34)\n", f.GetIndex(1))
	fmt.Printf("GetIndex(2) = %d (expected 20)\n", f.GetIndex(2))

	// Test case: multAll(0) then append
	f2 := Constructor()
	f2.Append(5)    // [5]
	f2.MultAll(0)   // [0]
	f2.Append(3)    // [0, 3]
	f2.AddAll(2)    // [2, 5]
	fmt.Printf("After multAll(0): GetIndex(0) = %d (expected 2)\n", f2.GetIndex(0))
	fmt.Printf("After multAll(0): GetIndex(1) = %d (expected 5)\n", f2.GetIndex(1))

	// Test case: empty
	f3 := Constructor()
	fmt.Printf("Empty getIndex(0) = %d (expected -1)\n", f3.GetIndex(0))
}
