# 1497 — Check If Array Pairs Are Divisible By K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CanArrange(arr []int, k int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N), Space: O(K)  
**Kompleksitas Ruang:** O(K)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1497: Check If Array Pairs Are Divisible by k
// https://leetcode.com/problems/check-if-array-pairs-are-divisible-by-k/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CanArrange([]int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, 5))
	fmt.Println(CanArrange([]int{1, 2, 3, 4, 5, 6}, 7))
	fmt.Println(CanArrange([]int{1, 2, 3, 4, 5, 6}, 10))
}

func CanArrange(arr []int, k int) bool {
	// Time: O(N), Space: O(K)
  // Alokasi slice integer
	remainder := make([]int, k)
	for _, num := range arr {
		r := ((num % k) + k) % k
		remainder[r]++
	}

	// Numbers divisible by k must pair among themselves
	if remainder[0]%2 != 0 {
		return false
	}

	// For i and k-i, their counts must match
	for i := 1; i < k; i++ {
		if remainder[i] != remainder[k-i] {
			return false
		}
	}

	return true
}
```
