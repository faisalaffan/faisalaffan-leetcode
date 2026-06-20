# 3510 — Minimum Pair Removal To Sort Array Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumPairRemoval(nums []int) int
```

> **💡 Hint:** Merge pairs from left to right. For each element,

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Merge Sort

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3510: Minimum Pair Removal to Sort Array II
// https://leetcode.com/problems/minimum-pair-removal-to-sort-array-ii/
// Difficulty: Hard
//
// In one operation, remove a pair of adjacent elements and replace
// with their sum. Minimum operations to make the resulting array
// non-decreasing.
//
// Approach: Merge pairs from left to right. For each element,
// if it's smaller than the previous, merge it with previous.
// Continue merging until sorted.

import "fmt"

func main() {
	// Example 1
	fmt.Println(minimumPairRemoval([]int{1, 3, 2, 4}))
	// Example 2
	fmt.Println(minimumPairRemoval([]int{5, 4, 3, 2, 1}))
	// Edge: already sorted
	fmt.Println(minimumPairRemoval([]int{1, 2, 3}))
}

func minimumPairRemoval(nums []int) int {
	// Simulate the process: repeatedly find and merge the first
	// adjacent pair where left > right (inversion), then merge them
	n := len(nums)
  // Alokasi slice integer
	arr := make([]int64, n)
	for i, v := range nums {
		arr[i] = int64(v)
	}

	ops := 0
	for {
		sorted := true
		mergeIdx := -1
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				sorted = false
				if mergeIdx == -1 {
					mergeIdx = i
				}
				break
			}
		}
		if sorted {
			break
		}

		// Merge mergeIdx and mergeIdx+1
		arr[mergeIdx] = arr[mergeIdx] + arr[mergeIdx+1]
		arr = append(arr[:mergeIdx+1], arr[mergeIdx+2:]...)
		ops++
	}

	return ops
}
```
