# 3637 — Trionic Array I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func TrionicArrayI(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3637: Trionic Array I
// https://leetcode.com/problems/trionic-array-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TrionicArrayI([]int{1, 3, 5, 4, 2, 6}))
	fmt.Println(TrionicArrayI([]int{2, 1, 3}))
}

// Time: O(n)
// Space: O(1)
func TrionicArrayI(nums []int) bool {
	n := len(nums)
	if n < 4 || nums[0] >= nums[1] || nums[n-2] >= nums[n-1] {
		return false
	}

	i := 1
	// Phase 1: strictly increasing
	for i < n && nums[i] > nums[i-1] {
		i++
	}
	p := i - 1

	// Phase 2: strictly decreasing
	for i < n && nums[i] < nums[i-1] {
		i++
	}
	q := i - 1

	// Phase 3: strictly increasing
	for i < n && nums[i] > nums[i-1] {
		i++
	}

	return i == n && p > 0 && q > p && q < n-1
}
```
