# 0948 — Bag Of Tokens

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func bagOfTokensScore(tokens []int, power int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #948: Bag of Tokens
// https://leetcode.com/problems/bag-of-tokens/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

// Time: O(n log n) | Space: O(1)
func bagOfTokensScore(tokens []int, power int) int {
  // Sort O(n log n)
	sort.Ints(tokens)
	left, right := 0, len(tokens)-1
	score, maxScore := 0, 0

  // Binary search loop
	for left <= right {
		if power >= tokens[left] {
			power -= tokens[left]
			left++
			score++
			if score > maxScore {
				maxScore = score
			}
		} else if score > 0 {
			power += tokens[right]
			right--
			score--
		} else {
			break
		}
	}
	return maxScore
}

func main() {
	fmt.Println(bagOfTokensScore([]int{100}, 50))
	fmt.Println(bagOfTokensScore([]int{200, 100}, 150))
	fmt.Println(bagOfTokensScore([]int{100, 200, 300, 400}, 200))
}
```
