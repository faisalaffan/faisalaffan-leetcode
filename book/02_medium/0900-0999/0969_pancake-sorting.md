# 0969 — Pancake Sorting

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func pancakeSort(arr []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #969: Pancake Sorting
// https://leetcode.com/problems/pancake-sorting/
// Difficulty: Medium

import "fmt"

// Time: O(n^2) | Space: O(n)
func pancakeSort(arr []int) []int {
  // Alokasi slice integer
	ans := make([]int, 0)
	n := len(arr)

	for i := n; i > 1; i-- {
		// Find index of max in arr[:i]
		maxIdx := 0
		for j := 1; j < i; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		if maxIdx == i-1 {
			continue
		}
		// Flip to bring max to front
		if maxIdx > 0 {
			reverse(arr, maxIdx)
			ans = append(ans, maxIdx+1)
		}
		// Flip to put max at correct position
		reverse(arr, i-1)
		ans = append(ans, i)
	}

	return ans
}

func reverse(arr []int, end int) {
	for i, j := 0, end; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	fmt.Println(pancakeSort([]int{3, 2, 4, 1}))
	fmt.Println(pancakeSort([]int{1, 2, 3}))
}
```
