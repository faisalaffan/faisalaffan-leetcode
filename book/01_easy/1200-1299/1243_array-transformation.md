# 1243 — Array Transformation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func transformArray(arr []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2) worst case  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1243: Array Transformation
// https://leetcode.com/problems/array-transformation/
// Difficulty: Easy [Paid]
// Time: O(n^2) worst case | Space: O(n)

import "fmt"

func main() {
	fmt.Println(transformArray([]int{6, 2, 3, 4})) // [6,3,3,4]
	fmt.Println(transformArray([]int{1, 6, 3, 4, 3, 5})) // [1,4,4,4,4,5]
}

// LeetCode submission: transformArray
func transformArray(arr []int) []int {
	if len(arr) <= 2 {
		return append([]int{}, arr...)
	}
	for {
		changed := false
  // Alokasi slice
		next := make([]int, len(arr))
		copy(next, arr)
		for i := 1; i < len(arr)-1; i++ {
			if arr[i] < arr[i-1] && arr[i] < arr[i+1] {
				next[i]++
				changed = true
			} else if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
				next[i]--
				changed = true
			}
		}
		if !changed {
			break
		}
		arr = next
	}
	return arr
}
```
