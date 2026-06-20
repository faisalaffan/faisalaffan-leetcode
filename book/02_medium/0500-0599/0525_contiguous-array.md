# 0525 — Contiguous Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindMaxLength(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #525: Contiguous Array
// https://leetcode.com/problems/contiguous-array/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(FindMaxLength([]int{0, 1}))
	fmt.Println(FindMaxLength([]int{0, 1, 0}))
}

func FindMaxLength(nums []int) int {
	// Map count -> first index. count = (#1 - #0)
  // Membuat map (HashMap) — pencarian O(1)
	countMap := make(map[int]int)
	countMap[0] = -1
	count := 0
	maxLen := 0

	for i, num := range nums {
		if num == 1 {
			count++
		} else {
			count--
		}
		if prevIdx, ok := countMap[count]; ok {
			if i-prevIdx > maxLen {
				maxLen = i - prevIdx
			}
		} else {
			countMap[count] = i
		}
	}

	return maxLen
}
```
