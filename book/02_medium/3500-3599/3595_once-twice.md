# 3595 — Once Twice

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func OnceTwice(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3595: Once Twice
// https://leetcode.com/problems/once-twice/
// Difficulty: Medium [Paid]
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", OnceTwice([]int{1, 1, 2, 2, 3}))
	// Test case 2
	fmt.Println("Test 2:", OnceTwice([]int{1, 1, 1, 2, 2}))
	// Test case 3
	fmt.Println("Test 3:", OnceTwice([]int{1, 2, 3}))
}

func OnceTwice(nums []int) int {
	// Count elements that appear twice vs once
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	once, twice := 0, 0
	for _, c := range freq {
		if c == 1 {
			once++
		} else if c >= 2 {
			twice++
		}
	}
	return once * twice
}
```
