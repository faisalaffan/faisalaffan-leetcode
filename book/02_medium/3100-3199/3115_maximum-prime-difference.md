# 3115 — Maximum Prime Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isPrime(n int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * sqrt(m))  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3115: Maximum Prime Difference
// https://leetcode.com/problems/maximum-prime-difference/
// Difficulty: Medium
// Time: O(n * sqrt(m)) | Space: O(1)

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func maximumPrimeDifference(nums []int) int {
	first, last := -1, -1
	for i, v := range nums {
		if isPrime(v) {
			if first == -1 {
				first = i
			}
			last = i
		}
	}
	return last - first
}

func main() {
	fmt.Println(maximumPrimeDifference([]int{4, 2, 9, 5, 3})) // Expected: 3
	fmt.Println(maximumPrimeDifference([]int{4, 8, 2, 8}))    // Expected: 0
}
```
