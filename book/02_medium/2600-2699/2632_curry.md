# 2632 — Curry

## Deskripsi

**Soal:** [2632. Curry](https://leetcode.com/problems/curry/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1) per call  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func curry2New(fn func(int, int) int) *curry2`

## Solusi Go

```go
package main

// LeetCode #2632: Curry
// https://leetcode.com/problems/curry/
// Difficulty: Medium [Paid]
// Time: O(1) per call | Space: O(1)

import "fmt"

// curry2 implements currying for a 2-argument function
type curry2 struct {
	fn     func(int, int) int
	args   []int
	arity  int
}

func curry2New(fn func(int, int) int) *curry2 {
	return &curry2{fn: fn, arity: 2}
}

func (c *curry2) call(args ...int) *curry2 {
	c.args = append(c.args, args...)
	return c
}

func (c *curry2) done() int {
	return c.fn(c.args[0], c.args[1])
}

func add(a, b int) int {
	return a + b
}

func main() {
	fn := add

	// Test case 1: currying with two separate calls
	c := curry2New(fn)
	r1 := c.call(1).call(2).done()
	fmt.Println("Test 1:", r1)
	// Expected: 3

	// Test case 2: currying with one call
	c2 := curry2New(fn)
	r2 := c2.call(3, 4).done()
	fmt.Println("Test 2:", r2)
	// Expected: 7
}
```
