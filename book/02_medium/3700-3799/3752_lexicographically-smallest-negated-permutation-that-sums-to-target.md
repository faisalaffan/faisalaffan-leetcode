# 3752 — Lexicographically Smallest Negated Permutation That Sums To Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func lexicographicallySmallestNegatedPermutationThatSumsToTarget(n int, target int64) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

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

  // Alokasi slice
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
