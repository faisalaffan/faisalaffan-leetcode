# 3355 — Zero Array Transformation I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isZeroArray(nums []int, queries [][]int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + q) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3355: Zero Array Transformation I
// https://leetcode.com/problems/zero-array-transformation-i/
// Difficulty: Medium
// Time: O(n + q) Space: O(n)

import "fmt"

func main() {
	fmt.Println(isZeroArray([]int{1, 0, 1}, [][]int{{0, 2}, {0, 2}})) // true
	fmt.Println(isZeroArray([]int{2, 0, 2}, [][]int{{0, 2}, {0, 2}, {1, 1}})) // true
}

func isZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
  // Alokasi slice integer
	diff := make([]int, n+1)

	for _, q := range queries {
		l, r := q[0], q[1]
		diff[l]++
		diff[r+1]--
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
