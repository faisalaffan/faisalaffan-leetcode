# 3752 — Lexicographically Smallest Negated Permutation That Sums To Target

## Deskripsi

**Soal:** [3752. Lexicographically Smallest Negated Permutation That Sums To Target](https://leetcode.com/problems/lexicographically-smallest-negated-permutation-that-sums-to-target/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func lexicographicallySmallestNegatedPermutationThatSumsToTarget(n int, target int64) []int`

## Solusi Go

```go
package main

// LeetCode #3752: Lexicographically Smallest Negated Permutation that Sums to Target
// https://leetcode.com/problems/lexicographically-smallest-negated-permutation-that-sums-to-target/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func lexicographicallySmallestNegatedPermutationThatSumsToTarget(n int, target int64) []int {
	s := int64(n) * int64(n+1) / 2
	drop := s - target
	if drop < 0 || drop%2 != 0 {
		return []int{}
	}

	delta := drop / 2
  // Membuat slice untuk menyimpan hasil
	used := make([]bool, n+1)
	negatedSum := int64(0)

	// Greedily pick largest numbers to negate
	for i := n; i > 0; i-- {
		if negatedSum+int64(i) <= delta {
			used[i] = true
			negatedSum += int64(i)
		}
	}

	if negatedSum != delta {
		return []int{}
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]int, 0, n)
	// Negated numbers first (from largest to smallest = lexicographically smallest)
	for i := n; i > 0; i-- {
		if used[i] {
			ans = append(ans, -i)
		}
	}
	// Then positive numbers in ascending order
	for i := 1; i <= n; i++ {
		if !used[i] {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	fmt.Println(lexicographicallySmallestNegatedPermutationThatSumsToTarget(3, 0))
	fmt.Println(lexicographicallySmallestNegatedPermutationThatSumsToTarget(4, 10))
	fmt.Println(lexicographicallySmallestNegatedPermutationThatSumsToTarget(5, 15))
}
```
