# 2610 — Convert An Array Into A 2D Array With Conditions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func findMatrix(nums []int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2610: Convert an Array Into a 2D Array With Conditions
// https://leetcode.com/problems/convert-an-array-into-a-2d-array-with-conditions/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func findMatrix(nums []int) [][]int {
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	ans := [][]int{}

	for _, v := range nums {
		freq[v]++
		if freq[v] > len(ans) {
			ans = append(ans, []int{})
		}
		ans[freq[v]-1] = append(ans[freq[v]-1], v)
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findMatrix([]int{1, 3, 4, 1, 2, 3, 1}))
	// Expected: [[1,3,4,2],[1,3],[1]]

	// Test case 2
	fmt.Println("Test 2:", findMatrix([]int{1, 2, 3, 4}))
	// Expected: [[1,2,3,4]]

	// Test case 3
	fmt.Println("Test 3:", findMatrix([]int{1, 1, 1, 1}))
	// Expected: [[1],[1],[1],[1]]
}
```
