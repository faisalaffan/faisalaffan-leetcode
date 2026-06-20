# 2870 — Minimum Number Of Operations To Make Array Empty

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumNumberOfOperationsToMakeArrayEmpty(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2870: Minimum Number of Operations to Make Array Empty
// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-empty/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func MinimumNumberOfOperationsToMakeArrayEmpty(nums []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	ops := 0
	for _, c := range freq {
		if c == 1 {
			return -1
		}
		// Use as many 3s as possible
		ops += c / 3
		if c%3 != 0 {
			ops++
		}
	}

	return ops
}

func main() {
	fmt.Println(MinimumNumberOfOperationsToMakeArrayEmpty([]int{2, 3, 3, 3, 3, 2}))
	fmt.Println(MinimumNumberOfOperationsToMakeArrayEmpty([]int{1, 1, 1, 1}))
	fmt.Println(MinimumNumberOfOperationsToMakeArrayEmpty([]int{1, 2, 3}))
}
```
