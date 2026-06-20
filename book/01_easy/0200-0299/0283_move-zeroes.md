# 0283 — Move Zeroes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MoveZeroes(nums []int) `

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #283: Move Zeroes
// https://leetcode.com/problems/move-zeroes/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func MoveZeroes(nums []int) {
	lastNonZero := 0
  // Linear scan O(n)
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[lastNonZero] = nums[lastNonZero], nums[i]
			lastNonZero++
		}
	}
}

func main() {
	n1 := []int{0, 1, 0, 3, 12}
	MoveZeroes(n1)
	fmt.Println(n1)
	n2 := []int{0}
	MoveZeroes(n2)
	fmt.Println(n2)
}
```
