# 2784 — Check If Array Is Good

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckIfArrayIsGood(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2784: Check if Array is Good
// https://leetcode.com/problems/check-if-array-is-good/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(CheckIfArrayIsGood([]int{2, 1, 3}))
	fmt.Println(CheckIfArrayIsGood([]int{3, 4, 4, 1, 2, 1}))
}

func CheckIfArrayIsGood(nums []int) bool {
	n := len(nums) - 1
  // Alokasi slice integer
	counts := make([]int, n+1)
	for _, v := range nums {
		if v > n {
			return false
		}
		counts[v]++
	}
	if counts[n] != 2 {
		return false
	}
	for i := 1; i < n; i++ {
		if counts[i] != 1 {
			return false
		}
	}
	return true
}
```
