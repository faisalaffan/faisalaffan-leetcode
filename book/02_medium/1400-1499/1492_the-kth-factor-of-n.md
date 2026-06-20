# 1492 — The Kth Factor Of N

## Deskripsi

**Soal:** [1492. The Kth Factor Of N](https://leetcode.com/problems/the-kth-factor-of-n/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(sqrt(N)), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1492: The kth Factor of n
// https://leetcode.com/problems/the-kth-factor-of-n/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(KthFactor(12, 3))
	fmt.Println(KthFactor(7, 2))
	fmt.Println(KthFactor(4, 4))
}

func KthFactor(n int, k int) int {
	// Time: O(sqrt(N)), Space: O(1)
	// Count factors from 1 to sqrt(n)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			k--
			if k == 0 {
				return i
			}
		}
	}

	// Count factors from sqrt(n) down to 1 (the paired factors)
	// Start from the largest paired factor
	for i := intSqrt(n); i >= 1; i-- {
		if n%i == 0 && i*i != n { // don't double count perfect square root
			k--
			if k == 0 {
				return n / i
			}
		}
	}

	return -1
}

func intSqrt(n int) int {
	for i := 1; i*i <= n; i++ {
		if i*i == n {
			return i
		}
	}
	// floor sqrt
	result := 0
	for result*result <= n {
		result++
	}
	return result - 1
}
```
