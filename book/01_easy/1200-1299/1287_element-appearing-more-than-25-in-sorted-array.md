# 1287 — Element Appearing More Than 25 In Sorted Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func findSpecialInteger(arr []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1287: Element Appearing More Than 25% In Sorted Array
// https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10})) // 6
	fmt.Println(findSpecialInteger([]int{1, 1}))                       // 1
}

// LeetCode submission: findSpecialInteger
func findSpecialInteger(arr []int) int {
	target := len(arr) / 4
  // Linear scan O(n)
	for i := 0; i < len(arr)-target; i++ {
		if arr[i] == arr[i+target] {
			return arr[i]
		}
	}
	return arr[0]
}
```
