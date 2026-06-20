# 2749 — Minimum Operations To Make The Integer Zero

## Deskripsi

**Soal:** [2749. Minimum Operations To Make The Integer Zero](https://leetcode.com/problems/minimum-operations-to-make-the-integer-zero/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func MinimumOperationsToMakeTheIntegerZero(num1 int, num2 int) int`

## Solusi Go

```go
package main

// LeetCode #2749: Minimum Operations to Make the Integer Zero
// https://leetcode.com/problems/minimum-operations-to-make-the-integer-zero/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import "fmt"

func MinimumOperationsToMakeTheIntegerZero(num1 int, num2 int) int {
	// Try k operations: num1 - k*num2 must have at most k bits set, >= k
	for k := 1; k <= 60; k++ {
		diff := num1 - k*num2
		if diff < 0 {
			return -1
		}
		bits := 0
		for tmp := diff; tmp > 0; tmp >>= 1 {
			bits += tmp & 1
		}
		if bits <= k && diff >= k {
			return k
		}
	}
	return -1
}

func main() {
	fmt.Println(MinimumOperationsToMakeTheIntegerZero(3, -2))
	fmt.Println(MinimumOperationsToMakeTheIntegerZero(5, 7))
}
```
