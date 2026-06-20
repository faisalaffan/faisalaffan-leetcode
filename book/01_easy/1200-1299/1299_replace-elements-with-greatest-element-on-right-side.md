# 1299 — Replace Elements With Greatest Element On Right Side

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func replaceElements(arr []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1) excluding output


## 💻 Solusi Go

```go
package main

// LeetCode #1299: Replace Elements with Greatest Element on Right Side
// https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/
// Difficulty: Easy
// Time: O(n) | Space: O(1) excluding output

import "fmt"

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1})) // [18,6,6,6,1,-1]
	fmt.Println(replaceElements([]int{400}))                 // [-1]
}

// LeetCode submission: replaceElements
func replaceElements(arr []int) []int {
  // Alokasi slice
	ans := make([]int, len(arr))
	maxRight := -1
	for i := len(arr) - 1; i >= 0; i-- {
		ans[i] = maxRight
		if arr[i] > maxRight {
			maxRight = arr[i]
		}
	}
	return ans
}
```
