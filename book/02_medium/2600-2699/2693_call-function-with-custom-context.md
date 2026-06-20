# 2693 — Call Function With Custom Context

## Deskripsi

**Soal:** [2693. Call Function With Custom Context](https://leetcode.com/problems/call-function-with-custom-context/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func callWithContext(fn func(Context, ...int) int, ctx Context, args ...int) int`

## Solusi Go

```go
package main

// LeetCode #2693: Call Function with Custom Context
// https://leetcode.com/problems/call-function-with-custom-context/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

type Context map[string]any

func callWithContext(fn func(Context, ...int) int, ctx Context, args ...int) int {
	return fn(ctx, args...)
}

func main() {
	// Test case 1
	sumFn := func(ctx Context, args ...int) int {
		total := 0
		for _, v := range args {
			total += v
		}
		return total
	}
	fmt.Println("Test 1:", callWithContext(sumFn, Context{}, 1, 2, 3))
	// Expected: 6

	// Test case 2
	ctxFn := func(ctx Context, args ...int) int {
		if multiplier, ok := ctx["mult"]; ok {
			m := multiplier.(int)
			result := 0
			for _, v := range args {
				result += v * m
			}
			return result
		}
		return 0
	}
	fmt.Println("Test 2:", callWithContext(ctxFn, Context{"mult": 3}, 1, 2, 3))
	// Expected: 18

	// Test case 3
	fmt.Println("Test 3:", callWithContext(sumFn, Context{}))
	// Expected: 0
}
```
