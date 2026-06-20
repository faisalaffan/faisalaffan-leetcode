# 1089 — Duplicate Zeros

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func duplicateZeros(arr []int) `

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1089: Duplicate Zeros
// https://leetcode.com/problems/duplicate-zeros/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	arr1 := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr1)
	fmt.Println(arr1) // [1,0,0,2,3,0,0,4]

	arr2 := []int{1, 2, 3}
	duplicateZeros(arr2)
	fmt.Println(arr2) // [1,2,3]
}

// LeetCode submission: duplicateZeros
func duplicateZeros(arr []int) {
	n := len(arr)
	possibleDups := 0
	for i := 0; i+possibleDups < n; i++ {
		if arr[i] == 0 {
			possibleDups++
		}
	}
	last := n - 1 - possibleDups
	for i := last; i >= 0; i-- {
		if i+possibleDups < n {
			arr[i+possibleDups] = arr[i]
		}
		if arr[i] == 0 {
			possibleDups--
			if i+possibleDups < n {
				arr[i+possibleDups] = 0
			}
		}
	}
}
```
