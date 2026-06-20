# 3012 — Minimize Length Of Array Using Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumArrayLength(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3012: Minimize Length of Array Using Operations
// https://leetcode.com/problems/minimize-length-of-array-using-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minimumArrayLength([]int{1, 4, 3, 1}))
	fmt.Println(minimumArrayLength([]int{5, 5, 5, 10, 5}))
	fmt.Println(minimumArrayLength([]int{3, 5}))
}

func minimumArrayLength(nums []int) int {
	minVal := nums[0]
	for _, x := range nums[1:] {
		if x < minVal {
			minVal = x
		}
	}
	cnt := 0
	for _, x := range nums {
		if x%minVal != 0 {
			return 1
		}
		if x == minVal {
			cnt++
		}
	}
	return (cnt + 1) / 2
}
```
