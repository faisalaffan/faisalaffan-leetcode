# 0027 — Remove Element

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func RemoveElement(nums []int, val int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #27: Remove Element
// https://leetcode.com/problems/remove-element/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func RemoveElement(nums []int, val int) int {
	k := 0
  // Linear scan O(n)
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	n1 := []int{3, 2, 2, 3}
	fmt.Println(RemoveElement(n1, 3), n1)
	n2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(RemoveElement(n2, 2), n2)
}
```
