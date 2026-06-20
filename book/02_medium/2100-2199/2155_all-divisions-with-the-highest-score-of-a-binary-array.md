# 2155 — All Divisions With The Highest Score Of A Binary Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxScoreIndices(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2155: All Divisions With the Highest Score of a Binary Array
// https://leetcode.com/problems/all-divisions-with-the-highest-score-of-a-binary-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func maxScoreIndices(nums []int) []int {
	n := len(nums)
	totalOnes := 0
	for _, v := range nums {
		totalOnes += v
	}

	maxScore := totalOnes // score at index 0 (0 zeros left + totalOnes ones right)
	result := []int{0}
	zerosLeft := 0

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			zerosLeft++
		}
		onesRight := totalOnes - (i + 1 - zerosLeft) // total - ones so far
		score := zerosLeft + onesRight
		if score > maxScore {
			maxScore = score
			result = []int{i + 1}
		} else if score == maxScore {
			result = append(result, i+1)
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxScoreIndices([]int{0, 0, 1, 0}))
	// Expected: [2, 4]

	// Test case 2
	fmt.Println("Test 2:", maxScoreIndices([]int{0, 0, 0}))
	// Expected: [3]

	// Test case 3
	fmt.Println("Test 3:", maxScoreIndices([]int{1, 1}))
	// Expected: [0]
}
```
