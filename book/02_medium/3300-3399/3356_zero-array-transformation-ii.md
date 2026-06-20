# 3356 — Zero Array Transformation Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minZeroArray(nums []int, queries [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O((n + q) log q) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3356: Zero Array Transformation II
// https://leetcode.com/problems/zero-array-transformation-ii/
// Difficulty: Medium
// Time: O((n + q) log q) Space: O(n)

import "fmt"

func main() {
	fmt.Println(minZeroArray([]int{2, 0, 2}, [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}})) // 2
	fmt.Println(minZeroArray([]int{4, 3, 2, 1}, [][]int{{1, 3, 2}, {0, 2, 1}}))         // -1
}

func minZeroArray(nums []int, queries [][]int) int {
	m := len(queries)

	// Check if already zero
	allZero := true
	for _, v := range nums {
		if v != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return 0
	}

	// Check if impossible
	if !canMakeZero(nums, queries, m) {
		return -1
	}

	lo, hi := 1, m
	for lo < hi {
		mid := (lo + hi) / 2
		if canMakeZero(nums, queries, mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func canMakeZero(nums []int, queries [][]int, k int) bool {
	n := len(nums)
  // Alokasi slice integer
	diff := make([]int, n+1)
	for i := 0; i < k; i++ {
		l, r, val := queries[i][0], queries[i][1], queries[i][2]
		diff[l] += val
		diff[r+1] -= val
	}
	cur := 0
	for i := 0; i < n; i++ {
		cur += diff[i]
		if cur < nums[i] {
			return false
		}
	}
	return true
}
```
