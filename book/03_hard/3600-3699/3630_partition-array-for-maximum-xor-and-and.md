# 3630 — Partition Array For Maximum Xor And And

## Deskripsi

**Soal:** [3630. Partition Array For Maximum Xor And And](https://leetcode.com/problems/partition-array-for-maximum-xor-and-and/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), Bitmask (representasi himpunan dengan bit)

> **Ide Kunci:** Bitmask DP over subsets. n <= 19, so 3^n is too large but we can

## Solusi Go

```go
package main

// LeetCode #3630: Partition Array for Maximum XOR and AND
// https://leetcode.com/problems/partition-array-for-maximum-xor-and-and/
// Difficulty: Hard
//
// Partition array into three subsequences A, B, C (each element in exactly one)
// to maximize XOR(A) + AND(B) + XOR(C). XOR(empty) = 0, AND(empty) = 0.
//
// Approach: Bitmask DP over subsets. n <= 19, so 3^n is too large but we can
// enumerate subsets for one partition and compute remaining values.

import "fmt"
import "math"

func main() {
	// Example 1
	fmt.Println(maximizeXorAndXor([]int{1, 2, 3, 4}))
	// Example 2
	fmt.Println(maximizeXorAndXor([]int{5, 1, 6}))
	// Edge: single element
	fmt.Println(maximizeXorAndXor([]int{10}))
	// Edge: all zeros
	fmt.Println(maximizeXorAndXor([]int{0, 0, 0}))
}

func maximizeXorAndXor(nums []int) int64 {
	n := len(nums)
	total := 1 << uint(n)

	// Precompute XOR for all subsets
  // Membuat slice untuk menyimpan hasil
	xor := make([]int64, total)
  // Membuat slice untuk menyimpan hasil
	and := make([]int64, total)
	for mask := 1; mask < total; mask++ {
		lsb := mask & -mask
		bit := int(math.Log2(float64(lsb)))
		prev := mask ^ lsb
		xor[mask] = xor[prev] ^ int64(nums[bit])
		and[mask] = and[prev] & int64(nums[bit])
		if prev == 0 {
			and[mask] = int64(nums[bit])
		}
	}

	var result int64

	// Enumerate A (maskA), then B (maskB) as subset of remaining, C = remaining ^ maskB
	remaining := total - 1
	for maskA := 0; maskA < total; maskA++ {
		xorA := xor[maskA]
		rest := remaining ^ maskA
		// Enumerate subsets of rest for B
		maskB := rest
		for {
			andB := and[maskB]
			xorC := xor[rest^maskB]
			val := xorA + andB + xorC
			if val > result {
				result = val
			}
			if maskB == 0 {
				break
			}
			maskB = (maskB - 1) & rest
		}
	}

	return result
}
```
