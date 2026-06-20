# 2543 — Check If Point Is Reachable

## Deskripsi

**Soal:** [2543. Check If Point Is Reachable](https://leetcode.com/problems/check-if-point-is-reachable/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func isReachable(targetX int, targetY int) bool`

## Solusi Go

```go
package main

// LeetCode #2543: Check if Point Is Reachable
// https://leetcode.com/problems/check-if-point-is-reachable/
// Difficulty: Hard

import "fmt"

// isReachable checks if (targetX, targetY) is reachable from (1, 1) using:
//   (x, y) -> (x, y-x), (x-y, y), (2x, y), (x, 2y)
//
// Key insight: (x, y) is reachable iff gcd(x, y) is a power of 2.
// The subtraction operations preserve GCD. Doubling multiplies a coordinate by 2.
// Starting from (1, 1), gcd = 1 = 2^0. The invariant is that gcd is always a power of 2.
// Conversely, any (x, y) with gcd = 2^k can be reduced back to (1, 1) by reversing
// the operations (reverse subtraction = addition; reverse doubling = halving).
//
// Complexity: O(log(min(x,y))) time, O(1) space
func isReachable(targetX int, targetY int) bool {
	g := gcd(targetX, targetY)
	// Check if g is a power of 2: a power of 2 has exactly one bit set
	return g&(g-1) == 0
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	// Test cases
	fmt.Println("Test 1: targetX=6, targetY=9 ->", isReachable(6, 9))   // false
	fmt.Println("Test 2: targetX=4, targetY=7 ->", isReachable(4, 7))   // false
	fmt.Println("Test 3: targetX=1, targetY=1 ->", isReachable(1, 1))   // true
	fmt.Println("Test 4: targetX=2, targetY=4 ->", isReachable(2, 4))   // true (gcd=2)
	fmt.Println("Test 5: targetX=3, targetY=5 ->", isReachable(3, 5))   // true (gcd=1=2^0)
	fmt.Println("Test 6: targetX=8, targetY=12 ->", isReachable(8, 12)) // true (gcd=4=2^2)
	fmt.Println("Test 7: targetX=4, targetY=8 ->", isReachable(4, 8))   // true (gcd=4=2^2)
	fmt.Println("Test 8: targetX=2, targetY=3 ->", isReachable(2, 3))   // true (gcd=1=2^0)
}
```
