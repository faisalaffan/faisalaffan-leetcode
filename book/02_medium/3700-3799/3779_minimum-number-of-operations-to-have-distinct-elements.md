# 3779 — Minimum Number Of Operations To Have Distinct Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumNumberOfOperationsToHaveDistinctElements(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3779: Minimum Number of Operations to Have Distinct Elements
// https://leetcode.com/problems/minimum-number-of-operations-to-have-distinct-elements/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumNumberOfOperationsToHaveDistinctElements(nums []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]bool)
	for i := len(nums) - 1; i >= 0; i-- {
		if seen[nums[i]] {
			return (i + 3) / 3
		}
		seen[nums[i]] = true
	}
	return 0
}

func main() {
	fmt.Println(minimumNumberOfOperationsToHaveDistinctElements([]int{3, 8, 3, 6, 5, 8}))
	fmt.Println(minimumNumberOfOperationsToHaveDistinctElements([]int{2, 2}))
	fmt.Println(minimumNumberOfOperationsToHaveDistinctElements([]int{4, 3, 5, 1, 2}))
}
```
