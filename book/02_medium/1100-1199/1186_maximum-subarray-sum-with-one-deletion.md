# 1186 — Maximum Subarray Sum With One Deletion

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumSum(arr []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1186: Maximum Subarray Sum with One Deletion
// https://leetcode.com/problems/maximum-subarray-sum-with-one-deletion/
// Difficulty: Medium

// Kadane's algorithm variant with one deletion allowed.
// dp_no_del[i] = max subarray sum ending at i without deletion
// dp_del[i] = max subarray sum ending at i with one deletion

// Time: O(n)
// Space: O(1)

func maximumSum(arr []int) int {
	n := len(arr)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}

	noDel := arr[0]
	withDel := arr[0]
	result := arr[0]

	for i := 1; i < n; i++ {
		withDel = max(withDel+arr[i], noDel)
		noDel = max(noDel+arr[i], arr[i])
		result = max(result, max(noDel, withDel))
	}

	return result
}

func main() {
	fmt.Printf("%d (expected: 4)\n", maximumSum([]int{1, -2, 0, 3}))
	fmt.Printf("%d (expected: -1)\n", maximumSum([]int{-1, -1, -1, -1}))
	fmt.Printf("%d (expected: 7)\n", maximumSum([]int{1, -2, -2, 3, -1, 4}))
}
```
