# 2789 — Largest Element In An Array After Merge Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LargestElementInAnArrayAfterMergeOperations(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2789: Largest Element in an Array after Merge Operations
// https://leetcode.com/problems/largest-element-in-an-array-after-merge-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func LargestElementInAnArrayAfterMergeOperations(nums []int) int64 {
	n := len(nums)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}

	result := int64(nums[n-1])
	for i := n - 2; i >= 0; i-- {
		if int64(nums[i]) <= result {
			result += int64(nums[i])
		} else {
			result = int64(nums[i])
		}
	}

	return result
}

func main() {
	fmt.Println(LargestElementInAnArrayAfterMergeOperations([]int{2, 3, 7, 9, 3}))
	fmt.Println(LargestElementInAnArrayAfterMergeOperations([]int{5, 3, 3}))
}
```
