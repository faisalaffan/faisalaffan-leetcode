# 3512 — Minimum Operations To Make Array Sum Divisible By K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumOperationsToMakeArraySumDivisibleByK(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3512: Minimum Operations to Make Array Sum Divisible by K
// https://leetcode.com/problems/minimum-operations-to-make-array-sum-divisible-by-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumOperationsToMakeArraySumDivisibleByK([]int{3, 9, 7}, 5))
	fmt.Println(MinimumOperationsToMakeArraySumDivisibleByK([]int{4, 1, 3}, 4))
}

// MinimumOperationsToMakeArraySumDivisibleByK returns the min operations (incrementing elements by 1) to make sum divisible by k.
// Time: O(n). Space: O(1).
func MinimumOperationsToMakeArraySumDivisibleByK(nums []int, k int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	rem := sum % k
	if rem == 0 {
		return 0
	}
	return rem
}
```
