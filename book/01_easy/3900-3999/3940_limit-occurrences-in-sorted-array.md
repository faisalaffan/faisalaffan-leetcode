# 3940 — Limit Occurrences In Sorted Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LimitOccurrencesInSortedArray(nums []int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3940: Limit Occurrences in Sorted Array
// https://leetcode.com/problems/limit-occurrences-in-sorted-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(LimitOccurrencesInSortedArray([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(LimitOccurrencesInSortedArray([]int{1, 2, 3}, 1))
}

// Time: O(n)
// Space: O(1)
func LimitOccurrencesInSortedArray(nums []int, k int) []int {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return nums
	}
	write := 0
	for _, v := range nums {
		if write < k || v != nums[write-k] {
			nums[write] = v
			write++
		}
	}
	return nums[:write]
}
```
