# 2161 — Partition Array According To Given Pivot

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func pivotArray(nums []int, pivot int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2161: Partition Array According to Given Pivot
// https://leetcode.com/problems/partition-array-according-to-given-pivot/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func pivotArray(nums []int, pivot int) []int {
	less, equal, greater := []int{}, []int{}, []int{}
	for _, v := range nums {
		if v < pivot {
			less = append(less, v)
		} else if v == pivot {
			equal = append(equal, v)
		} else {
			greater = append(greater, v)
		}
	}
	return append(append(less, equal...), greater...)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", pivotArray([]int{9, 12, 5, 10, 14, 3, 10}, 10))
	// Expected: [9,5,3,10,10,12,14]

	// Test case 2
	fmt.Println("Test 2:", pivotArray([]int{-3, 4, 3, 2}, 2))
	// Expected: [-3,2,4,3]

	// Test case 3
	fmt.Println("Test 3:", pivotArray([]int{1, 2, 3, 4, 5}, 3))
	// Expected: [1,2,3,4,5]
}
```
