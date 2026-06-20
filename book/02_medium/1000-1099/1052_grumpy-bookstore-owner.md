# 1052 — Grumpy Bookstore Owner

## Deskripsi

**Soal:** [1052. Grumpy Bookstore Owner](https://leetcode.com/problems/grumpy-bookstore-owner/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Sliding Window (jendela geser)

> **Ide Kunci:** Sliding window - find best minutes to use secret technique

## Solusi Go

```go
package main

// LeetCode #1052: Grumpy Bookstore Owner
// https://leetcode.com/problems/grumpy-bookstore-owner/
// Difficulty: Medium
//
// Approach: Sliding window - find best minutes to use secret technique
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3)) // 16
	fmt.Println(maxSatisfied([]int{1}, []int{0}, 1)) // 1
}

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n := len(customers)
	baseSatisfied := 0
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			baseSatisfied += customers[i]
		}
	}

	extraSatisfied := 0
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			extraSatisfied += customers[i]
		}
	}

	maxExtra := extraSatisfied
	for i := minutes; i < n; i++ {
		if grumpy[i-minutes] == 1 {
			extraSatisfied -= customers[i-minutes]
		}
		if grumpy[i] == 1 {
			extraSatisfied += customers[i]
		}
		if extraSatisfied > maxExtra {
			maxExtra = extraSatisfied
		}
	}

	return baseSatisfied + maxExtra
}
```
