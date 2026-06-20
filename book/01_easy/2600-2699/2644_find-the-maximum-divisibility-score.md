# 2644 — Find The Maximum Divisibility Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindTheMaximumDivisibilityScore(nums []int, divisors []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2644: Find the Maximum Divisibility Score
// https://leetcode.com/problems/find-the-maximum-divisibility-score/
// Difficulty: Easy
// Time: O(|nums| * |divisors|) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(FindTheMaximumDivisibilityScore([]int{2, 3, 4, 5, 6}, []int{2, 3, 4}))
	fmt.Println(FindTheMaximumDivisibilityScore([]int{4, 7, 9, 3, 9}, []int{5, 2, 3}))
}

func FindTheMaximumDivisibilityScore(nums []int, divisors []int) int {
	ans := divisors[0]
	maxScore := 0
	for _, d := range divisors {
		score := 0
		for _, n := range nums {
			if n%d == 0 {
				score++
			}
		}
		if score > maxScore || (score == maxScore && d < ans) {
			maxScore = score
			ans = d
		}
	}
	return ans
}
```
