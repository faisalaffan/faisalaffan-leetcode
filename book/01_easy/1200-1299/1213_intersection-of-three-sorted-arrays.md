# 1213 — Intersection Of Three Sorted Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func arraysIntersection(arr1, arr2, arr3 []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1) excluding output


## 💻 Solusi Go

```go
package main

// LeetCode #1213: Intersection of Three Sorted Arrays
// https://leetcode.com/problems/intersection-of-three-sorted-arrays/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1) excluding output

import "fmt"

func main() {
	fmt.Println(arraysIntersection([]int{1, 2, 3, 4, 5}, []int{1, 2, 5, 7, 9}, []int{1, 3, 4, 5, 8}))
	// [1,5]
	fmt.Println(arraysIntersection([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}))
	// []
}

// LeetCode submission: arraysIntersection
func arraysIntersection(arr1, arr2, arr3 []int) []int {
	var ans []int
	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) && k < len(arr3) {
		if arr1[i] == arr2[j] && arr2[j] == arr3[k] {
			ans = append(ans, arr1[i])
			i++; j++; k++
		} else if arr1[i] < arr2[j] {
			i++
		} else if arr2[j] < arr3[k] {
			j++
		} else {
			k++
		}
	}
	return ans
}
```
